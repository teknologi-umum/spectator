import { store } from "@/store";
import { setState } from "@/store/slices/signalRSlice";
import * as SignalR from "@microsoft/signalr";

export default class SpokeBase {
  protected _hubConnection: SignalR.HubConnection;
  protected _accessToken: string;

  constructor(hubUrl: string) {
    this._accessToken =
      document.cookie
        .split("; ")
        .find((c) => c.startsWith("ACCESS_TOKEN="))
        ?.split("=")
        ?.at(1) ?? "";
    this._hubConnection = new SignalR.HubConnectionBuilder()
      .withUrl(hubUrl, {
        accessTokenFactory: () => this._accessToken,
        withCredentials: true
      })
      .withHubProtocol(new SignalR.JsonHubProtocol())
      // TODO(elianiva): might want to change this later since Debug is very
      //                 verbose
      .configureLogging(SignalR.LogLevel.Debug)
      .build();

    this._hubConnection.onclose(async () => {
      store.dispatch(setState(SignalR.HubConnectionState.Disconnected));
      await this.start();
    });

    this._hubConnection.onreconnecting(() => {
      store.dispatch(setState(SignalR.HubConnectionState.Reconnecting));
    });

    this._hubConnection.onreconnected(() => {
      store.dispatch(setState(SignalR.HubConnectionState.Connected));
    });
  }

  public onClose(cb: (err?: Error | undefined) => void) {
    return this._hubConnection.onclose(cb);
  }

  public setAccessToken(accessToken: string) {
    document.cookie = `ACCESS_TOKEN=${accessToken}`;
    this._accessToken = accessToken;
  }

  public isDisconnected(): boolean {
    return (
      this._hubConnection.state === SignalR.HubConnectionState.Disconnected
    );
  }

  protected async _startIfDisconnected() {
    if (this.isDisconnected()) await this.start();
  }

  public async start(): Promise<void> {
    if (!this.isDisconnected()) return;

    try {
      await this._hubConnection.start();
      store.dispatch(setState(SignalR.HubConnectionState.Connected));
      console.log("SignalR connected.");
    } catch (err) {
      console.log(err);
      setTimeout(this.start, 5000);
    }
  }
}
