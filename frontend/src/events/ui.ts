import { eventSpoke } from "@/spoke";
import { WindowSizeInfo } from "@/stub/input";

export function windowResizeHandler(
  questionNumber: number | null,
  accessToken: string | null
) {
  return async () => {
    if (questionNumber === null || accessToken === null) return;

    const data: WindowSizeInfo = {
      accessToken: accessToken,
      questionNumber: questionNumber,
      width: window.innerWidth,
      height: window.innerHeight,
      time: Date.now() as unknown as bigint
    };

    try {
      await eventSpoke.windowResized(data);
    } catch (err) {
      console.error(err);
    }
  };
}
