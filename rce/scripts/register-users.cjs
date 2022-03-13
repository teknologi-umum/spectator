#!/usr/bin/env node
/* eslint-disable @typescript-eslint/no-var-requires */

const fs = require("fs/promises");
const cp = require("child_process");
const console = require("console");

// This file should be in CommonJS as it will be called directly.


function execute(command, workingDirectory = process.cwd()) {
    return new Promise((resolve, reject) => {
        const cmd = cp.exec(
            command,
            { cwd: workingDirectory },
            (error) => {
                if (error) {
                    reject(error);
                }
            }
        );

        let stdout = "";
        let stderr = "";

        cmd.stdout.on("data", (data) => {
            console.log(data.toString());
            stdout += data.toString();
        });

        cmd.stderr.on("data", (data) => {
            console.error(data.toString());
            stderr += data.toString();
        });

        cmd.on("close", (code) => {
            if (code !== 0) {
                reject(new Error(stderr));
                return;
            }

            resolve(stdout);
        });
    });
}

const globalGroupID = 64101;

(async () => {
    await fs.mkdir("/code", { recursive: true });

    // Create a new group
    const groupAddStdout = await execute(`groupadd -g ${globalGroupID.toString()} code_executors`);
    console.log(groupAddStdout.toString());

    const ints = Array.from(Array(50).keys());

    for await (const i of ints) {
        const uid = (globalGroupID + i).toString();
        const homeDir = `/code/code_executor_${uid}`;
        const stdout = await execute(`useradd -M --base-dir ${homeDir} --uid ${uid} --gid ${globalGroupID.toString()} --shell /bin/bash --home ${homeDir} --comment "Code executor ${uid}" code_executor_${uid}`);
        console.log(stdout);
        await fs.mkdir(homeDir, { recursive: true });
    }
})();
