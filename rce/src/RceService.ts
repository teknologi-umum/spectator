import fs from "fs/promises";
import path from "path";
import toml from "toml";
import { fileURLToPath } from "url";
import { ICodeExecutionEngineService } from "@/stub/rce_pb.grpc-server";
import { sendUnaryData, ServerUnaryCall, UntypedHandleCall } from "@grpc/grpc-js";
import {
  PingResponse,
  Runtimes,
  Runtime,
  CodeRequest,
  CodeResponse
} from "@/stub/rce_pb";
import { Runtime as RceRuntime } from "@/runtime/runtime";
import { SystemUsers } from "./user/user";
import { Job } from "./job/job";

export class RceServiceImpl implements ICodeExecutionEngineService {
  // eslint-disable-next-line no-undef
  [name: string]: UntypedHandleCall;
  // TODO(aldy505): To elianiva, help please.
  registeredRuntimes: RceRuntime[];
  users: SystemUsers;

  constructor(registeredRuntimes: RceRuntime[], users: SystemUsers) {
    this.registeredRuntimes = registeredRuntimes;
    this.users = users;
  }

  public async listRuntimes(_call, callback: sendUnaryData<Runtimes>) {
    // TODO(aldy505): To elianiva, please acquire the data from this.registeredRuntimes
    const PACKAGES_DIR = path.join(
      fileURLToPath(import.meta.url),
      "..",
      "..",
      "packages"
    );

    const packages = await fs.readdir(PACKAGES_DIR, {
      withFileTypes: true
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
        aliases: config.aliases
      };
    });
    const runtimes = await Promise.all(runtimesPromise);

    callback(null, { runtime: runtimes });
  }

  public ping(_call, callback: sendUnaryData<PingResponse>) {
    callback(null, { message: "OK" });
  }

  public async execute(call: ServerUnaryCall<CodeRequest, CodeResponse>, callback: sendUnaryData<CodeResponse>) {
    // TODO(aldy505): should add try catch?
    const req = call.request;

    // TODO: validate if the runtime is supported, then we acquire the runtime.
    const runtimeIndex = this.registeredRuntimes.findIndex(r => r.language === req.language && r.version === req.version);
    if (runtimeIndex < 0) {
      callback(new Error("Runtime not found"), null);
      return;
    }

    const runtime = this.registeredRuntimes[runtimeIndex];

    // Acquire the available user.
    const user = this.users.acquire();
    if (user === null) {
      callback(new Error("No user available"), null);
      return;
    }

    // Create a job.
    const job = new Job(user, runtime, req.code, req.compileTimeout);
    const filePath = await job.createFile();
    if (runtime.compiled) {
      job.compile(filePath);
    }

    const commandOutput = job.run(filePath);
    // Release the user.
    this.users.release(user.uid);

    callback(null, {
      exitCode: commandOutput.exitCode,
      language: runtime.language,
      output: commandOutput.output,
      stderr: commandOutput.stderr,
      stdout: commandOutput.stdout,
      version: runtime.version
    });
  }
}
