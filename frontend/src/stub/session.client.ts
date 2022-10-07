// @generated by protobuf-ts 2.5.0
// @generated from protobuf file "session.proto" (package "session", syntax proto3)
// tslint:disable
import type { RpcTransport } from "@protobuf-ts/runtime-rpc";
import type { ServiceInfo } from "@protobuf-ts/runtime-rpc";
import { SessionService } from "./session";
import type { SubmitSolutionSAMRequest } from "./session";
import type { SubmissionResult } from "./session";
import type { SubmissionRequest } from "./session";
import type { ExamResult } from "./session";
import type { Exam } from "./session";
import type { EmptyRequest } from "./session";
import type { SubmitSAMRequest } from "./session";
import type { SubmitPersonalInfoRequest } from "./session";
import type { EmptyReply } from "./session";
import type { SetLocaleRequest } from "./session";
import { stackIntercept } from "@protobuf-ts/runtime-rpc";
import type { SessionReply } from "./session";
import type { StartSessionRequest } from "./session";
import type { UnaryCall } from "@protobuf-ts/runtime-rpc";
import type { RpcOptions } from "@protobuf-ts/runtime-rpc";
/**
 * @generated from protobuf service session.SessionService
 */
export interface ISessionServiceClient {
    /**
     * @generated from protobuf rpc: StartSession(session.StartSessionRequest) returns (session.SessionReply);
     */
    startSession(input: StartSessionRequest, options?: RpcOptions): UnaryCall<StartSessionRequest, SessionReply>;
    /**
     * @generated from protobuf rpc: SetLocale(session.SetLocaleRequest) returns (session.EmptyReply);
     */
    setLocale(input: SetLocaleRequest, options?: RpcOptions): UnaryCall<SetLocaleRequest, EmptyReply>;
    /**
     * @generated from protobuf rpc: SubmitPersonalInfo(session.SubmitPersonalInfoRequest) returns (session.EmptyReply);
     */
    submitPersonalInfo(input: SubmitPersonalInfoRequest, options?: RpcOptions): UnaryCall<SubmitPersonalInfoRequest, EmptyReply>;
    /**
     * @generated from protobuf rpc: SubmitBeforeExamSAM(session.SubmitSAMRequest) returns (session.EmptyReply);
     */
    submitBeforeExamSAM(input: SubmitSAMRequest, options?: RpcOptions): UnaryCall<SubmitSAMRequest, EmptyReply>;
    /**
     * @generated from protobuf rpc: StartExam(session.EmptyRequest) returns (session.Exam);
     */
    startExam(input: EmptyRequest, options?: RpcOptions): UnaryCall<EmptyRequest, Exam>;
    /**
     * @generated from protobuf rpc: ResumeExam(session.EmptyRequest) returns (session.Exam);
     */
    resumeExam(input: EmptyRequest, options?: RpcOptions): UnaryCall<EmptyRequest, Exam>;
    /**
     * @generated from protobuf rpc: EndExam(session.EmptyRequest) returns (session.ExamResult);
     */
    endExam(input: EmptyRequest, options?: RpcOptions): UnaryCall<EmptyRequest, ExamResult>;
    /**
     * @generated from protobuf rpc: SubmitSolution(session.SubmissionRequest) returns (session.SubmissionResult);
     */
    submitSolution(input: SubmissionRequest, options?: RpcOptions): UnaryCall<SubmissionRequest, SubmissionResult>;
    /**
     * @generated from protobuf rpc: SubmitAfterExamSAM(session.SubmitSAMRequest) returns (session.EmptyReply);
     */
    submitAfterExamSAM(input: SubmitSAMRequest, options?: RpcOptions): UnaryCall<SubmitSAMRequest, EmptyReply>;
    /**
     * @generated from protobuf rpc: SubmitSolutionSAM(session.SubmitSolutionSAMRequest) returns (session.EmptyReply);
     */
    submitSolutionSAM(input: SubmitSolutionSAMRequest, options?: RpcOptions): UnaryCall<SubmitSolutionSAMRequest, EmptyReply>;
}
/**
 * @generated from protobuf service session.SessionService
 */
