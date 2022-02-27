export class Runtime {
    language: string;
    version: string;
    extension: string;
    compiled: boolean;
    buildCommand: string[];
    runCommand: string[];

    constructor(language: string, version: string, extension: string, compiled: boolean, buildCommand: string[], runCommand: string[]) {
        if (!language || !version || !extension || !runCommand) {
            throw new TypeError("Invalid runtime parameters");
        }

        if (compiled && buildCommand.length === 0) {
            throw new TypeError("Invalid runtime parameters: buildCommand is empty yet compiled is true");
        }

        this.language = language;
        this.version = version;
        this.extension = extension;
        this.compiled = compiled;
        this.buildCommand = buildCommand;
        this.runCommand = runCommand;
    }
}
