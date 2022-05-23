import { VIDEO_STREAM_URL } from "@/constants";
import { loggerInstance } from "@/spoke/logger";
import { LogLevel } from "@microsoft/signalr";

type Fetcher = () => Promise<void>;

class RequestQueue {
  private readonly _queue: Array<Fetcher> = [];

  public push(fetcher: Fetcher) {
    this._queue.push(fetcher);
  }

  public async execute() {
    while (this._queue.length > 0) {
      const fetcher = this._queue.shift();
      if (fetcher) {
        // eslint-disable-next-line no-await-in-loop
        await fetcher();
      }
    }
    setTimeout(this.execute.bind(this), 0);
  }
}

const queue = new RequestQueue();
queue.execute();

export interface UploadRequest {
  blob: Blob;
  size: string;
  accessToken: string;
  startedAt: number;
}

const worker = self as unknown as Worker;
worker.onmessage = (e: MessageEvent<UploadRequest>) => {
  const fetcher = async () => {
    const formData = new FormData();
    formData.append("startedAt", e.data.startedAt.toString());
    formData.append("stoppedAt", Date.now().toString());
    formData.append("file", e.data.blob);

    try {
      await fetch(VIDEO_STREAM_URL, {
        method: "POST",
        headers: {
          "Content-Length": e.data.size,
          Authorization: e.data.accessToken
        },
        body: formData
      });
    } catch (err) {
      if (err instanceof Error) {
        await loggerInstance.log(LogLevel.Error, err.message);
      }
    }
  };

  queue.push(fetcher);
};
