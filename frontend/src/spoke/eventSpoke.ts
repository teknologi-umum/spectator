import type {
  CodingEventKeystroke,
  CodingEventMouseClick,
  CodingEventMouseMove
} from "@/events/types";
import SpokeBase from "@/spoke/spokeBase";

// TODO(elianiva): replace with the proper method names
//                 currently these are just some fake methods to
//                 make it easier to replace later
class EventSpoke extends SpokeBase {
  public async mouseScrolled(request: any /* TODO(elianiva): add type */) {
    await this._startIfDisconnected();
    return this._hubConnection.invoke("MouseScrolledAsync");
  }

  public async mouseClicked(request: CodingEventMouseClick) {
    await this._startIfDisconnected();
    return this._hubConnection.invoke("MouseClickedAsync", request);
  }

  public async mouseMoved(request: CodingEventMouseMove) {
    await this._startIfDisconnected();
    return this._hubConnection.invoke("MouseMovedAsync", request);
  }

  public async keyboardPressed(request: CodingEventKeystroke) {
    await this._startIfDisconnected();
    return this._hubConnection.invoke("KeyboardPressedAsync", request);
  }
}

export default new EventSpoke(import.meta.env.VITE_EVENT_HUB_URL);
