import SpokeBase from "@/spoke/spokeBase";
import { KeystrokeRequest, MouseClickRequest, MouseMoveRequest, MouseScrollRequest } from "@/stub/events";

// TODO(elianiva): replace with the proper method names
//                 currently these are just some fake methods to
//                 make it easier to replace later
class EventSpoke extends SpokeBase {
  public async mouseScrolled(request: MouseScrollRequest) {
    await this._startIfDisconnected();
    return this._hubConnection.invoke("MouseScrollAsync", request);
  }

  public async mouseClicked(request: MouseClickRequest) {
    await this._startIfDisconnected();
    return this._hubConnection.invoke("MouseClickAsync", request);
  }

  public async mouseMoved(request: MouseMoveRequest) {
    await this._startIfDisconnected();
    return this._hubConnection.invoke("MouseMoveAsync", request);
  }

  public async keystroke(request: KeystrokeRequest) {
    await this._startIfDisconnected();
    return this._hubConnection.invoke("KeystrokeAsync", request);
  }
}

export default new EventSpoke(import.meta.env.VITE_EVENT_HUB_URL);
