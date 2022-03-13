import path from "path";
import fs from "fs/promises";
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
    constructor(
        public user: User,
        public runtime: Runtime,
        public code: string,
        public timeout?: number,
        public memoryLimit?: number
    ) {
        if (!user || !runtime || !code) {
            throw new TypeError("Invalid job parameters");
        }

        if (!timeout || timeout < 1) {
            this.timeout = 5_000;
        }

        if (!memoryLimit || memoryLimit < 1) {
            this.memoryLimit = 128 * 1024 * 1024;
        }
    }

    async createFile(): Promise<string> {
        const filePath = path.resolve("/code", `/${this.user.username}`, `/code.${this.runtime.extension}`);
        await fs.writeFile(filePath, this.code);
        await fs.chmod(filePath, 0o555);
        await fs.chown(filePath, this.user.uid, this.user.gid);
        return filePath;
    }

    async compile(filePath: string): Promise<void> {
        if (!this.runtime.compiled) {
            return;
        }

        const fileName = path.basename(filePath);
        const buildCommand: string[] = [
            "nice",
            "prlimit",
            "--nproc=128",
            "--nofile=2048",
            "--fsize=10000000", // 10MB
            "--rttime="+this.timeout?.toString(),
            "--as="+this.memoryLimit?.toString(),
            "nosocket",
            "runuser",
            "-u",
            this.user.username,
            "--",
            ...this.runtime.buildCommand.map(arg => arg.replace("{file}", fileName))
        ];
        const buildCommandOutput = await this.executeCommand(buildCommand);
        if (buildCommandOutput.exitCode !== 0 || buildCommandOutput.stderr) {
            throw new Error(buildCommandOutput.stderr);
        }
    }

    async run(filePath: string): Promise<CommandOutput> {
        const fileName = path.basename(filePath);
        const runCommand: string[] = [
            "nice",
            "prlimit",
            "--nproc=64",
            "--nofile=2048",
            "--fsize=10000000", // 10MB
            "--rttime="+this.timeout?.toString(),
            "--as="+this.memoryLimit?.toString(),
            "nosocket",
            "runuser",
            "-u",
            this.user.username,
            "--",
            ...this.runtime.runCommand.map(
                arg => arg.replace(
                    "{file}",
                    fileName.replace(`.${this.runtime.extension}`, "")
                ))
        ];
        const result = await this.executeCommand(runCommand);
        return result;
    }

    private executeCommand(command: string[]): Promise<CommandOutput> {
        const { gid, uid } = this.user;
        const timeout = this.timeout;

        return new Promise((resolve, reject) => {
            let stdout = "";
            let stderr = "";
            let output = "";
            let exitCode = 0;
            let exitSignal = "";

            const cmd = childProcess.exec(command.join(" "), {
                cwd: "/code/" + uid.toString(),
                gid: gid,
                uid: uid,
                timeout: timeout ?? 5_000,
                encoding: "utf8",
                shell: "/bin/bash"
            }, (error) => {
                if (error !== undefined && error !== null) {
                    stderr = error.message;
                    exitCode = 1;
                    reject(error.message);
                }
            });

            if (cmd.stdout !== null) {
                cmd.stdout.on("data", (data) => {
                    stdout += data.toString();
                });
            }

            if (cmd.stderr !== null) {
                cmd.stderr.on("data", (data) => {
                    stderr += data.toString();
                });
            }

            cmd.on("close", (code, signal) => {
                exitCode = code ?? 0;
                exitSignal = signal ?? "";

                if (stdout !== "") {
                    output += stdout;
                } else {
                    output += stderr;
                }

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
