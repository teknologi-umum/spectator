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

    compile(filePath: string): void {
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
        const buildCommandOutput = this.executeCommand(buildCommand);
        if (buildCommandOutput.exitCode !== 0 || buildCommandOutput.stderr) {
            throw new Error(buildCommandOutput.stderr);
        }
    }

    run(filePath: string): CommandOutput {
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
        return this.executeCommand(runCommand);
    }

    private executeCommand(command: string[]): CommandOutput {
        let stdout = "";
        let stderr = "";
        let output = "";
        let exitCode = 0;
        let exitSignal = "";

        const cmd = childProcess.spawn(command[0], command.slice(1), {
            cwd: "/code/" + this.user.uid.toString(),
            gid: this.user.gid,
            uid: this.user.uid,
            shell: false
        });

        cmd.stdout.on("data", (data) => {
            stdout += data.toString();
        });

        cmd.stderr.on("data", (data) => {
            stderr += data.toString();
        });

        cmd.on("close", (code, signal) => {
            exitCode = code ?? 0;
            exitSignal = signal ?? "";
        });

        if (stdout !== "") {
            output += stdout;
        } else {
            output += stderr;
        }

        return {
            stdout,
            stderr,
            output,
            exitCode,
            signal: exitSignal
        };
    }
}
