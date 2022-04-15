// @generated by protobuf-ts 2.1.0
// @generated from protobuf file "events.proto" (package "events", syntax proto3)
// tslint:disable
import type { RpcTransport } from "@protobuf-ts/runtime-rpc";
import type { ServiceInfo } from "@protobuf-ts/runtime-rpc";
import { EventsService } from "./events";
import type { KeystrokeRequest } from "./events";
import type { MouseScrollRequest } from "./events";
import type { MouseMoveRequest } from "./events";
import { stackIntercept } from "@protobuf-ts/runtime-rpc";
import type { EventReply } from "./events";
import type { MouseClickRequest } from "./events";
import type { UnaryCall } from "@protobuf-ts/runtime-rpc";
import type { RpcOptions } from "@protobuf-ts/runtime-rpc";
/**
 * @generated from protobuf service events.EventsService
 */
export interface IEventsServiceClient {
    /**
     * @generated from protobuf rpc: MouseClick(events.MouseClickRequest) returns (events.EventReply);
     */
    mouseClick(input: MouseClickRequest, options?: RpcOptions): UnaryCall<MouseClickRequest, EventReply>;
    /**
     * @generated from protobuf rpc: MouseMove(events.MouseMoveRequest) returns (events.EventReply);
     */
    mouseMove(input: MouseMoveRequest, options?: RpcOptions): UnaryCall<MouseMoveRequest, EventReply>;
    /**
     * @generated from protobuf rpc: MouseScroll(events.MouseScrollRequest) returns (events.EventReply);
     */
    mouseScroll(input: MouseScrollRequest, options?: RpcOptions): UnaryCall<MouseScrollRequest, EventReply>;
    /**
     * @generated from protobuf rpc: Keystroke(events.KeystrokeRequest) returns (events.EventReply);
     */
    keystroke(input: KeystrokeRequest, options?: RpcOptions): UnaryCall<KeystrokeRequest, EventReply>;
}
/**
 * @generated from protobuf service events.EventsService
 */
export class EventsServiceClient implements IEventsServiceClient, ServiceInfo {
    typeName = EventsService.typeName;
    methods = EventsService.methods;
    options = EventsService.options;
    constructor(private readonly _transport: RpcTransport) {
    }
    /**
     * @generated from protobuf rpc: MouseClick(events.MouseClickRequest) returns (events.EventReply);
     */
    mouseClick(input: MouseClickRequest, options?: RpcOptions): UnaryCall<MouseClickRequest, EventReply> {
        const method = this.methods[0], opt = this._transport.mergeOptions(options);
        return stackIntercept<MouseClickRequest, EventReply>("unary", this._transport, method, opt, input);
    }
    /**
     * @generated from protobuf rpc: MouseMove(events.MouseMoveRequest) returns (events.EventReply);
     */
    mouseMove(input: MouseMoveRequest, options?: RpcOptions): UnaryCall<MouseMoveRequest, EventReply> {
        const method = this.methods[1], opt = this._transport.mergeOptions(options);
        return stackIntercept<MouseMoveRequest, EventReply>("unary", this._transport, method, opt, input);
    }
    /**
     * @generated from protobuf rpc: MouseScroll(events.MouseScrollRequest) returns (events.EventReply);
     */
    mouseScroll(input: MouseScrollRequest, options?: RpcOptions): UnaryCall<MouseScrollRequest, EventReply> {
        const method = this.methods[2], opt = this._transport.mergeOptions(options);
        return stackIntercept<MouseScrollRequest, EventReply>("unary", this._transport, method, opt, input);
    }
    /**
     * @generated from protobuf rpc: Keystroke(events.KeystrokeRequest) returns (events.EventReply);
     */
    keystroke(input: KeystrokeRequest, options?: RpcOptions): UnaryCall<KeystrokeRequest, EventReply> {
        const method = this.methods[3], opt = this._transport.mergeOptions(options);
        return stackIntercept<KeystrokeRequest, EventReply>("unary", this._transport, method, opt, input);
    }
}
