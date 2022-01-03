// @generated by protobuf-ts 2.1.0
// @generated from protobuf file "session.proto" (package "session", syntax proto3)
// tslint:disable
import { ServiceType } from "@protobuf-ts/runtime-rpc";
import { WireType } from "@protobuf-ts/runtime";
import type { BinaryWriteOptions } from "@protobuf-ts/runtime";
import type { IBinaryWriter } from "@protobuf-ts/runtime";
import { UnknownFieldHandler } from "@protobuf-ts/runtime";
import type { BinaryReadOptions } from "@protobuf-ts/runtime";
import type { IBinaryReader } from "@protobuf-ts/runtime";
import type { PartialMessage } from "@protobuf-ts/runtime";
import { reflectionMergePartial } from "@protobuf-ts/runtime";
import { MESSAGE_TYPE } from "@protobuf-ts/runtime";
import { MessageType } from "@protobuf-ts/runtime";
import { Language } from "./enums";
/**
 * @generated from protobuf message session.StartSessionRequest
 */
export interface StartSessionRequest {
}
/**
 * @generated from protobuf message session.SessionReply
 */
export interface SessionReply {
    /**
     * @generated from protobuf field: string access_token = 1;
     */
    accessToken: string;
}
/**
 * @generated from protobuf message session.PersonalInfo
 */
export interface PersonalInfo {
    /**
     * @generated from protobuf field: string student_number = 1;
     */
    studentNumber: string;
    /**
     * @generated from protobuf field: int32 years_of_experience = 2;
     */
    yearsOfExperience: number;
    /**
     * @generated from protobuf field: int32 hours_of_practice = 3;
     */
    hoursOfPractice: number;
    /**
     * @generated from protobuf field: string familiar_languages = 4;
     */
    familiarLanguages: string;
}
/**
 * @generated from protobuf message session.SAM
 */
export interface SAM {
    /**
     * @generated from protobuf field: int32 aroused_level = 1;
     */
    arousedLevel: number;
    /**
     * @generated from protobuf field: int32 pleased_level = 2;
     */
    pleasedLevel: number;
}
/**
 * @generated from protobuf message session.Question
 */
export interface Question {
    /**
     * @generated from protobuf field: int32 question_number = 1;
     */
    questionNumber: number;
    /**
     * @generated from protobuf field: string title = 2;
     */
    title: string;
    /**
     * @generated from protobuf field: string instruction = 3;
     */
    instruction: string;
    /**
     * @generated from protobuf field: repeated enums.Language allowed_languages = 4;
     */
    allowedLanguages: Language[];
    /**
     * @generated from protobuf field: string boilerplate = 5;
     */
    boilerplate: string;
}
/**
 * @generated from protobuf message session.Exam
 */
export interface Exam {
    /**
     * @generated from protobuf field: int64 deadline = 1;
     */
    deadline: bigint;
    /**
     * @generated from protobuf field: repeated session.Question questions = 2;
     */
    questions: Question[];
    /**
     * @generated from protobuf field: repeated int32 answered_question_numbers = 3;
     */
    answeredQuestionNumbers: number[];
}
/**
 * @generated from protobuf message session.ExamResult
 */
export interface ExamResult {
    /**
     * @generated from protobuf field: int64 duration = 1;
     */
    duration: bigint;
    /**
     * @generated from protobuf field: repeated int32 answered_question_numbers = 2;
     */
    answeredQuestionNumbers: number[];
}
/**
 * @generated from protobuf message session.SubmissionRequest
 */
export interface SubmissionRequest {
    /**
     * @generated from protobuf field: int32 question_number = 1;
     */
    questionNumber: number;
    /**
     * @generated from protobuf field: enums.Language language = 2;
     */
    language: Language;
    /**
     * @generated from protobuf field: string solution = 3;
     */
    solution: string;
    /**
     * @generated from protobuf field: string scratch_pad = 4;
     */
    scratchPad: string;
}
/**
 * @generated from protobuf message session.SubmissionResult
 */
export interface SubmissionResult {
    /**
     * @generated from protobuf field: bool accepted = 1;
     */
    accepted: boolean;
    /**
     * @generated from protobuf field: repeated session.SubmissionResult.TestResult test_results = 2;
     */
    testResults: SubmissionResult_TestResult[];
}
/**
 * @generated from protobuf message session.SubmissionResult.TestResult
 */
export interface SubmissionResult_TestResult {
    /**
     * @generated from protobuf field: bool success = 1;
     */
    success: boolean;
    /**
     * @generated from protobuf field: string message = 2;
     */
    message: string;
}
/**
 * @generated from protobuf message session.EmptyRequest
 */
