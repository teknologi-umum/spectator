import SpokeBase from "@/spoke/spokeBase";

class EventSpoke extends SpokeBase {
  public async mouseMove() {
    // TODO: implementation, or something
  }
}

export default new EventSpoke(import.meta.env.VITE_EVENT_HUB_URL);
