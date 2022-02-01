import * as SignalR from "@microsoft/signalr";

export default class SpokeBase {
  protected _hubConnection: SignalR.HubConnection;

  constructor(hubUrl: string) {
    this._hubConnection = new SignalR.HubConnectionBuilder()
      .withUrl(hubUrl)
      .withHubProtocol(new SignalR.JsonHubProtocol())
      // TODO(elianiva): might want to change this later since Debug is very
      //                 verbose
      .configureLogging(SignalR.LogLevel.Debug)
      .build();

    this._hubConnection.onclose(async () => {
      await this.start();
    });
  }

  public isDisconnected() {
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
      console.log("SignalR connected.");
    } catch (err) {
      console.log(err);
      setTimeout(this.start, 5000);
    }
  }
}
