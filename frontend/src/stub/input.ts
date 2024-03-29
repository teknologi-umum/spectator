// @generated by protobuf-ts 2.5.0
// @generated from protobuf file "input.proto" (package "input", syntax proto3)
// tslint:disable
import { ServiceType } from "@protobuf-ts/runtime-rpc";
import type { BinaryWriteOptions } from "@protobuf-ts/runtime";
import type { IBinaryWriter } from "@protobuf-ts/runtime";
import { WireType } from "@protobuf-ts/runtime";
import type { BinaryReadOptions } from "@protobuf-ts/runtime";
import type { IBinaryReader } from "@protobuf-ts/runtime";
import { UnknownFieldHandler } from "@protobuf-ts/runtime";
import type { PartialMessage } from "@protobuf-ts/runtime";
import { reflectionMergePartial } from "@protobuf-ts/runtime";
import { MESSAGE_TYPE } from "@protobuf-ts/runtime";
import { MessageType } from "@protobuf-ts/runtime";
import { MouseDirection } from "./enums";
import { MouseButton } from "./enums";
/**
 * @generated from protobuf message input.MouseClickInfo
 */
export interface MouseClickInfo {
    /**
     * @generated from protobuf field: string access_token = 1;
     */
    accessToken: string;
    /**
     * @generated from protobuf field: int32 question_number = 2;
     */
    questionNumber: number;
    /**
     * @generated from protobuf field: int32 x = 3;
     */
    x: number;
    /**
     * @generated from protobuf field: int32 y = 4;
     */
    y: number;
    /**
     * @generated from protobuf field: enums.MouseButton button = 5;
     */
    button: MouseButton;
    /**
     * @generated from protobuf field: int64 time = 6;
     */
    time: bigint;
}
/**
 * @generated from protobuf message input.MouseMoveInfo
 */
export interface MouseMoveInfo {
    /**
     * @generated from protobuf field: string access_token = 1;
     */
    accessToken: string;
    /**
     * @generated from protobuf field: int32 question_number = 2;
     */
    questionNumber: number;
    /**
     * @generated from protobuf field: int32 x = 3;
     */
    x: number;
    /**
     * @generated from protobuf field: int32 y = 4;
     */
    y: number;
    /**
     * @generated from protobuf field: enums.MouseDirection direction = 5;
     */
    direction: MouseDirection;
    /**
     * @generated from protobuf field: int64 time = 6;
     */
    time: bigint;
}
/**
 * @generated from protobuf message input.MouseScrollInfo
 */
export interface MouseScrollInfo {
    /**
     * @generated from protobuf field: string access_token = 1;
     */
    accessToken: string;
    /**
     * @generated from protobuf field: int32 question_number = 2;
     */
    questionNumber: number;
    /**
     * @generated from protobuf field: int32 x = 3;
     */
    x: number;
    /**
     * @generated from protobuf field: int32 y = 4;
     */
    y: number;
    /**
     * @generated from protobuf field: int32 delta = 5;
     */
    delta: number;
    /**
     * @generated from protobuf field: int64 time = 6;
     */
    time: bigint;
}
/**
 * @generated from protobuf message input.KeystrokeInfo
 */
export interface KeystrokeInfo {
    /**
     * @generated from protobuf field: string access_token = 1;
     */
    accessToken: string;
    /**
     * @generated from protobuf field: int32 question_number = 2;
     */
    questionNumber: number;
    /**
     * @generated from protobuf field: string key_char = 3;
     */
    keyChar: string;
    /**
     * @generated from protobuf field: bool shift = 4;
     */
    shift: boolean;
    /**
     * @generated from protobuf field: bool alt = 5;
     */
    alt: boolean;
    /**
     * @generated from protobuf field: bool control = 6;
     */
    control: boolean;
    /**
     * @generated from protobuf field: bool meta = 7;
     */
    meta: boolean;
    /**
     * @generated from protobuf field: bool unrelated_key = 8;
     */
    unrelatedKey: boolean;
    /**
     * @generated from protobuf field: int64 time = 9;
     */
    time: bigint;
}
/**
 * @generated from protobuf message input.WindowSizeInfo
 */
export interface WindowSizeInfo {
    /**
     * @generated from protobuf field: string access_token = 1;
     */
    accessToken: string;
    /**
     * @generated from protobuf field: int32 question_number = 2;
     */
    questionNumber: number;
    /**
     * @generated from protobuf field: int32 width = 3;
     */
    width: number;
    /**
     * @generated from protobuf field: int32 height = 4;
     */
    height: number;
    /**
     * @generated from protobuf field: int64 time = 5;
     */
    time: bigint;
}
/**
 * @generated from protobuf message input.EventReply
 */
