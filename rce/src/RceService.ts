import { ICodeExecutionEngineService } from "@/stub/rce_pb.grpc-server";
import { sendUnaryData, ServerUnaryCall, UntypedHandleCall } from "@grpc/grpc-js";
import {
  PingResponse,
  Runtimes,
  Runtime,
  CodeRequest,
  CodeResponse
} from "@/stub/rce_pb";
import fs from "fs/promises";
import path from "path";
import toml from "toml";
import { fileURLToPath } from "url";

export class RceServiceImpl implements ICodeExecutionEngineService {
  // eslint-disable-next-line no-undef
  [name: string]: UntypedHandleCall;

  public async listRuntimes(_call, callback: sendUnaryData<Runtimes>) {
    const PACKAGES_DIR = path.join(
      fileURLToPath(import.meta.url),
      "..",
      "..",
      "packages"
    );

    const packages = await fs.readdir(PACKAGES_DIR, {
      withFileTypes: true,
    });
    const runtimesPromise = packages.map(async (p): Promise<Runtime> => {
      const packagePath = path.join(PACKAGES_DIR, p.name);
      const configPath = path.join(packagePath, "config.toml");
      const configFile = await fs.readFile(configPath, "utf8");
      const config = toml.parse(configFile);

      return {
        language: config.language,
        version: config.version,
        compiled: config.compiled,
        aliases: config.aliases,
      };
    });
    const runtimes = await Promise.all(runtimesPromise);

    callback(null, { runtime: runtimes });
  }

  public ping(_call, callback: sendUnaryData<PingResponse>) {
    callback(null, { message: "OK" });
  }

  public execute(call: ServerUnaryCall<CodeRequest, CodeResponse>, callback: sendUnaryData<CodeResponse>) {
    const req = call.request;

    // TODO(elianiva): call execute() with the correct parameter from request
    // eslint-disable-next-line no-console
    console.log(req);

    callback(null, {
      exitCode: 0,
      language: "Javascript",
      output: "Hello World",
      stderr: "",
      stdout: "",
      version: "1.0.0"
    });
  }
}
