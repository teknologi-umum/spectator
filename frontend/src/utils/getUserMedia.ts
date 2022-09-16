export function getUserMedia(deviceId?: string | null) {
  return navigator.mediaDevices.getUserMedia({
    audio: false,
    video: {
      deviceId:
        deviceId !== null && deviceId !== undefined
          ? { exact: deviceId }
          : undefined,
      width: { ideal: 640, max: 1280 },
      height: { ideal: 360, max: 720 }
    }
  });
}
