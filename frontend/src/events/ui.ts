import { eventSpoke } from "@/spoke";
import { loggerInstance } from "@/spoke/logger";
import { WindowSizeInfo } from "@/stub/input";
import { LogLevel } from "@microsoft/signalr";

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
      if (err instanceof Error) {
        await loggerInstance.log(LogLevel.Error, err.message);
      }
    }
  };
}
