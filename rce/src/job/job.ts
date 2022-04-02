import path from "path";
import fs from "fs/promises";
import console from "node:console";
import childProcess from "child_process";
import { Runtime } from "@/runtime/runtime";
import { User } from "@/user/user";

export interface JobPrerequisites {
    user: User
    runtime: Runtime
    code: string
    timeout?: number
    memoryLimit?: number
}

export interface CommandOutput {
    stdout: string
    stderr: string
    output: string
    exitCode: number
    signal: string
}

export class Job implements JobPrerequisites {
    private sourceFilePath: string;
    private builtFilePath: string;

    constructor(
        public user: User,
        public runtime: Runtime,
        public code: string,
        public timeout?: number,
        public memoryLimit?: number
    ) {
        if (user === undefined
            || Object.keys(user).length === 0
            || runtime === undefined
            || Object.keys(runtime).length === 0
            || code === null
            || code === undefined
            || code === "") {
            throw new TypeError("Invalid job parameters");
        }

        if (timeout === undefined || timeout < 1) {
            this.timeout = 5_000;
        }

        if (memoryLimit === undefined || memoryLimit < 1) {
            this.memoryLimit = 128 * 1024 * 1024;
        }

        this.sourceFilePath = "";
        this.builtFilePath = "";
    }

    async createFile(): Promise<void> {
        const filePath = path.join("/code", `/${this.user.username}`, `/code.${this.runtime.extension}`);
        await fs.writeFile(filePath, this.code, { encoding: "utf-8" });
        await fs.chmod(filePath, 0o700);
        await fs.chown(filePath, this.user.uid, this.user.gid);

        // Make sure the file is written properly.
        const stat = await fs.stat(filePath);
        console.log(`File path: ${filePath}`);
        console.log(`File stat: ${stat.uid} ${stat.gid} ${stat.mode} ${stat.size}`);
        this.sourceFilePath = filePath;
    }

    async compile(): Promise<void> {
        if (!this.runtime.compiled) {
            return;
        }

        const fileName = path.basename(this.sourceFilePath);
        const buildCommand: string[] = [
            "/usr/bin/nice",
            "prlimit",
            "--nproc=128",
            "--nofile=2048",
            "--fsize=10000000", // 10MB
            "--rttime="+this.timeout?.toString(),
            "--as="+this.memoryLimit?.toString(),
            "nosocket",
            ...this.runtime.buildCommand.map(arg => arg.replace("{file}", fileName))
        ];

        const buildCommandOutput = await this.executeCommand(buildCommand);
        if (buildCommandOutput.exitCode !== 0) {
            throw new Error(buildCommandOutput.output);
        }

        this.builtFilePath = this.sourceFilePath.replace(`code.${this.runtime.extension}`, "code");
    }

    async run(): Promise<CommandOutput> {
        try {
            let finalFileName: string = path.basename(this.sourceFilePath);
            if (this.runtime.compiled) {
                finalFileName = this.builtFilePath.replace(`.${this.runtime.extension}`, "");
            }

            const runCommand: string[] = [
                "/usr/bin/nice",
                "prlimit",
                "--nproc=64",
                "--nofile=2048",
                "--fsize=10000000", // 10MB
                "--rttime="+this.timeout?.toString(),
                "--as="+this.memoryLimit?.toString(),
                "nosocket",
                ...this.runtime.runCommand.map(
                    arg => arg.replace(
                        "{file}",
                        finalFileName
                    ))
            ];

            const result = await this.executeCommand(runCommand);
            await this.cleanup();
            return result;
        } catch (error) {
            await this.cleanup();
            throw error;
        }
    }

    private async cleanup(): Promise<void> {
        await fs.rm(this.sourceFilePath);
        if (process.env.ENVIRONMENT === "development") {
            console.log(`Cleaned up: ${this.sourceFilePath}`);
        }

        if (this.runtime.compiled) {
            await fs.rm(this.builtFilePath);
            if (process.env.ENVIRONMENT === "development") {
                console.log(`Cleaned up: ${this.builtFilePath}`);
            }
        }
    }

    private executeCommand(command: string[]): Promise<CommandOutput> {
        const { gid, uid, username } = this.user;
        const timeout = this.timeout;

        return new Promise((resolve, reject) => {
            let stdout = "";
            let stderr = "";
            let output = "";
            let exitCode = 0;
            let exitSignal = "";

            const cmd = childProcess.spawn(command[0], command.slice(1), {
                env: {
                    PATH: process.env?.PATH ?? "/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin",
                    LOGGER_TOKEN: "",
                    LOGGER_SERVER_ADDRESS: "",
                    ENVIRONMENT: ""
                },
                cwd: "/code/" + username,
                gid: gid,
                uid: uid,
                timeout: timeout ?? 5_000,
                stdio: "pipe",
                detached: true
            });

            cmd.stdout.on("data", (data) => {
                stdout += data.toString();
                output += data.toString();

                if (process.env.ENVIRONMENT === "development") {
                    console.log(data.toString());
                }
            });

            cmd.stderr.on("data", (data) => {
                stderr += data.toString();
                output += data.toString();

                if (process.env.ENVIRONMENT === "development") {
                    console.log(data.toString());
                }
            });

            cmd.on("error", (error) => {
                cmd.stdout.destroy();
                cmd.stderr.destroy();

                reject(error.message);
            });

            cmd.on("close", (code, signal) => {
                cmd.stdout.destroy();
                cmd.stderr.destroy();

                exitCode = code ?? 0;
                exitSignal = signal ?? "";

                resolve({
                    stdout,
                    stderr,
                    output,
                    exitCode,
                    signal: exitSignal
                });
            });
        });
    }
}
