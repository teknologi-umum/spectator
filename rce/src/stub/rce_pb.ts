// @generated by protobuf-ts 2.2.2 with parameter server_grpc1,client_none,generate_dependencies,optimize_code_size,add_pb_suffix
// @generated from protobuf file "rce.proto" (package "rce", syntax proto3)
// tslint:disable
import { ServiceType } from "@protobuf-ts/runtime-rpc";
import { MessageType } from "@protobuf-ts/runtime";
/**
 * @generated from protobuf message rce.EmptyRequest
 */
export interface EmptyRequest {
}
/**
 * @generated from protobuf message rce.Runtimes
 */
export interface Runtimes {
    /**
     * @generated from protobuf field: repeated rce.Runtime runtime = 1;
     */
    runtime: Runtime[];
}
/**
 * @generated from protobuf message rce.Runtime
 */
export interface Runtime {
    /**
     * @generated from protobuf field: string language = 1;
     */
    language: string;
    /**
     * @generated from protobuf field: string version = 2;
     */
    version: string;
    /**
     * @generated from protobuf field: repeated string aliases = 3;
     */
    aliases: string[];
    /**
     * @generated from protobuf field: bool compiled = 4;
     */
    compiled: boolean;
}
/**
 * @generated from protobuf message rce.PingResponse
 */
export interface PingResponse {
    /**
     * @generated from protobuf field: string message = 1;
     */
    message: string;
}
/**
 * @generated from protobuf message rce.CodeRequest
 */
export interface CodeRequest {
    /**
     * @generated from protobuf field: string language = 1;
     */
    language: string;
    /**
     * @generated from protobuf field: string version = 2;
     */
    version: string;
    /**
     * @generated from protobuf field: string code = 3;
     */
    code: string;
    /**
     * @generated from protobuf field: int32 compile_timeout = 4;
     */
    compileTimeout: number;
    /**
     * @generated from protobuf field: int32 run_timeout = 5;
     */
    runTimeout: number;
    /**
     * @generated from protobuf field: int32 memory_limit = 6;
     */
    memoryLimit: number;
}
/**
 * @generated from protobuf message rce.CodeResponse
 */
export interface CodeResponse {
    /**
     * @generated from protobuf field: string language = 1;
     */
    language: string;
    /**
     * @generated from protobuf field: string version = 2;
     */
    version: string;
    /**
     * @generated from protobuf field: rce.CompileResult compile_result = 3;
     */
    compileResult?: CompileResult;
    /**
     * @generated from protobuf field: rce.RunResult run_result = 4;
     */
    runResult?: RunResult;
}
/**
 * @generated from protobuf message rce.CompileResult
 */
export interface CompileResult {
    /**
     * @generated from protobuf field: string stdout = 1;
     */
    stdout: string;
    /**
     * @generated from protobuf field: string stderr = 2;
     */
    stderr: string;
    /**
     * @generated from protobuf field: string output = 3;
     */
    output: string;
    /**
     * @generated from protobuf field: int32 exitCode = 4;
     */
    exitCode: number;
}
/**
 * @generated from protobuf message rce.RunResult
 */
export interface RunResult {
    /**
     * @generated from protobuf field: string stdout = 1;
     */
    stdout: string;
    /**
     * @generated from protobuf field: string stderr = 2;
     */
    stderr: string;
    /**
     * @generated from protobuf field: string output = 3;
     */
    output: string;
    /**
     * @generated from protobuf field: int32 exitCode = 4;
     */
    exitCode: number;
}
// @generated message type with reflection information, may provide speed optimized methods
class EmptyRequest$Type extends MessageType<EmptyRequest> {
    constructor() {
        super("rce.EmptyRequest", []);
    }
}
/**
 * @generated MessageType for protobuf message rce.EmptyRequest
 */
export const EmptyRequest = new EmptyRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class Runtimes$Type extends MessageType<Runtimes> {
    constructor() {
        super("rce.Runtimes", [
            { no: 1, name: "runtime", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => Runtime }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message rce.Runtimes
 */
export const Runtimes = new Runtimes$Type();
// @generated message type with reflection information, may provide speed optimized methods
class Runtime$Type extends MessageType<Runtime> {
    constructor() {
        super("rce.Runtime", [
            { no: 1, name: "language", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "version", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 3, name: "aliases", kind: "scalar", repeat: 2 /*RepeatType.UNPACKED*/, T: 9 /*ScalarType.STRING*/ },
            { no: 4, name: "compiled", kind: "scalar", T: 8 /*ScalarType.BOOL*/ }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message rce.Runtime
 */
export const Runtime = new Runtime$Type();
// @generated message type with reflection information, may provide speed optimized methods
class PingResponse$Type extends MessageType<PingResponse> {
    constructor() {
        super("rce.PingResponse", [
            { no: 1, name: "message", kind: "scalar", T: 9 /*ScalarType.STRING*/ }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message rce.PingResponse
 */
export const PingResponse = new PingResponse$Type();
// @generated message type with reflection information, may provide speed optimized methods
class CodeRequest$Type extends MessageType<CodeRequest> {
    constructor() {
        super("rce.CodeRequest", [
            { no: 1, name: "language", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "version", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 3, name: "code", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 4, name: "compile_timeout", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 5, name: "run_timeout", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 6, name: "memory_limit", kind: "scalar", T: 5 /*ScalarType.INT32*/ }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message rce.CodeRequest
 */
export const CodeRequest = new CodeRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class CodeResponse$Type extends MessageType<CodeResponse> {
    constructor() {
        super("rce.CodeResponse", [
            { no: 1, name: "language", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "version", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 3, name: "compile_result", kind: "message", T: () => CompileResult },
            { no: 4, name: "run_result", kind: "message", T: () => RunResult }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message rce.CodeResponse
 */
export const CodeResponse = new CodeResponse$Type();
// @generated message type with reflection information, may provide speed optimized methods
class CompileResult$Type extends MessageType<CompileResult> {
    constructor() {
        super("rce.CompileResult", [
            { no: 1, name: "stdout", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "stderr", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 3, name: "output", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 4, name: "exitCode", kind: "scalar", T: 5 /*ScalarType.INT32*/ }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message rce.CompileResult
 */
export const CompileResult = new CompileResult$Type();
// @generated message type with reflection information, may provide speed optimized methods
class RunResult$Type extends MessageType<RunResult> {
    constructor() {
        super("rce.RunResult", [
            { no: 1, name: "stdout", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "stderr", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 3, name: "output", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 4, name: "exitCode", kind: "scalar", T: 5 /*ScalarType.INT32*/ }
        ]);
    }
}
/**
 * @generated MessageType for protobuf message rce.RunResult
 */
export const RunResult = new RunResult$Type();
/**
 * @generated ServiceType for protobuf service rce.CodeExecutionEngineService
 */
export const CodeExecutionEngineService = new ServiceType("rce.CodeExecutionEngineService", [
    { name: "ListRuntimes", options: {}, I: EmptyRequest, O: Runtimes },
    { name: "Execute", options: {}, I: CodeRequest, O: CodeResponse },
    { name: "Ping", options: {}, I: EmptyRequest, O: PingResponse }
]);
