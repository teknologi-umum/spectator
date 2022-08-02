export function getUserMedia(deviceId?: string) {
  return navigator.mediaDevices.getUserMedia({
    audio: false,
    video: {
      deviceId: deviceId ? { exact: deviceId } : undefined,
      width: 640,
      height: 360
    }
  });
}
