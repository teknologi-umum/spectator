import { sessionSpoke, eventSpoke } from "@/spoke";
import { useAppSelector } from "@/store";

/**
 * `useSignalR` will give you a connection to a SignalR hub with the access
 * token automatically attached if it exists
 * @param hubUrl - The hub URL
 */
export function useSignalR() {
  const { accessToken } = useAppSelector((state) => state.session);

  if (accessToken !== null) {
    sessionSpoke.setAccessToken(accessToken);
    eventSpoke.setAccessToken(accessToken);
  }

  return { sessionSpoke, eventSpoke };
}
