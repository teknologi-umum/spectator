import { emit } from "@/events/emitter";
import type { CodingEventKeystroke } from "./types";

const F_KEYS: Record<string, boolean> = {
  F1: true,
  F2: true,
  F3: true,
  F4: true,
  F5: true,
  F6: true,
  F7: true,
  F8: true,
  F9: true,
  F10: true,
  F11: true,
  F12: true
};

export function keystrokeHandler(connection: unknown, questionNumber: number) {
  return async (e: KeyboardEvent) => {
    const data: CodingEventKeystroke = {
      // TODO(elianiva): revisit session_id
      session_id: "TBD",
      type: "coding_event_keystroke",
      question_number: questionNumber,
      key_char: e.key,
      key_code: e.keyCode,
      shift: e.shiftKey,
      alt: e.altKey,
      control: e.ctrlKey,
      meta: e.metaKey,
      unrelated_key: false,
      time: new Date(Date.now())
    };

    // ignore if it's triggered from codemirror because we it has separate
    // listener
    if ((e.target as HTMLDivElement).classList[0] === "cm-content") {
      // everything INSIDE the editor is always related except F-keys
      data.unrelated_key = F_KEYS[e.key] !== undefined;

      // don't allow pressing F-keys inside the editor
      if (F_KEYS[e.key]) e.preventDefault();

      try {
        await emit(connection, data);
      } catch (err) {
        // TODO(elianiva): replace with proper logging
        console.error(err);
      }
      return;
    }

    // everything OUTSIDE the editor is always unrelated
    data.unrelated_key = true;
    try {
      await emit(connection, data);
    } catch (err) {
      // TODO(elianiva): replace with proper logging
      console.error(err);
    }

    // don't allow to do anything outside of the code editor
    e.preventDefault();
  };
}
