import grpc from "@grpc/grpc-js";
import { codeExecutionEngineServiceDefinition } from "./stub/rce.grpc-server";
import { RceServiceImpl } from "./RceService";

const HOST = "0.0.0.0:50051";

const rceServiceImpl = new RceServiceImpl();
const server = new grpc.Server();
server.addService(codeExecutionEngineServiceDefinition, rceServiceImpl);

server.bindAsync(HOST, grpc.ServerCredentials.createInsecure(), () => {
  server.start();
});
