export interface SessionState {
  sessionId: string | null;
  accessToken: string | null;
  firstSAMSubmitted: boolean;
  secondSAMSubmitted: boolean;
  hasPermission: boolean;
  deviceId: string | null;
  tourCompleted: {
    personalInfo: boolean;
    samTest: boolean;
    codingTest: boolean;
  };
}