export interface EventReply {
    /**
     * @generated from protobuf field: string message = 1;
     */
    message: string;
}
// @generated message type with reflection information, may provide speed optimized methods
class MouseClickInfo$Type extends MessageType<MouseClickInfo> {
    constructor() {
        super("input.MouseClickInfo", [
            { no: 1, name: "access_token", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "question_number", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 3, name: "x", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 4, name: "y", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 5, name: "button", kind: "enum", T: () => ["enums.MouseButton", MouseButton] },
            { no: 6, name: "time", kind: "scalar", T: 3 /*ScalarType.INT64*/, L: 0 /*LongType.BIGINT*/ }
        ]);
    }
    create(value?: PartialMessage<MouseClickInfo>): MouseClickInfo {
        const message = { accessToken: "", questionNumber: 0, x: 0, y: 0, button: 0, time: 0n };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<MouseClickInfo>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: MouseClickInfo): MouseClickInfo {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* string access_token */ 1:
                    message.accessToken = reader.string();
                    break;
                case /* int32 question_number */ 2:
                    message.questionNumber = reader.int32();
                    break;
                case /* int32 x */ 3:
                    message.x = reader.int32();
                    break;
                case /* int32 y */ 4:
                    message.y = reader.int32();
                    break;
                case /* enums.MouseButton button */ 5:
                    message.button = reader.int32();
                    break;
                case /* int64 time */ 6:
                    message.time = reader.int64().toBigInt();
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message: MouseClickInfo, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* string access_token = 1; */
        if (message.accessToken !== "")
            writer.tag(1, WireType.LengthDelimited).string(message.accessToken);
        /* int32 question_number = 2; */
        if (message.questionNumber !== 0)
            writer.tag(2, WireType.Varint).int32(message.questionNumber);
        /* int32 x = 3; */
        if (message.x !== 0)
            writer.tag(3, WireType.Varint).int32(message.x);
        /* int32 y = 4; */
        if (message.y !== 0)
            writer.tag(4, WireType.Varint).int32(message.y);
        /* enums.MouseButton button = 5; */
        if (message.button !== 0)
            writer.tag(5, WireType.Varint).int32(message.button);
        /* int64 time = 6; */
        if (message.time !== 0n)
            writer.tag(6, WireType.Varint).int64(message.time);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message input.MouseClickInfo
 */
export const MouseClickInfo = new MouseClickInfo$Type();
// @generated message type with reflection information, may provide speed optimized methods
class MouseMoveInfo$Type extends MessageType<MouseMoveInfo> {
    constructor() {
        super("input.MouseMoveInfo", [
            { no: 1, name: "access_token", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "question_number", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 3, name: "x", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 4, name: "y", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 5, name: "direction", kind: "enum", T: () => ["enums.MouseDirection", MouseDirection] },
            { no: 6, name: "time", kind: "scalar", T: 3 /*ScalarType.INT64*/, L: 0 /*LongType.BIGINT*/ }
        ]);
    }
    create(value?: PartialMessage<MouseMoveInfo>): MouseMoveInfo {
        const message = { accessToken: "", questionNumber: 0, x: 0, y: 0, direction: 0, time: 0n };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<MouseMoveInfo>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: MouseMoveInfo): MouseMoveInfo {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* string access_token */ 1:
                    message.accessToken = reader.string();
                    break;
                case /* int32 question_number */ 2:
                    message.questionNumber = reader.int32();
                    break;
                case /* int32 x */ 3:
                    message.x = reader.int32();
                    break;
                case /* int32 y */ 4:
                    message.y = reader.int32();
                    break;
                case /* enums.MouseDirection direction */ 5:
                    message.direction = reader.int32();
                    break;
                case /* int64 time */ 6:
                    message.time = reader.int64().toBigInt();
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message: MouseMoveInfo, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* string access_token = 1; */
        if (message.accessToken !== "")
            writer.tag(1, WireType.LengthDelimited).string(message.accessToken);
        /* int32 question_number = 2; */
        if (message.questionNumber !== 0)
            writer.tag(2, WireType.Varint).int32(message.questionNumber);
        /* int32 x = 3; */
        if (message.x !== 0)
            writer.tag(3, WireType.Varint).int32(message.x);
        /* int32 y = 4; */
        if (message.y !== 0)
            writer.tag(4, WireType.Varint).int32(message.y);
        /* enums.MouseDirection direction = 5; */
        if (message.direction !== 0)
            writer.tag(5, WireType.Varint).int32(message.direction);
        /* int64 time = 6; */
        if (message.time !== 0n)
            writer.tag(6, WireType.Varint).int64(message.time);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message input.MouseMoveInfo
 */
export const MouseMoveInfo = new MouseMoveInfo$Type();
// @generated message type with reflection information, may provide speed optimized methods
class MouseScrollInfo$Type extends MessageType<MouseScrollInfo> {
    constructor() {
        super("input.MouseScrollInfo", [
            { no: 1, name: "access_token", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "question_number", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 3, name: "x", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 4, name: "y", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 5, name: "delta", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 6, name: "time", kind: "scalar", T: 3 /*ScalarType.INT64*/, L: 0 /*LongType.BIGINT*/ }
        ]);
    }
    create(value?: PartialMessage<MouseScrollInfo>): MouseScrollInfo {
        const message = { accessToken: "", questionNumber: 0, x: 0, y: 0, delta: 0, time: 0n };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<MouseScrollInfo>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: MouseScrollInfo): MouseScrollInfo {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* string access_token */ 1:
                    message.accessToken = reader.string();
                    break;
                case /* int32 question_number */ 2:
                    message.questionNumber = reader.int32();
                    break;
                case /* int32 x */ 3:
                    message.x = reader.int32();
                    break;
                case /* int32 y */ 4:
                    message.y = reader.int32();
                    break;
                case /* int32 delta */ 5:
                    message.delta = reader.int32();
                    break;
                case /* int64 time */ 6:
                    message.time = reader.int64().toBigInt();
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message: MouseScrollInfo, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* string access_token = 1; */
        if (message.accessToken !== "")
            writer.tag(1, WireType.LengthDelimited).string(message.accessToken);
        /* int32 question_number = 2; */
        if (message.questionNumber !== 0)
            writer.tag(2, WireType.Varint).int32(message.questionNumber);
        /* int32 x = 3; */
        if (message.x !== 0)
            writer.tag(3, WireType.Varint).int32(message.x);
        /* int32 y = 4; */
        if (message.y !== 0)
            writer.tag(4, WireType.Varint).int32(message.y);
        /* int32 delta = 5; */
        if (message.delta !== 0)
            writer.tag(5, WireType.Varint).int32(message.delta);
        /* int64 time = 6; */
        if (message.time !== 0n)
            writer.tag(6, WireType.Varint).int64(message.time);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message input.MouseScrollInfo
 */
export const MouseScrollInfo = new MouseScrollInfo$Type();
// @generated message type with reflection information, may provide speed optimized methods
class KeystrokeInfo$Type extends MessageType<KeystrokeInfo> {
    constructor() {
        super("input.KeystrokeInfo", [
            { no: 1, name: "access_token", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "question_number", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 3, name: "key_char", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 4, name: "shift", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 5, name: "alt", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 6, name: "control", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 7, name: "meta", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 8, name: "unrelated_key", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 9, name: "time", kind: "scalar", T: 3 /*ScalarType.INT64*/, L: 0 /*LongType.BIGINT*/ }
        ]);
    }
    create(value?: PartialMessage<KeystrokeInfo>): KeystrokeInfo {
        const message = { accessToken: "", questionNumber: 0, keyChar: "", shift: false, alt: false, control: false, meta: false, unrelatedKey: false, time: 0n };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<KeystrokeInfo>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: KeystrokeInfo): KeystrokeInfo {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* string access_token */ 1:
                    message.accessToken = reader.string();
                    break;
                case /* int32 question_number */ 2:
                    message.questionNumber = reader.int32();
                    break;
                case /* string key_char */ 3:
                    message.keyChar = reader.string();
                    break;
                case /* bool shift */ 4:
                    message.shift = reader.bool();
                    break;
                case /* bool alt */ 5:
                    message.alt = reader.bool();
                    break;
                case /* bool control */ 6:
                    message.control = reader.bool();
                    break;
                case /* bool meta */ 7:
                    message.meta = reader.bool();
                    break;
                case /* bool unrelated_key */ 8:
                    message.unrelatedKey = reader.bool();
                    break;
                case /* int64 time */ 9:
                    message.time = reader.int64().toBigInt();
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message: KeystrokeInfo, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* string access_token = 1; */
        if (message.accessToken !== "")
            writer.tag(1, WireType.LengthDelimited).string(message.accessToken);
        /* int32 question_number = 2; */
        if (message.questionNumber !== 0)
            writer.tag(2, WireType.Varint).int32(message.questionNumber);
        /* string key_char = 3; */
        if (message.keyChar !== "")
            writer.tag(3, WireType.LengthDelimited).string(message.keyChar);
        /* bool shift = 4; */
        if (message.shift !== false)
            writer.tag(4, WireType.Varint).bool(message.shift);
        /* bool alt = 5; */
        if (message.alt !== false)
            writer.tag(5, WireType.Varint).bool(message.alt);
        /* bool control = 6; */
        if (message.control !== false)
            writer.tag(6, WireType.Varint).bool(message.control);
        /* bool meta = 7; */
        if (message.meta !== false)
            writer.tag(7, WireType.Varint).bool(message.meta);
        /* bool unrelated_key = 8; */
        if (message.unrelatedKey !== false)
            writer.tag(8, WireType.Varint).bool(message.unrelatedKey);
        /* int64 time = 9; */
        if (message.time !== 0n)
            writer.tag(9, WireType.Varint).int64(message.time);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message input.KeystrokeInfo
 */
export const KeystrokeInfo = new KeystrokeInfo$Type();
// @generated message type with reflection information, may provide speed optimized methods
class WindowSizeInfo$Type extends MessageType<WindowSizeInfo> {
    constructor() {
        super("input.WindowSizeInfo", [
            { no: 1, name: "access_token", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "question_number", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 3, name: "width", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 4, name: "height", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 5, name: "time", kind: "scalar", T: 3 /*ScalarType.INT64*/, L: 0 /*LongType.BIGINT*/ }
        ]);
    }
    create(value?: PartialMessage<WindowSizeInfo>): WindowSizeInfo {
        const message = { accessToken: "", questionNumber: 0, width: 0, height: 0, time: 0n };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<WindowSizeInfo>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: WindowSizeInfo): WindowSizeInfo {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* string access_token */ 1:
                    message.accessToken = reader.string();
                    break;
                case /* int32 question_number */ 2:
                    message.questionNumber = reader.int32();
                    break;
                case /* int32 width */ 3:
                    message.width = reader.int32();
                    break;
                case /* int32 height */ 4:
                    message.height = reader.int32();
                    break;
                case /* int64 time */ 5:
                    message.time = reader.int64().toBigInt();
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message: WindowSizeInfo, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* string access_token = 1; */
        if (message.accessToken !== "")
            writer.tag(1, WireType.LengthDelimited).string(message.accessToken);
        /* int32 question_number = 2; */
        if (message.questionNumber !== 0)
            writer.tag(2, WireType.Varint).int32(message.questionNumber);
        /* int32 width = 3; */
        if (message.width !== 0)
            writer.tag(3, WireType.Varint).int32(message.width);
        /* int32 height = 4; */
        if (message.height !== 0)
            writer.tag(4, WireType.Varint).int32(message.height);
        /* int64 time = 5; */
        if (message.time !== 0n)
            writer.tag(5, WireType.Varint).int64(message.time);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message input.WindowSizeInfo
 */
export const WindowSizeInfo = new WindowSizeInfo$Type();
// @generated message type with reflection information, may provide speed optimized methods
class EventReply$Type extends MessageType<EventReply> {
    constructor() {
        super("input.EventReply", [
            { no: 1, name: "message", kind: "scalar", T: 9 /*ScalarType.STRING*/ }
        ]);
    }
    create(value?: PartialMessage<EventReply>): EventReply {
        const message = { message: "" };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<EventReply>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: EventReply): EventReply {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* string message */ 1:
                    message.message = reader.string();
                    break;
                default:
                    let u = options.readUnknownField;
                    if (u === "throw")
                        throw new globalThis.Error(`Unknown field ${fieldNo} (wire type ${wireType}) for ${this.typeName}`);
                    let d = reader.skip(wireType);
                    if (u !== false)
                        (u === true ? UnknownFieldHandler.onRead : u)(this.typeName, message, fieldNo, wireType, d);
            }
        }
        return message;
    }
    internalBinaryWrite(message: EventReply, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* string message = 1; */
        if (message.message !== "")
            writer.tag(1, WireType.LengthDelimited).string(message.message);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message input.EventReply
 */
export const EventReply = new EventReply$Type();
/**
 * @generated ServiceType for protobuf service input.InputService
 */
export const InputService = new ServiceType("input.InputService", [
    { name: "LogMouseUp", options: {}, I: MouseClickInfo, O: EventReply },
    { name: "LogMouseDown", options: {}, I: MouseClickInfo, O: EventReply },
    { name: "LogMouseMoved", options: {}, I: MouseMoveInfo, O: EventReply },
    { name: "LogMouseScrolled", options: {}, I: MouseScrollInfo, O: EventReply },
    { name: "LogKeystroke", options: {}, I: KeystrokeInfo, O: EventReply },
    { name: "LogWindowSized", options: {}, I: WindowSizeInfo, O: EventReply }
]);
