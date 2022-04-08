import { eventSpoke } from "@/spoke";
import { MouseButton } from "@/stub/enums";
import { MouseClickRequest, MouseMoveRequest } from "@/stub/events";
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

    const data: MouseClickRequest = {
      accessToken: accessToken,
      questionNumber: questionNumber,
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

// TODO(elianiva): emit position and direction as a single event??
export function mouseMoveHandler(
  questionNumber: number | null,
  accessToken: string | null
) {
  return async (e: MouseEvent) => {
    if (questionNumber === null || accessToken === null) return;

    // only emit if it's actually moving
    const direction = calculateDirection(e);
    if (direction === null) return;

    const data: MouseMoveRequest = {
      accessToken: accessToken,
      direction: direction,
      xPosition: e.pageX,
      yPosition: e.pageY,
      windowWidth: window.innerWidth,
      windowHeight: window.innerHeight,
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
