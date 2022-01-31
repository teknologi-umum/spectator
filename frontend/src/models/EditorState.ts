import type { EditorSnapshot } from "@/models/EditorSnapshot";
import type { Language } from "@/models/Language";
import type { Question } from "@/models/Question";

export interface EditorState {
  deadlineUtc: number | null;
  questions: Question[] | null;
  currentQuestionNumber: number | null;
  currentLanguage: Language;
  fontSize: number;
  snapshotByQuestionNumber: Record<number, EditorSnapshot>;
}