export interface EmptyRequest {
}
/**
 * @generated from protobuf message session.EmptyReply
 */
export interface EmptyReply {
}
// @generated message type with reflection information, may provide speed optimized methods
class StartSessionRequest$Type extends MessageType<StartSessionRequest> {
    constructor() {
        super("session.StartSessionRequest", []);
    }
    create(value?: PartialMessage<StartSessionRequest>): StartSessionRequest {
        const message = {};
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<StartSessionRequest>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: StartSessionRequest): StartSessionRequest {
        return target ?? this.create();
    }
    internalBinaryWrite(message: StartSessionRequest, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message session.StartSessionRequest
 */
export const StartSessionRequest = new StartSessionRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class SessionReply$Type extends MessageType<SessionReply> {
    constructor() {
        super("session.SessionReply", [
            { no: 1, name: "access_token", kind: "scalar", T: 9 /*ScalarType.STRING*/ }
        ]);
    }
    create(value?: PartialMessage<SessionReply>): SessionReply {
        const message = { accessToken: "" };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<SessionReply>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: SessionReply): SessionReply {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* string access_token */ 1:
                    message.accessToken = reader.string();
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
    internalBinaryWrite(message: SessionReply, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* string access_token = 1; */
        if (message.accessToken !== "")
            writer.tag(1, WireType.LengthDelimited).string(message.accessToken);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message session.SessionReply
 */
export const SessionReply = new SessionReply$Type();
// @generated message type with reflection information, may provide speed optimized methods
class PersonalInfo$Type extends MessageType<PersonalInfo> {
    constructor() {
        super("session.PersonalInfo", [
            { no: 1, name: "student_number", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 2, name: "years_of_experience", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 3, name: "hours_of_practice", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 4, name: "familiar_languages", kind: "scalar", T: 9 /*ScalarType.STRING*/ }
        ]);
    }
    create(value?: PartialMessage<PersonalInfo>): PersonalInfo {
        const message = { studentNumber: "", yearsOfExperience: 0, hoursOfPractice: 0, familiarLanguages: "" };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<PersonalInfo>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: PersonalInfo): PersonalInfo {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* string student_number */ 1:
                    message.studentNumber = reader.string();
                    break;
                case /* int32 years_of_experience */ 2:
                    message.yearsOfExperience = reader.int32();
                    break;
                case /* int32 hours_of_practice */ 3:
                    message.hoursOfPractice = reader.int32();
                    break;
                case /* string familiar_languages */ 4:
                    message.familiarLanguages = reader.string();
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
    internalBinaryWrite(message: PersonalInfo, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* string student_number = 1; */
        if (message.studentNumber !== "")
            writer.tag(1, WireType.LengthDelimited).string(message.studentNumber);
        /* int32 years_of_experience = 2; */
        if (message.yearsOfExperience !== 0)
            writer.tag(2, WireType.Varint).int32(message.yearsOfExperience);
        /* int32 hours_of_practice = 3; */
        if (message.hoursOfPractice !== 0)
            writer.tag(3, WireType.Varint).int32(message.hoursOfPractice);
        /* string familiar_languages = 4; */
        if (message.familiarLanguages !== "")
            writer.tag(4, WireType.LengthDelimited).string(message.familiarLanguages);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message session.PersonalInfo
 */
export const PersonalInfo = new PersonalInfo$Type();
// @generated message type with reflection information, may provide speed optimized methods
class SAM$Type extends MessageType<SAM> {
    constructor() {
        super("session.SAM", [
            { no: 1, name: "aroused_level", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 2, name: "pleased_level", kind: "scalar", T: 5 /*ScalarType.INT32*/ }
        ]);
    }
    create(value?: PartialMessage<SAM>): SAM {
        const message = { arousedLevel: 0, pleasedLevel: 0 };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<SAM>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: SAM): SAM {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* int32 aroused_level */ 1:
                    message.arousedLevel = reader.int32();
                    break;
                case /* int32 pleased_level */ 2:
                    message.pleasedLevel = reader.int32();
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
    internalBinaryWrite(message: SAM, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* int32 aroused_level = 1; */
        if (message.arousedLevel !== 0)
            writer.tag(1, WireType.Varint).int32(message.arousedLevel);
        /* int32 pleased_level = 2; */
        if (message.pleasedLevel !== 0)
            writer.tag(2, WireType.Varint).int32(message.pleasedLevel);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message session.SAM
 */
export const SAM = new SAM$Type();
// @generated message type with reflection information, may provide speed optimized methods
class Question$Type extends MessageType<Question> {
    constructor() {
        super("session.Question", [
            { no: 1, name: "question_number", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 2, name: "title", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 3, name: "instruction", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 4, name: "allowed_languages", kind: "enum", repeat: 1 /*RepeatType.PACKED*/, T: () => ["enums.Language", Language] },
            { no: 5, name: "boilerplate", kind: "scalar", T: 9 /*ScalarType.STRING*/ }
        ]);
    }
    create(value?: PartialMessage<Question>): Question {
        const message = { questionNumber: 0, title: "", instruction: "", allowedLanguages: [], boilerplate: "" };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<Question>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: Question): Question {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* int32 question_number */ 1:
                    message.questionNumber = reader.int32();
                    break;
                case /* string title */ 2:
                    message.title = reader.string();
                    break;
                case /* string instruction */ 3:
                    message.instruction = reader.string();
                    break;
                case /* repeated enums.Language allowed_languages */ 4:
                    if (wireType === WireType.LengthDelimited)
                        for (let e = reader.int32() + reader.pos; reader.pos < e;)
                            message.allowedLanguages.push(reader.int32());
                    else
                        message.allowedLanguages.push(reader.int32());
                    break;
                case /* string boilerplate */ 5:
                    message.boilerplate = reader.string();
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
    internalBinaryWrite(message: Question, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* int32 question_number = 1; */
        if (message.questionNumber !== 0)
            writer.tag(1, WireType.Varint).int32(message.questionNumber);
        /* string title = 2; */
        if (message.title !== "")
            writer.tag(2, WireType.LengthDelimited).string(message.title);
        /* string instruction = 3; */
        if (message.instruction !== "")
            writer.tag(3, WireType.LengthDelimited).string(message.instruction);
        /* repeated enums.Language allowed_languages = 4; */
        if (message.allowedLanguages.length) {
            writer.tag(4, WireType.LengthDelimited).fork();
            for (let i = 0; i < message.allowedLanguages.length; i++)
                writer.int32(message.allowedLanguages[i]);
            writer.join();
        }
        /* string boilerplate = 5; */
        if (message.boilerplate !== "")
            writer.tag(5, WireType.LengthDelimited).string(message.boilerplate);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message session.Question
 */
export const Question = new Question$Type();
// @generated message type with reflection information, may provide speed optimized methods
class Exam$Type extends MessageType<Exam> {
    constructor() {
        super("session.Exam", [
            { no: 1, name: "deadline", kind: "scalar", T: 3 /*ScalarType.INT64*/, L: 0 /*LongType.BIGINT*/ },
            { no: 2, name: "questions", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => Question },
            { no: 3, name: "answered_question_numbers", kind: "scalar", repeat: 1 /*RepeatType.PACKED*/, T: 5 /*ScalarType.INT32*/ }
        ]);
    }
    create(value?: PartialMessage<Exam>): Exam {
        const message = { deadline: 0n, questions: [], answeredQuestionNumbers: [] };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<Exam>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: Exam): Exam {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* int64 deadline */ 1:
                    message.deadline = reader.int64().toBigInt();
                    break;
                case /* repeated session.Question questions */ 2:
                    message.questions.push(Question.internalBinaryRead(reader, reader.uint32(), options));
                    break;
                case /* repeated int32 answered_question_numbers */ 3:
                    if (wireType === WireType.LengthDelimited)
                        for (let e = reader.int32() + reader.pos; reader.pos < e;)
                            message.answeredQuestionNumbers.push(reader.int32());
                    else
                        message.answeredQuestionNumbers.push(reader.int32());
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
    internalBinaryWrite(message: Exam, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* int64 deadline = 1; */
        if (message.deadline !== 0n)
            writer.tag(1, WireType.Varint).int64(message.deadline);
        /* repeated session.Question questions = 2; */
        for (let i = 0; i < message.questions.length; i++)
            Question.internalBinaryWrite(message.questions[i], writer.tag(2, WireType.LengthDelimited).fork(), options).join();
        /* repeated int32 answered_question_numbers = 3; */
        if (message.answeredQuestionNumbers.length) {
            writer.tag(3, WireType.LengthDelimited).fork();
            for (let i = 0; i < message.answeredQuestionNumbers.length; i++)
                writer.int32(message.answeredQuestionNumbers[i]);
            writer.join();
        }
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message session.Exam
 */
export const Exam = new Exam$Type();
// @generated message type with reflection information, may provide speed optimized methods
class ExamResult$Type extends MessageType<ExamResult> {
    constructor() {
        super("session.ExamResult", [
            { no: 1, name: "duration", kind: "scalar", T: 3 /*ScalarType.INT64*/, L: 0 /*LongType.BIGINT*/ },
            { no: 2, name: "answered_question_numbers", kind: "scalar", repeat: 1 /*RepeatType.PACKED*/, T: 5 /*ScalarType.INT32*/ }
        ]);
    }
    create(value?: PartialMessage<ExamResult>): ExamResult {
        const message = { duration: 0n, answeredQuestionNumbers: [] };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<ExamResult>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: ExamResult): ExamResult {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* int64 duration */ 1:
                    message.duration = reader.int64().toBigInt();
                    break;
                case /* repeated int32 answered_question_numbers */ 2:
                    if (wireType === WireType.LengthDelimited)
                        for (let e = reader.int32() + reader.pos; reader.pos < e;)
                            message.answeredQuestionNumbers.push(reader.int32());
                    else
                        message.answeredQuestionNumbers.push(reader.int32());
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
    internalBinaryWrite(message: ExamResult, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* int64 duration = 1; */
        if (message.duration !== 0n)
            writer.tag(1, WireType.Varint).int64(message.duration);
        /* repeated int32 answered_question_numbers = 2; */
        if (message.answeredQuestionNumbers.length) {
            writer.tag(2, WireType.LengthDelimited).fork();
            for (let i = 0; i < message.answeredQuestionNumbers.length; i++)
                writer.int32(message.answeredQuestionNumbers[i]);
            writer.join();
        }
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message session.ExamResult
 */
export const ExamResult = new ExamResult$Type();
// @generated message type with reflection information, may provide speed optimized methods
class SubmissionRequest$Type extends MessageType<SubmissionRequest> {
    constructor() {
        super("session.SubmissionRequest", [
            { no: 1, name: "question_number", kind: "scalar", T: 5 /*ScalarType.INT32*/ },
            { no: 2, name: "language", kind: "enum", T: () => ["enums.Language", Language] },
            { no: 3, name: "solution", kind: "scalar", T: 9 /*ScalarType.STRING*/ },
            { no: 4, name: "scratch_pad", kind: "scalar", T: 9 /*ScalarType.STRING*/ }
        ]);
    }
    create(value?: PartialMessage<SubmissionRequest>): SubmissionRequest {
        const message = { questionNumber: 0, language: 0, solution: "", scratchPad: "" };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<SubmissionRequest>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: SubmissionRequest): SubmissionRequest {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* int32 question_number */ 1:
                    message.questionNumber = reader.int32();
                    break;
                case /* enums.Language language */ 2:
                    message.language = reader.int32();
                    break;
                case /* string solution */ 3:
                    message.solution = reader.string();
                    break;
                case /* string scratch_pad */ 4:
                    message.scratchPad = reader.string();
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
    internalBinaryWrite(message: SubmissionRequest, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* int32 question_number = 1; */
        if (message.questionNumber !== 0)
            writer.tag(1, WireType.Varint).int32(message.questionNumber);
        /* enums.Language language = 2; */
        if (message.language !== 0)
            writer.tag(2, WireType.Varint).int32(message.language);
        /* string solution = 3; */
        if (message.solution !== "")
            writer.tag(3, WireType.LengthDelimited).string(message.solution);
        /* string scratch_pad = 4; */
        if (message.scratchPad !== "")
            writer.tag(4, WireType.LengthDelimited).string(message.scratchPad);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message session.SubmissionRequest
 */
export const SubmissionRequest = new SubmissionRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class SubmissionResult$Type extends MessageType<SubmissionResult> {
    constructor() {
        super("session.SubmissionResult", [
            { no: 1, name: "accepted", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 2, name: "test_results", kind: "message", repeat: 1 /*RepeatType.PACKED*/, T: () => SubmissionResult_TestResult }
        ]);
    }
    create(value?: PartialMessage<SubmissionResult>): SubmissionResult {
        const message = { accepted: false, testResults: [] };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<SubmissionResult>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: SubmissionResult): SubmissionResult {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* bool accepted */ 1:
                    message.accepted = reader.bool();
                    break;
                case /* repeated session.SubmissionResult.TestResult test_results */ 2:
                    message.testResults.push(SubmissionResult_TestResult.internalBinaryRead(reader, reader.uint32(), options));
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
    internalBinaryWrite(message: SubmissionResult, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* bool accepted = 1; */
        if (message.accepted !== false)
            writer.tag(1, WireType.Varint).bool(message.accepted);
        /* repeated session.SubmissionResult.TestResult test_results = 2; */
        for (let i = 0; i < message.testResults.length; i++)
            SubmissionResult_TestResult.internalBinaryWrite(message.testResults[i], writer.tag(2, WireType.LengthDelimited).fork(), options).join();
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message session.SubmissionResult
 */
export const SubmissionResult = new SubmissionResult$Type();
// @generated message type with reflection information, may provide speed optimized methods
class SubmissionResult_TestResult$Type extends MessageType<SubmissionResult_TestResult> {
    constructor() {
        super("session.SubmissionResult.TestResult", [
            { no: 1, name: "success", kind: "scalar", T: 8 /*ScalarType.BOOL*/ },
            { no: 2, name: "message", kind: "scalar", T: 9 /*ScalarType.STRING*/ }
        ]);
    }
    create(value?: PartialMessage<SubmissionResult_TestResult>): SubmissionResult_TestResult {
        const message = { success: false, message: "" };
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<SubmissionResult_TestResult>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: SubmissionResult_TestResult): SubmissionResult_TestResult {
        let message = target ?? this.create(), end = reader.pos + length;
        while (reader.pos < end) {
            let [fieldNo, wireType] = reader.tag();
            switch (fieldNo) {
                case /* bool success */ 1:
                    message.success = reader.bool();
                    break;
                case /* string message */ 2:
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
    internalBinaryWrite(message: SubmissionResult_TestResult, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        /* bool success = 1; */
        if (message.success !== false)
            writer.tag(1, WireType.Varint).bool(message.success);
        /* string message = 2; */
        if (message.message !== "")
            writer.tag(2, WireType.LengthDelimited).string(message.message);
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message session.SubmissionResult.TestResult
 */
export const SubmissionResult_TestResult = new SubmissionResult_TestResult$Type();
// @generated message type with reflection information, may provide speed optimized methods
class EmptyRequest$Type extends MessageType<EmptyRequest> {
    constructor() {
        super("session.EmptyRequest", []);
    }
    create(value?: PartialMessage<EmptyRequest>): EmptyRequest {
        const message = {};
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<EmptyRequest>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: EmptyRequest): EmptyRequest {
        return target ?? this.create();
    }
    internalBinaryWrite(message: EmptyRequest, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message session.EmptyRequest
 */
export const EmptyRequest = new EmptyRequest$Type();
// @generated message type with reflection information, may provide speed optimized methods
class EmptyReply$Type extends MessageType<EmptyReply> {
    constructor() {
        super("session.EmptyReply", []);
    }
    create(value?: PartialMessage<EmptyReply>): EmptyReply {
        const message = {};
        globalThis.Object.defineProperty(message, MESSAGE_TYPE, { enumerable: false, value: this });
        if (value !== undefined)
            reflectionMergePartial<EmptyReply>(this, message, value);
        return message;
    }
    internalBinaryRead(reader: IBinaryReader, length: number, options: BinaryReadOptions, target?: EmptyReply): EmptyReply {
        return target ?? this.create();
    }
    internalBinaryWrite(message: EmptyReply, writer: IBinaryWriter, options: BinaryWriteOptions): IBinaryWriter {
        let u = options.writeUnknownFields;
        if (u !== false)
            (u == true ? UnknownFieldHandler.onWrite : u)(this.typeName, message, writer);
        return writer;
    }
}
/**
 * @generated MessageType for protobuf message session.EmptyReply
 */
export const EmptyReply = new EmptyReply$Type();
/**
 * @generated ServiceType for protobuf service session.SessionService
 */
export const SessionService = new ServiceType("session.SessionService", [
    { name: "StartSession", options: {}, I: EmptyRequest, O: SessionReply },
    { name: "SubmitPersonalInfo", options: {}, I: PersonalInfo, O: EmptyReply },
    { name: "SubmitBeforeCodeSAM", options: {}, I: SAM, O: EmptyReply },
    { name: "StartExam", options: {}, I: EmptyRequest, O: Exam },
    { name: "ResumeExam", options: {}, I: EmptyRequest, O: Exam },
    { name: "EndExam", options: {}, I: EmptyRequest, O: ExamResult },
    { name: "SubmitSolution", options: {}, I: SubmissionRequest, O: SubmissionResult },
    { name: "SubmitAfterCodeSAM", options: {}, I: SAM, O: EmptyReply }
]);
