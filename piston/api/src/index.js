#!/usr/bin/env node
require("nocamel");
const Logger = require("logplease");
const express = require("express");
const expressWs = require("express-ws");
const globals = require("./globals");
const config = require("./config");
const path = require("path");
const fs = require("fs/promises");
const fss = require("fs");
const runtime = require("./runtime");
const Package = require("./package");

const logger = Logger.create("index");
const app = express();
expressWs(app);

(async () => {
    logger.info("Setting loglevel to", config.log_level);
    Logger.setLogLevel(config.log_level);
    logger.debug("Ensuring data directories exist");

    Object.values(globals.data_directories).for_each(dir => {
        const data_path = path.join(config.data_directory, dir);

        logger.debug(`Ensuring ${data_path} exists`);

        if (!fss.exists_sync(data_path)) {
            logger.info(`${data_path} does not exist.. Creating..`);

            try {
                fss.mkdir_sync(data_path);
            } catch (e) {
                logger.error(`Failed to create ${data_path}: `, e.message);
            }
        }
    });

    // Install all packages from the /packages directory
    // We fetch the list first
    const pkgList = Package.get_package_list();
    for (const pkg of pkgList) {
        const {language, version} = await pkg.install();
        console.log(`Installed! ${language}-${version}`);
    }

    logger.info("Loading packages");
    const pkgdir = path.join(
        config.data_directory,
        globals.data_directories.packages
    );

    const pkglist = await fs.readdir(pkgdir, { encoding: "utf-8", withFileTypes: true });
    logger.info(`Got pkglist: ${pkglist.map(o => o.name).join(", ")}`);
    const languages = await Promise.all(
        pkglist.map(async lang => {
            if (lang.isDirectory()) {
                const x = await fs.readdir(path.join(pkgdir, lang.name));
                return x.map(y => {
                    logger.info(`[${lang.name}] y is ${y}`);
                    return path.join(pkgdir, lang.name, y);
                });
            }
        })
    );

    const installed_languages = languages
        .flat()
        .filter(pkg =>
            fss.exists_sync(path.join(pkg, globals.pkg_installed_file))
        );

    installed_languages.for_each(pkg => runtime.load_package(pkg));

    logger.info("Starting API Server");
    logger.info("Constructing Express App");
    logger.info("Registering middleware");

    app.use(express.urlencoded({ extended: true }));
    app.use(express.json());

    app.use((err, req, res, next) => {
        return res.status(400).send({
            stack: err.stack
        });
    });

    logger.info("Registering Routes");

    const api_v2 = require("./api/v2");
    app.use("/api/v2", api_v2);
    app.use("/api/v2", api_v2);

    app.use((req, res, next) => {
        return res.status(404).send({ message: "Not Found" });
    });

    logger.info("Calling app.listen");
    const [address, port] = config.bind_address.split(":");

    app.listen(port, address, () => {
        logger.info("API server started on", config.bind_address);
    });
})();
