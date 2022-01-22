import type { Language } from "./Language";

export interface Question {
  questionNumber: number;
  title: string;
  instruction: string;
  templateByLanguage: Record<Language, string>;
}