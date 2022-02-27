import { ICodeExecutionEngineService } from "@/stub/rce.grpc-server";
import { sendUnaryData } from "@grpc/grpc-js";
import { PingResponse } from "@/stub/rce";

export class RceServiceImpl implements ICodeExecutionEngineService {
  public listRuntimes() {
    // TODO: implementation
  }

  public ping(_call, callback: sendUnaryData<PingResponse>) {
    callback(null, { message: "OK" });
  }

  public execute() {
    // TODO: implementation
  }
}
