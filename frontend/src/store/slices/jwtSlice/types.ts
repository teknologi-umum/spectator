export interface InitialState {
  jwt: string;
  jwtPayload: Record<string, unknown>;
  hasFinished: boolean;
}
