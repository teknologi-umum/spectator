import SpokeBase from "@/spoke/spokeBase";

class EventSpoke extends SpokeBase {
  public async mouseMove() {
    // TODO: implementation, or something
  }
}

export default new EventSpoke("http://localhost:5000/hubs/events");
