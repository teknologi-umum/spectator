import { useEffect, useState } from "react";

export function useVideoSources(options: { isAllowed: boolean }) {
  const [videoSources, setVideoSources] = useState<InputDeviceInfo[]>([]);

  useEffect(() => {
    // only try to get the source list when it's allowed
    // because the browser will still give us the list when it's not allowed but the information is useless
    if (!options.isAllowed) return;

    (async () => {
      const sources = await navigator.mediaDevices.enumerateDevices();
      setVideoSources(
        sources.filter(
          (source): source is InputDeviceInfo => source.kind === "videoinput"
        )
      );
    })();
  }, [options.isAllowed]);

  return videoSources;
}
