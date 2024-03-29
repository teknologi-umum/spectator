// @generated by protobuf-ts 2.2.2 with parameter server_grpc1,client_none,generate_dependencies,optimize_code_size,add_pb_suffix
// @generated from protobuf file "rce.proto" (package "rce", syntax proto3)
// tslint:disable
import { PingResponse } from "./rce_pb";
import { CodeResponse } from "./rce_pb";
import { CodeRequest } from "./rce_pb";
import { Runtimes } from "./rce_pb";
import { EmptyRequest } from "./rce_pb";
import * as grpc from "@grpc/grpc-js";
/**
 * @generated from protobuf service rce.CodeExecutionEngineService
 */
export interface ICodeExecutionEngineService extends grpc.UntypedServiceImplementation {
    /**
     * @generated from protobuf rpc: ListRuntimes(rce.EmptyRequest) returns (rce.Runtimes);
     */
    listRuntimes: grpc.handleUnaryCall<EmptyRequest, Runtimes>;
    /**
     * @generated from protobuf rpc: Execute(rce.CodeRequest) returns (rce.CodeResponse);
     */
    execute: grpc.handleUnaryCall<CodeRequest, CodeResponse>;
    /**
     * @generated from protobuf rpc: Ping(rce.EmptyRequest) returns (rce.PingResponse);
     */
    ping: grpc.handleUnaryCall<EmptyRequest, PingResponse>;
}
/**
 * @grpc/grpc-js definition for the protobuf service rce.CodeExecutionEngineService.
 *
 * Usage: Implement the interface ICodeExecutionEngineService and add to a grpc server.
 *
 * ```typescript
 * const server = new grpc.Server();
 * const service: ICodeExecutionEngineService = ...
 * server.addService(codeExecutionEngineServiceDefinition, service);
 * ```
 */
export const codeExecutionEngineServiceDefinition: grpc.ServiceDefinition<ICodeExecutionEngineService> = {
    listRuntimes: {
        path: "/rce.CodeExecutionEngineService/ListRuntimes",
        originalName: "ListRuntimes",
        requestStream: false,
        responseStream: false,
        responseDeserialize: bytes => Runtimes.fromBinary(bytes),
        requestDeserialize: bytes => EmptyRequest.fromBinary(bytes),
        responseSerialize: value => Buffer.from(Runtimes.toBinary(value)),
        requestSerialize: value => Buffer.from(EmptyRequest.toBinary(value))
    },
    execute: {
        path: "/rce.CodeExecutionEngineService/Execute",
        originalName: "Execute",
        requestStream: false,
        responseStream: false,
        responseDeserialize: bytes => CodeResponse.fromBinary(bytes),
        requestDeserialize: bytes => CodeRequest.fromBinary(bytes),
        responseSerialize: value => Buffer.from(CodeResponse.toBinary(value)),
        requestSerialize: value => Buffer.from(CodeRequest.toBinary(value))
    },
    ping: {
        path: "/rce.CodeExecutionEngineService/Ping",
        originalName: "Ping",
        requestStream: false,
        responseStream: false,
        responseDeserialize: bytes => PingResponse.fromBinary(bytes),
        requestDeserialize: bytes => EmptyRequest.fromBinary(bytes),
        responseSerialize: value => Buffer.from(PingResponse.toBinary(value)),
        requestSerialize: value => Buffer.from(EmptyRequest.toBinary(value))
    }
};
