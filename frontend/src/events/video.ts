import type { UploadRequest } from "@/events/video.worker";

const worker = new Worker(new URL("./video.worker.ts", import.meta.url), {
  type: "module"
});

export function uploadVideo(accessToken: string, startedAt: number) {
  return (e: BlobEvent) => {
    const data: UploadRequest = {
      blob: e.data,
      size: e.data.size.toString(),
      accessToken: accessToken,
      startedAt: startedAt
    };
    worker.postMessage(data);
  };
}
