import type { UploadRequest } from "@/events/video.worker";
import UploaderWorker from "@/events/video.worker.ts?worker";

const worker = new UploaderWorker();

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
