import { eventSpoke } from "@/spoke";
import { loggerInstance } from "@/spoke/logger";
import { KeystrokeInfo } from "@/stub/input";
import { LogLevel } from "@microsoft/signalr";

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

export function keystrokeHandler(
  questionNumber: number | null,
  accessToken: string | null
) {
  return async (e: KeyboardEvent) => {
    if (questionNumber === null || accessToken === null) return;

    const data: KeystrokeInfo = {
      accessToken: accessToken,
      questionNumber: questionNumber,
      keyChar: e.code,
      shift: e.shiftKey,
      alt: e.altKey,
      control: e.ctrlKey,
      meta: e.metaKey,
      unrelatedKey: false,
      time: Date.now() as unknown as bigint
    };

    // ignore if it's triggered from codemirror because we it has separate
    // listener
    if ((e.target as HTMLDivElement).classList[0] === "cm-content") {
      // everything INSIDE the editor is always related except F-keys
      data.unrelatedKey = F_KEYS[e.key] !== undefined;

      // don't allow pressing F-keys inside the editor
      if (F_KEYS[e.key]) e.preventDefault();

      try {
        await eventSpoke.keystroke(data);
      } catch (err) {
        if (err instanceof Error) {
          loggerInstance.log(LogLevel.Error, err.message);
        }
      }

      return;
    }

    // everything OUTSIDE the editor is always unrelated
    data.unrelatedKey = true;

    try {
      await eventSpoke.keystroke(data);
    } catch (err) {
      if (err instanceof Error) {
        loggerInstance.log(LogLevel.Error, err.message);
      }
    }

    // don't allow to do anything outside of the code editor
    e.preventDefault();
  };
}
