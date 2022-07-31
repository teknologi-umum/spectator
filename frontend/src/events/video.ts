import type { UploadRequest } from "@/events/video.worker";

const worker = new Worker(new URL("./video.worker.ts", import.meta.url), {
  type: "module",
});

export function uploadVideo(accessToken: string, startedAt: number) {
  return (e: BlobEvent) => {
    const data: UploadRequest = {
      blob: e.data,
      accessToken: accessToken,
      startedAt: startedAt,
      stoppedAt: Math.round(e.timeStamp),
    };
    worker.postMessage(data);
  };
}
