import type { UploadRequest } from "@/workers/uploader.worker";
import UploaderWorker from "@/workers/uploader.worker.ts?worker";

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
