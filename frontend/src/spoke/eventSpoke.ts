import SpokeBase from "@/spoke/spokeBase";
import {
  KeystrokeInfo,
  MouseClickInfo,
  MouseMoveInfo,
  MouseScrollInfo,
  WindowSizeInfo
} from "@/stub/input";

class EventSpoke extends SpokeBase {
  public async mouseScrolled(request: MouseScrollInfo) {
    await this._startIfDisconnected();
    return this._hubConnection.invoke("LogMouseScrolledAsync", request);
  }

  public async mouseUp(request: MouseClickInfo) {
    await this._startIfDisconnected();
    return this._hubConnection.invoke("LogMouseUpAsync", request);
  }

  public async mouseDown(request: MouseClickInfo) {
    await this._startIfDisconnected();
    return this._hubConnection.invoke("LogMouseUpAsync", request);
  }

  public async mouseMoved(request: MouseMoveInfo) {
    await this._startIfDisconnected();
    return this._hubConnection.invoke("LogMouseMovedAsync", request);
  }

  public async keystroke(request: KeystrokeInfo) {
    await this._startIfDisconnected();
    return this._hubConnection.invoke("LogKeystrokeAsync", request);
  }

  public async windowResized(request: WindowSizeInfo) {
    await this._startIfDisconnected();
    return this._hubConnection.invoke("LogWindowSizedAsync", request);
  }
}

export default new EventSpoke(import.meta.env.VITE_EVENT_HUB_URL);
