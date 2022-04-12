import SpokeBase from "@/spoke/spokeBase";
import {
  KeystrokeInfo,
  MouseClickInfo,
  MouseMoveInfo,
  MouseScrollInfo
} from "@/stub/input";

// TODO(elianiva): replace with the proper method names
//                 currently these are just some fake methods to
//                 make it easier to replace later
class EventSpoke extends SpokeBase {
  public async mouseScrolled(request: MouseScrollInfo) {
    await this._startIfDisconnected();
    return this._hubConnection.invoke("MouseScrollAsync", request);
  }

  public async mouseClicked(request: MouseClickInfo) {
    await this._startIfDisconnected();
    return this._hubConnection.invoke("MouseClickAsync", request);
  }

  public async mouseMoved(request: MouseMoveInfo) {
    await this._startIfDisconnected();
    return this._hubConnection.invoke("MouseMoveAsync", request);
  }

  public async keystroke(request: KeystrokeInfo) {
    await this._startIfDisconnected();
    return this._hubConnection.invoke("KeystrokeAsync", request);
  }
}

export default new EventSpoke(import.meta.env.VITE_EVENT_HUB_URL);
