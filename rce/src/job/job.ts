import path from "path";
import fs from "fs/promises";
import childProcess from "child_process";
import { Runtime } from "@/runtime/runtime";
import { User } from "@/user/user";

export interface JobPrerequisites {
    user: User
    runtime: Runtime
    code: string
    timeout: number
    memoryLimit: number
}

export interface CommandOutput {
    stdout: string
    stderr: string
    output: string
    exitCode: number
    signal: string
}

export class Job implements JobPrerequisites {
    user: User;
    runtime: Runtime;
    code: string;
    timeout: number;
    memoryLimit: number;

    constructor(user: User, runtime: Runtime, code: string, timeout: number, memoryLimit: number) {
        this.user = user;
        this.runtime = runtime;
        this.code = code;
        this.timeout = timeout;
        this.memoryLimit = memoryLimit;
    }

    async createFile(): Promise<string> {
        const filePath = path.resolve("/code", `/${this.user.uid.toString()}`, `/code.${this.runtime.extension}`);
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
        // TODO: append nice and prlimit to the runCommand below
        // TODO: run as user. see https://www.cyberciti.biz/open-source/command-line-hacks/linux-run-command-as-different-user/
        const buildCommand = this.runtime.buildCommand.map(arg => arg.replace("{file}", fileName));
        const buildCommandOutput = this.executeCommand(buildCommand);
        if (buildCommandOutput.exitCode !== 0 || buildCommandOutput.stderr) {
            throw new Error(buildCommandOutput.stderr);
        }
    }

    run(filePath: string): CommandOutput {
        const fileName = path.basename(filePath);
        // TODO: append nice and prlimit to the runCommand below
        // TODO: run as user. see https://www.cyberciti.biz/open-source/command-line-hacks/linux-run-command-as-different-user/
        const runCommand = this.runtime.runCommand.map(
            arg => arg.replace(
                "{file}",
                fileName.replace(`.${this.runtime.extension}`, "")
            )
        );
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

        if (stdout) {
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
