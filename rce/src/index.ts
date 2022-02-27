import grpc from "@grpc/grpc-js";
import { codeExecutionEngineServiceDefinition } from "@/stub/rce_pb.grpc-server";
import { RceServiceImpl } from "@/RceService";
import { Logger } from "@/Logger";

const HOST = "0.0.0.0:50051";

const logger = new Logger(
  process.env.LOGGER_SERVER_ADDRESS,
  process.env.LOGGER_TOKEN,
  process.env.ENVIRONMENT
);
const rceServiceImpl = new RceServiceImpl(logger);

const server = new grpc.Server();
server.addService(codeExecutionEngineServiceDefinition, rceServiceImpl);

server.bindAsync(HOST, grpc.ServerCredentials.createInsecure(), () => {
  server.start();
});
