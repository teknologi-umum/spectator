// TODO(elianiva): use typings from proto

import type { Direction } from "@/utils/getMouseDirection";

// Bucket name: public
export interface Session {
  // Tags
  type: "session_started";
  session_id: string;
  // Fields
  token: string;
  time: Date;
}

export interface PersonalInfo {
  // Tags
  type: "personal_info";
  session_id: string;
  // Fields
  student_number: string;
  hours_of_practice: number;
  years_of_experience: number;
  familiar_languages: string;
  time: Date;
}

export interface CodingSubmission {
  // Tags
  session_id: string;
  type: "code_submission" | "code_test_attempt";
  question_number: number;
  // Fields
  code: string;
  language: string;
  _time: Date;
}

export interface CodingStatus {
  // Tags
  session_id: string;
  type:
  | "coding_started"
  | "coding_completed"
  | "coding_paused"
  | "coding_resumed";
  question_number: number;
  // Fields
  time: Date;
}

export interface CodingEventKeystroke {
  // Tags
  session_id: string;
  type: "coding_event_keystroke";
  question_number: number;
  // Fields
  key_char: string;
  key_code: number;
  shift: boolean;
  alt: boolean;
  control: boolean;
  meta: boolean;
  unrelated_key: boolean;
  time: Date;
}

export interface CodingEventMouseMove {
  // Tags
  session_id: string;
  type: "coding_event_mousemove";
  question_number: number;
  // Fields
  direction: Direction;
  x_position: number;
  y_position: number;
  window_width: number;
  window_height: number;
  time: Date;
}

export interface CodingEventMouseClick {
  // Tags
  session_id: string;
  type: "coding_event_mouseclick";
  question_number: number;
  // Fields
  right_click: boolean;
  left_click: boolean;
  middle_click: boolean;
  time: Date;
}

export interface SelfAssessmentManikinSubmitted {
  // Tags
  session_id: string;
  type: "sam_test_before";
  // Fields
  aroused_level: number;
  pleased_level: number;
  time: Date;
}

// Bucket name: results
export interface TestResult {
  // Tags
  session_id: string;
  student_number: string;
  // Fields
  file_url_json: string;
  file_url_csv: string;
  time: Date;
}
