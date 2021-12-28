// TODO(elianiva): replace this with proper JWT stuff once the server is up and
//                 running. This is only used for testing purpose. PLZ DELET DIS

const FAKE_HEADER = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9";
const FAKE_SIGNATURE = "ySxmfXC7SSP4NR7Go2qitririWvL-vWMLZUDjY0w6U8";

export function getJwt(): string {
  const fakePayload = window.btoa(
    JSON.stringify({
      studentNumber: "1234567890",
      iat: Date.now(),
      // ideally this should be 90 minutes, but for testing purpose a minute
      // is enough
      exp: 60 * 1000
    })
  );

  return [FAKE_HEADER, fakePayload, FAKE_SIGNATURE].join(".");
}
