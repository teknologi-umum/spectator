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
          Token: accessToken
        },
        body: formData
      });
    } catch (err) {
      // TODO(elianiva): proper logging
      if (err instanceof Error) {
        console.error(err.message);
      }
    }
  };
}
