#!/usr/bin/env node
/* eslint-disable @typescript-eslint/no-var-requires */

const fs = require("fs");
const cp = require("child_process");

// This file should be in CommonJS as it will be called directly.

const globalGroupID = 64101;

fs.mkdirSync("/code", { recursive: true });

// Create a new group
cp.execSync(`groupadd -g ${globalGroupID.toString()} code_executors`);
for (let i = 0; i < 50; i++) {
    const uid = (globalGroupID + i).toString();
    const homeDir = `/code/code_executor_${uid}`;
    cp.execSync(`useradd -M --base-dir ${homeDir} --uid ${uid} --gid ${globalGroupID.toString()} --shell /bin/bash --home ${homeDir} --comment "Code executor ${uid}" code_executor_${uid}`);
}
