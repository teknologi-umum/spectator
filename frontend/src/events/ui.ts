import { eventSpoke } from "@/spoke";
import { MouseScrollInfo } from "@/stub/input";
import type { UIEvent } from "react";

export function scrollHandler<T extends Element>(
  questionNumber: number | null,
  accessToken: string | null
) {
  return async (e: Event | UIEvent<T, globalThis.UIEvent>) => {
    if (questionNumber === null || accessToken === null) return;

    const data: MouseScrollInfo = {
      accessToken: accessToken,
      // TODO(elianiva): implement proper logic
      delta: (e.currentTarget as HTMLElement).scrollTop,
      x: 0,
      y: 0,
      questionNumber: questionNumber,
      time: Date.now() as unknown as bigint
    };

    try {
      await eventSpoke.mouseScrolled(data);
    } catch (err) {
      console.error(err);
    }
  };
}
