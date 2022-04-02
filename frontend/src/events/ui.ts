import { eventSpoke } from "@/spoke";
import type { UIEvent } from "react";

// Only for testing purposes
// could be deleted later or being adjusted with different concepts
// after backend is ready
export function scrollHandler<T extends Element>(
  questionNumber: number | null
) {
  return async (e: Event | UIEvent<T, globalThis.UIEvent>) => {
    if (questionNumber === null) return;

    try {
      await eventSpoke.mouseScrolled(e);
    } catch (err) {
      console.error(err);
    }
  };
}
