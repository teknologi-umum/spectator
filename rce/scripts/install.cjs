#!/usr/bin/env node
/* eslint-disable @typescript-eslint/no-var-requires */

const fs = require("fs/promises");
const cp = require("child_process");
const path = require("path");
const console = require("console");
const toml = require("toml");

// This file should be in CommonJS as it will be called directly.

function execute(command, workingDirectory = process.cwd()) {
    return new Promise((resolve, reject) => {
        cp.exec(
            command,
            { cwd: workingDirectory },
            (error, stdout, stderr) => {
                if (error) {
                    reject(error);
                    return;
                }

                if (stderr) {
                    reject(stderr);
                    return;
                }

                resolve(stdout);
            }
        );
    });
}

(async () => {
    const packages = await fs.readdir(path.join(__dirname, "..", "packages"), { withFileTypes: true });
    for await (const pkg of packages) {
        if (!pkg.isDirectory()) {
            continue;
        }

        const packagePath = path.join(__dirname, "..", "packages", pkg.name);
        const installResult = await execute(`./packages/${pkg.name}/install.sh`);
        console.log(installResult);

        const configPath = path.join(packagePath, "config.toml");
        const configFile = await fs.readFile(configPath, "utf8");
        const config = toml.parse(configFile);

        // Run the Hello World file.
        if (config.compiled) {
            const compiled = await execute(config["build_command"].join(" ").replace("{file}", config["test_file"]), packagePath);
            console.log(compiled);

            // Run the test file.
            const testResult = await execute(config["run_command"].join(" "), packagePath);
            console.log(testResult);

            if (testResult.trim() !== "Hello world!") {
                throw new Error(`Test file failed for package ${pkg.name}. Expecting "Hello world!", got "${testResult.trim()}"`);
            }
        } else {
            const testResult = await execute(config["run_command"].join(" ").replace("{file}", config["test_file"]), packagePath);
            console.log(testResult);

            if (testResult.trim() !== "Hello world!") {
                throw new Error(`Test file failed for package ${pkg.name}. Expecting "Hello world!", got "${testResult.trim()}"`);
            }
        }

        console.log(`Package ${pkg.name} installed successfully.`);
    }
})();
