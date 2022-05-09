import { useEffect, useRef } from "react";
import type { Dispatch, SetStateAction } from "react";
import { uploadVideo } from "@/events";

interface VideoRecorderHookOptions {
  setRecordingStatus: Dispatch<
    SetStateAction<"stopped" | "started" | "unknown">
  >;
  accessToken: string | null;
}

export function useVideoRecorder({
  setRecordingStatus,
  accessToken
}: VideoRecorderHookOptions) {
  const videoStream = useRef<MediaStream | null>();
  const mediaRecorder = useRef<MediaRecorder | null>();

  useEffect(() => {
    if (accessToken === null) return;

    // We want to generate the StartedAt field here because we need it as a marker
    // for the start of the recording.
    // `MediaRecorder` splice the video stream into chunks, only the first chunk is a valid video
    // so we need to mark this first chunk.
    // A user could theoretically refresh the page and the recording would get restarted.
    // If we don't have this marker, we wouldn't know which one is the first chunk for each recording session.
    const startedAt = Date.now();

    navigator.mediaDevices
      .getUserMedia({
        audio: false,
        video: { width: 640, height: 320 }
      })
      .then((stream) => {
        videoStream.current = stream;

        mediaRecorder.current = new MediaRecorder(stream, {
          mimeType: "video/webm;codecs=vp9",
          videoBitsPerSecond: 200_000 // 0.2Mbits / sec
        });

        mediaRecorder.current.onstart = () => {
          setRecordingStatus("started");
        };
        mediaRecorder.current.onstop = () => {
          setRecordingStatus("stopped");
        };

        mediaRecorder.current.start(1000); // send blob every second
        mediaRecorder.current.ondataavailable = uploadVideo(
          accessToken,
          startedAt
        );
      });

    // stop recording when the component is unmounted
    () => {
      if (mediaRecorder.current) {
        mediaRecorder.current.stop();

        videoStream.current = null;
        mediaRecorder.current = null;
      }
    };
  }, []);
}
