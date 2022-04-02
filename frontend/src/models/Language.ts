export const LANGUAGES = [
  "java",
  "javascript",
  "php",
  "python",
  "cpp",
  "c"
] as const;

export type Language = typeof LANGUAGES[number];
