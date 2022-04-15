import { eventSpoke } from "@/spoke";
import { MouseScrollRequest } from "@/stub/events";
import type { UIEvent } from "react";

// Only for testing purposes
// could be deleted later or being adjusted with different concepts
// after backend is ready
export function scrollHandler<T extends Element>(
  questionNumber: number | null,
  accessToken: string | null
) {
  return async (e: Event | UIEvent<T, globalThis.UIEvent>) => {
    if (questionNumber === null || accessToken === null) return;

    const data: MouseScrollRequest = {
      accessToken: accessToken,
      time: Date.now() as unknown as bigint
    };

    try {
      await eventSpoke.mouseScrolled(data);
    } catch (err) {
      console.error(err);
    }
  };
}