export class SessionServiceClient implements ISessionServiceClient, ServiceInfo {
    typeName = SessionService.typeName;
    methods = SessionService.methods;
    options = SessionService.options;
    constructor(private readonly _transport: RpcTransport) {
    }
    /**
     * @generated from protobuf rpc: StartSession(session.StartSessionRequest) returns (session.SessionReply);
     */
    startSession(input: StartSessionRequest, options?: RpcOptions): UnaryCall<StartSessionRequest, SessionReply> {
        const method = this.methods[0], opt = this._transport.mergeOptions(options);
        return stackIntercept<StartSessionRequest, SessionReply>("unary", this._transport, method, opt, input);
    }
    /**
     * @generated from protobuf rpc: SetLocale(session.SetLocaleRequest) returns (session.EmptyReply);
     */
    setLocale(input: SetLocaleRequest, options?: RpcOptions): UnaryCall<SetLocaleRequest, EmptyReply> {
        const method = this.methods[1], opt = this._transport.mergeOptions(options);
        return stackIntercept<SetLocaleRequest, EmptyReply>("unary", this._transport, method, opt, input);
    }
    /**
     * @generated from protobuf rpc: SubmitPersonalInfo(session.SubmitPersonalInfoRequest) returns (session.EmptyReply);
     */
    submitPersonalInfo(input: SubmitPersonalInfoRequest, options?: RpcOptions): UnaryCall<SubmitPersonalInfoRequest, EmptyReply> {
        const method = this.methods[2], opt = this._transport.mergeOptions(options);
        return stackIntercept<SubmitPersonalInfoRequest, EmptyReply>("unary", this._transport, method, opt, input);
    }
    /**
     * @generated from protobuf rpc: SubmitBeforeExamSAM(session.SubmitSAMRequest) returns (session.EmptyReply);
     */
    submitBeforeExamSAM(input: SubmitSAMRequest, options?: RpcOptions): UnaryCall<SubmitSAMRequest, EmptyReply> {
        const method = this.methods[3], opt = this._transport.mergeOptions(options);
        return stackIntercept<SubmitSAMRequest, EmptyReply>("unary", this._transport, method, opt, input);
    }
    /**
     * @generated from protobuf rpc: StartExam(session.EmptyRequest) returns (session.Exam);
     */
    startExam(input: EmptyRequest, options?: RpcOptions): UnaryCall<EmptyRequest, Exam> {
        const method = this.methods[4], opt = this._transport.mergeOptions(options);
        return stackIntercept<EmptyRequest, Exam>("unary", this._transport, method, opt, input);
    }
    /**
     * @generated from protobuf rpc: ResumeExam(session.EmptyRequest) returns (session.Exam);
     */
    resumeExam(input: EmptyRequest, options?: RpcOptions): UnaryCall<EmptyRequest, Exam> {
        const method = this.methods[5], opt = this._transport.mergeOptions(options);
        return stackIntercept<EmptyRequest, Exam>("unary", this._transport, method, opt, input);
    }
    /**
     * @generated from protobuf rpc: EndExam(session.EmptyRequest) returns (session.ExamResult);
     */
    endExam(input: EmptyRequest, options?: RpcOptions): UnaryCall<EmptyRequest, ExamResult> {
        const method = this.methods[6], opt = this._transport.mergeOptions(options);
        return stackIntercept<EmptyRequest, ExamResult>("unary", this._transport, method, opt, input);
    }
    /**
     * @generated from protobuf rpc: SubmitSolution(session.SubmissionRequest) returns (session.SubmissionResult);
     */
    submitSolution(input: SubmissionRequest, options?: RpcOptions): UnaryCall<SubmissionRequest, SubmissionResult> {
        const method = this.methods[7], opt = this._transport.mergeOptions(options);
        return stackIntercept<SubmissionRequest, SubmissionResult>("unary", this._transport, method, opt, input);
    }
    /**
     * @generated from protobuf rpc: SubmitAfterExamSAM(session.SubmitSAMRequest) returns (session.EmptyReply);
     */
    submitAfterExamSAM(input: SubmitSAMRequest, options?: RpcOptions): UnaryCall<SubmitSAMRequest, EmptyReply> {
        const method = this.methods[8], opt = this._transport.mergeOptions(options);
        return stackIntercept<SubmitSAMRequest, EmptyReply>("unary", this._transport, method, opt, input);
    }
    /**
     * @generated from protobuf rpc: SubmitSolutionSAM(session.SubmitSolutionSAMRequest) returns (session.EmptyReply);
     */
    submitSolutionSAM(input: SubmitSolutionSAMRequest, options?: RpcOptions): UnaryCall<SubmitSolutionSAMRequest, EmptyReply> {
        const method = this.methods[9], opt = this._transport.mergeOptions(options);
        return stackIntercept<SubmitSolutionSAMRequest, EmptyReply>("unary", this._transport, method, opt, input);
    }
}
