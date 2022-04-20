import { loggerInstance } from "@/spoke/logger";
import { LogLevel } from "@microsoft/signalr";

export function uploadVideo(accessToken: string, startedAt: number) {
  return async (e: BlobEvent) => {
    const formData = new FormData();
    formData.append("startedAt", startedAt.toString());
    formData.append("stoppedAt", Date.now().toString());
    formData.append("file", e.data);

    try {
      await fetch(import.meta.env.VITE_VIDEO_STREAM_URL, {
        method: "POST",
        headers: {
          "Content-Length": e.data.size.toString(),
          "Authorization": accessToken
        },
        body: formData
      });
    } catch (err) {
      if (err instanceof Error) {
        await loggerInstance.log(LogLevel.Error, err.message);
      }
    }
  };
}
