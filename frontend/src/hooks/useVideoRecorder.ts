import { useEffect, useRef, useState } from "react";
import { uploadVideo } from "@/events";
import { getUserMedia } from "@/utils/getUserMedia";
import { useAppSelector } from "@/store";

export enum RecordingStatus {
  UNKNOWN,
  STARTED,
  STOPPED,
}

function getSupportedCodec() {
  const codecs = ["vp9", "vp8", "opus,vp8", "vp8.0", "h264", "opus"];
  for (const codec of codecs) {
    const mimeType = `video/webm;codecs=${codec}`;
    if (MediaRecorder.isTypeSupported(mimeType)) {
      return mimeType;
    }
  }
  throw new Error("No codec was supported");
}

export function useVideoRecorder(accessToken: string | null) {
  const { deviceId } = useAppSelector((state) => state.session);
  const [recordingStatus, setRecordingStatus] = useState<RecordingStatus>(
    RecordingStatus.UNKNOWN
  );
  const mediaRecorder = useRef<MediaRecorder | null>();

  useEffect(() => {
    if (accessToken === null || deviceId === null) return undefined;

    // We want to generate the StartedAt field here because we need it as a marker
    // for the start of the recording.
    // `MediaRecorder` splice the video stream into chunks, only the first chunk is a valid video
    // so we need to mark this first chunk.
    // A user could theoretically refresh the page and the recording would get restarted.
    // If we don't have this marker, we wouldn't know which one is the first chunk for each recording session.
    const startedAt = Date.now();

    let stream: MediaStream;
    (async () => {
      stream = await getUserMedia(deviceId);
      mediaRecorder.current = new MediaRecorder(stream, {
        mimeType: getSupportedCodec(),
        videoBitsPerSecond: 200_000 // 0.2Mbits / sec
      });

      mediaRecorder.current.onstart = () => {
        setRecordingStatus(RecordingStatus.STARTED);
      };
      mediaRecorder.current.onstop = () => {
        setRecordingStatus(RecordingStatus.STOPPED);
      };

      mediaRecorder.current.start(1000); // send blob every second
      mediaRecorder.current.ondataavailable = uploadVideo(
        accessToken,
        startedAt
      );
    })();

    // stop recording when the component is unmounted
    return () => {
      if (stream?.active) {
        for (const videoTrack of stream.getVideoTracks()) {
          videoTrack.stop();
        }
      }

      if (mediaRecorder.current) {
        mediaRecorder.current.stop();
        mediaRecorder.current = null;
      }
    };
  }, []);

  return recordingStatus;
}
