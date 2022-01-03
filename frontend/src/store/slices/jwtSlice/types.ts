export interface InitialState {
  jwt: string;
  // TODO(elianiva): revisit this later when we have a proper JWT support
  jwtPayload: {
    exp: number;
    iat: number;
    studentNumber: string;
  };
  hasFinished: boolean;
}
