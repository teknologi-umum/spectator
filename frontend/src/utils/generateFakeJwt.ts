// TODO(elianiva): replace this with proper JWT stuff once the server is up and
//                 running. This is only used for testing purpose. PLZ DELET DIS

const FAKE_HEADER = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9";
const FAKE_SIGNATURE = "ySxmfXC7SSP4NR7Go2qitririWvL-vWMLZUDjY0w6U8";

export function getJwt(): string {
  const fakePayload = window
    .btoa(
      JSON.stringify({
        studentNumber: "1234567890",
        iat: Date.now(),
        // ubah ini kalo mau nge-tes
        exp: Date.now() + 5 * 1000 // 90 minutes from iat
      })
    )
    .slice(0, -2); // remove trailing `=` from btoa

  return [FAKE_HEADER, fakePayload, FAKE_SIGNATURE].join(".");
}
