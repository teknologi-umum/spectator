import * as SignalR from "@microsoft/signalr";
import { LocaleInfo, SessionReply } from "@/stub/session";

const connection = new SignalR.HubConnectionBuilder()
  .withUrl("http://localhost:5000/hubs/session")
  .withHubProtocol(new SignalR.JsonHubProtocol())
  .configureLogging(SignalR.LogLevel.Information)
  .build();

async function start() {
  try {
    await connection.start();
    console.log("SignalR connected.");
  } catch (err) {
    console.log(err);
    setTimeout(start, 5000);
  }
}

connection.onclose(async () => {
  await start();
});

export async function startSession(localeInfo: LocaleInfo): Promise<SessionReply> {
  await start();
  return connection.invoke("StartSessionAsync", localeInfo);
}
