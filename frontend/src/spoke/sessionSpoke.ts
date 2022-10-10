import {
  StartSessionRequest,
  SessionReply,
  SetLocaleRequest,
  SubmitPersonalInfoRequest,
  SubmitSAMRequest,
  EmptyRequest,
  SubmissionRequest,
  Exam,
  ExamResult,
  SubmitSolutionSAMRequest
} from "@/stub/session";
import SpokeBase from "@/spoke/spokeBase";
import { SESSION_HUB_URL } from "@/constants";
import { SubmissionResult } from "@/models/SubmissionResult";

class SessionSpoke extends SpokeBase {
  public async startSession(
    request: StartSessionRequest
  ): Promise<SessionReply> {
    await super._startIfDisconnected();
    return this._hubConnection.invoke("StartSessionAsync", request);
  }

  public resumeSession(): Promise<void> {
    return super._startIfDisconnected();
  }

  public async setLocale(request: SetLocaleRequest): Promise<void> {
    await super._startIfDisconnected();
    return this._hubConnection.invoke("SetLocaleAsync", request);
  }

  public async submitPersonalInfo(
    request: SubmitPersonalInfoRequest
  ): Promise<void> {
    await super._startIfDisconnected();
    return this._hubConnection.invoke("SubmitPersonalInfoAsync", request);
  }

  public async submitBeforeExamSAM(request: SubmitSAMRequest): Promise<void> {
    await super._startIfDisconnected();
    return this._hubConnection.invoke("SubmitBeforeExamSAMAsync", request);
  }

  public async startExam(request: EmptyRequest): Promise<Exam> {
    await super._startIfDisconnected();
    return this._hubConnection.invoke("StartExamAsync", request);
  }

  public async resumeExam(request: EmptyRequest): Promise<Exam> {
    await super._startIfDisconnected();
    return this._hubConnection.invoke("ResumeExamAsync", request);
  }

  public async submitSolution(
    request: SubmissionRequest
  ): Promise<SubmissionResult> {
    await super._startIfDisconnected();
    return this._hubConnection.invoke("SubmitSolutionAsync", request);
  }

  public async testSolution(
    request: SubmissionRequest
  ): Promise<SubmissionResult> {
    await super._startIfDisconnected();
    return this._hubConnection.invoke("TestSolutionAsync", request);
  }

  public async endExam(request: EmptyRequest): Promise<ExamResult> {
    await super._startIfDisconnected();
    return this._hubConnection.invoke("EndExamAsync", request);
  }

  public async passDeadline(request: EmptyRequest): Promise<ExamResult> {
    await super._startIfDisconnected();
    return this._hubConnection.invoke("PassDeadlineAsync", request);
  }

  public async forfeitExam(request: EmptyRequest): Promise<ExamResult> {
    await super._startIfDisconnected();
    return this._hubConnection.invoke("ForfeitExamAsync", request);
  }

  public async submitAfterExamSAM(request: SubmitSAMRequest): Promise<void> {
    await super._startIfDisconnected();
    return this._hubConnection.invoke("SubmitAfterExamSAMAsync", request);
  }

  public async submitSolutionSAM(request: SubmitSolutionSAMRequest): Promise<void> {
    await super._startIfDisconnected();
    return this._hubConnection.invoke("SubmitSolutionSAMAsync", request);
  }
}

export default new SessionSpoke(SESSION_HUB_URL);
