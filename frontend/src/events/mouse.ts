import { calculateDirection } from "@/utils/getMouseDirection";
import { emit } from "@/events/emitter";

export function mouseClickHandler(connection: unknown) {
  return async () => {
    const data = { event: "mouse", value: "clicked", timestamp: Date.now() };
    try {
      await emit(connection, data);
    } catch (err) {
      // TODO(elianiva): pake logging?
      console.error(err);
    }
  };
}

// TODO(elianiva): emit position and direction as a single event??
export function mouseMoveHandler(connection: unknown) {
  return async (e: MouseEvent) => {
    const data = {
      event: "mouse",
      value: JSON.stringify({
        x: e.pageX,
        y: e.pageY
      }),
      timestamp: Date.now()
    };

    try {
      await emit(connection, data);
    } catch (err) {
      // TODO(elianiva): pake logging?
      console.error(err);
    }

    // only emit if it's actully moving
    const direction = calculateDirection(e);
    if (direction) {
      const data = {
        event: "mouse",
        value: JSON.stringify({ direction }),
        timestamp: Date.now()
      };

      try {
        await emit(connection, data);
      } catch (err) {
        // TODO(elianiva): pake logging?
        console.error(err);
      }
    }
  };
}
