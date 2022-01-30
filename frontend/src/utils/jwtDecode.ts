/**
 * Decode a JWT token
 * @param token - The JWT token to decode
 * @returns The decoded token
 */
export function jwtDecode(token: string): {
  iat: number;
  exp: number;
} {
  return JSON.parse(window.atob(token.split(".")[1]));
}
