/**
 * `useSignalR` will give you a connection to a SignalR hub
 * @param hubUrl - The hub URL
 */
export function useSignalR(hubUrl: string) {
  // kira kira implementasinya nanti kaya gini:
  // const connection = new signalR.HubConnectionBuilder()
  //   .withUrl(hubUrl)
  //   .configureLogging(signalR.LogLevel.Information)
  //   .build();
  // tapi karena servernya belom ada, jadi implement nanti aja
  const connection = "not an actual implementation";

  return connection;
}
