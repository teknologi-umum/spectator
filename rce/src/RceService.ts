import { ICodeExecutionEngineService } from "@/stub/rce.grpc-server";
import { sendUnaryData, UntypedHandleCall } from "@grpc/grpc-js";
import { PingResponse } from "@/stub/rce";

export class RceServiceImpl implements ICodeExecutionEngineService {
  [name: string]: UntypedHandleCall;

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
