import {
  LocaleInfo,
  SessionReply,
  PersonalInfo,
  SAM,
  SubmissionRequest,
  Exam,
  SubmissionResult,
  ExamResult
} from "@/stub/session";
import SpokeBase from "@/spoke/spokeBase";

class SessionSpoke extends SpokeBase {
  public async startSession(localeInfo: LocaleInfo): Promise<SessionReply> {
    await this._startIfDisconnected();
    return this._hubConnection.invoke("StartSessionAsync", localeInfo);
  }

  public async setLocale(localeInfo: LocaleInfo): Promise<void> {
    await this._startIfDisconnected();
    return this._hubConnection.invoke("SetLocaleAsync", localeInfo);
  }

  public async submitPersonalInfo(personalInfo: PersonalInfo): Promise<void> {
    await this._startIfDisconnected();
    return this._hubConnection.invoke("SubmitPersonalInfoAsync", personalInfo);
  }

  public async submitBeforeExamSAM(sam: SAM): Promise<void> {
    await this._startIfDisconnected();
    return this._hubConnection.invoke("SubmitBeforeExamSAMAsync", sam);
  }

  public async startExam(): Promise<Exam> {
    await this._startIfDisconnected();
    return this._hubConnection.invoke("StartExamAsync");
  }

  public async resumeExam(): Promise<Exam> {
    await this._startIfDisconnected();
    return this._hubConnection.invoke("ResumeExamAsync");
  }

  public async submitSolution(
    submissionRequest: SubmissionRequest
  ): Promise<SubmissionResult> {
    await this._startIfDisconnected();
    return this._hubConnection.invoke("SubmitSolutionAsync", submissionRequest);
  }

  public async endExam(): Promise<ExamResult> {
    await this._startIfDisconnected();
    return this._hubConnection.invoke("EndExamAsync");
  }

  public async passDeadline(): Promise<ExamResult> {
    await this._startIfDisconnected();
    return this._hubConnection.invoke("PassDeadlineAsync");
  }

  public async forfeitExam(): Promise<ExamResult> {
    await this._startIfDisconnected();
    return this._hubConnection.invoke("ForfeitExamAsync");
  }

  public async submitAfterExamSAM(sam: SAM): Promise<void> {
    await this._startIfDisconnected();
    return this._hubConnection.invoke("SubmitAfterExamSAMAsync", sam);
  }
}

export default new SessionSpoke(import.meta.env.VITE_SESSION_HUB_URL);
