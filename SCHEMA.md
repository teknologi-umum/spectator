# Schema for Spectator

Based on [backend/Spectator.DomainEvents](https://github.com/teknologi-umum/spectator/tree/master/backend/Spectator.DomainEvents/SessionDomain)

<details>
<summary>Enums</summary>

```typescript
enum Locale {
  EN = 0;
  ID = 1;
}

enum Language {
  Undefined = 0,
  C = 1,
  CPP = 2,
  PHP = 3,
  Javascript = 4,
  Java = 5,
  Python = 6
}

enum MouseButton {
  Left,
  Right,
  Middle
}
```

</details>

<details>
<summary>Session</summary>

```typescript
// measurement: after_exam_sam_submitted
interface AfterExamSAMSubmitted {
  session_id: string; // tag
  // fields
  aroused_level: number;
  pleased_level: number;
}

// measurement: before_exam_sam_submitted
interface BeforeExamSAMSubmitted {
  session_id: string; // tag
  // fields
  aroused_level: number;
  pleased_level: number;
}

// measurement: deadline_passed
interface DeadlinePassed {
  session_id: string; // tag
}

// measurement: exam_ended
interface ExamEnded {
  session_id: string; // tag
}

// measurement: exam_forfeited
interface ExamForfeited {
  session_id: string; // tag
}

// measurement: exam_ide_reloaded
interface ExamIDEReloaded {
  session_id: string; // tag
}

// measurement: exam_started
interface ExamStarted {
  session_id: string; // tag
  question_numbers: number[]; // field
  deadline: Date; // field
}

// measurement: locale_set
interface LocaleSet {
  session_id: string; // tag
  locale: Locale; // field
}

// measurement: personal_info_submitted
interface PersonalInfoSubmitted {
  session_id: string; // tag

  // fields
  student_number: string;
  years_of_experience: number;
  hours_of_practice: number;
  familiar_languages: string;
}

// measurement: session_started
interface SessionStarted {
  session_id: string; // tag
  locale: Locale; // field
}

// measurement: solution_accepted
interface SolutionAccepted {
  session_id: string; // tag

  // fields
  question_number: number;
  language: Language;
  solution: string;
  scratchpad: string;
  serialized_test_result: string;
}

// measurement: solution_rejected
interface SolutionRejected {
  session_id: string; // tag

  // fields
  question_number: number;
  language: Language;
  solution: string;
  scratchpad: string;
  serialized_test_result: string;
}
```

</details>

<details>
<summary>Input</summary>

```typescript
// measurement: keystroke
interface Keystroke {
  session_id: string; // tag
  
  // fields
  key_code: string;
  key_char: string;
  shift: boolean;
  alt: boolean;
  ctrl: boolean;
}

// measurement: mouse_down
interface MouseDown {
  session_id: string; // tag

  // fields
  x: number;
  y: number;
  button: MouseButton;
}

// measurement: mouse_up
interface MouseScrolled {
  session_id: string; // tag

  // fields
  x: number;
  y: number;
}

// measurement: mouse_moved
interface MouseMoved {
  session_id: string; // tag

  // fields
  x: number;
  y: number;
}

// measurement: mouse_scrolled
interface MouseScrolled {
  session_id: string; // tag

  // fields
  x: number;
  y: number;
  delta: number;
}

// measurement: window_sized
interface WindowSized {
  session_id: string; // tag

  // fields
  width: number;
  height: number;
}
```

</details>
