import { eventSpoke } from "@/spoke";
import { MouseButton } from "@/stub/enums";
import { MouseScrollInfo, MouseClickInfo, MouseMoveInfo } from "@/stub/input";
import { calculateDirection } from "@/utils/getMouseDirection";

const eventButtonToEnum = [
  MouseButton.LEFT_BUTTON,
  MouseButton.MIDDLE_BUTTON,
  MouseButton.RIGHT_BUTTON
];

export function mouseClickHandler(
  questionNumber: number | null,
  accessToken: string | null
) {
  return async (e: MouseEvent) => {
    if (questionNumber === null || accessToken === null) return;

    const data: MouseClickInfo = {
      accessToken: accessToken,
      questionNumber: questionNumber,
      x: e.clientX,
      y: e.clientY,
      button: eventButtonToEnum[e.button],
      time: Date.now() as unknown as bigint
    };

    try {
      await eventSpoke.mouseClicked(data);
    } catch (err) {
      // TODO(elianiva): replace with proper logging
      console.error(err);
    }
  };
}

export function mouseMoveHandler(
  questionNumber: number | null,
  accessToken: string | null
) {
  return async (e: MouseEvent) => {
    if (questionNumber === null || accessToken === null) return;

    // only emit if it's actually moving
    const direction = calculateDirection(e);
    if (direction === null) return;

    const data: MouseMoveInfo = {
      accessToken: accessToken,
      direction: direction,
      x: e.pageX,
      y: e.pageY,
      questionNumber: questionNumber,
      time: Date.now() as unknown as bigint
    };

    try {
      await eventSpoke.mouseMoved(data);
    } catch (err) {
      // TODO(elianiva): replace with proper logging
      console.error(err);
    }
  };
}

export function mouseScrollHandler(
  questionNumber: number | null,
  accessToken: string | null
) {
  return async (e: WheelEvent) => {
    if (questionNumber === null || accessToken === null) return;


    const data: MouseScrollInfo = {
      accessToken: accessToken,
      delta: e.deltaY,
      x: e.clientX,
      y: e.clientY,
      questionNumber: questionNumber,
      time: Date.now() as unknown as bigint
    };

    try {
      await eventSpoke.mouseScrolled(data);
    } catch (err) {
      console.error(err);
    }
  };
}
