export const THEMES = ["light", "dimmed", "dark"] as const;

export type Theme = typeof THEMES[number];
