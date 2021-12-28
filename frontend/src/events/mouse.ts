import { calculateDirection } from "@/utils/getMouseDirection";
import { emit } from "@/events/emitter";
import type { CodingEventMouseClick, CodingEventMouseMove } from "./types";

export function mouseClickHandler(connection: unknown, questionNumber: number) {
  return async (e: MouseEvent) => {
    const data: CodingEventMouseClick = {
      // TODO(elianiva): revisit session_id
      session_id: "TBD",
      type: "coding_event_mouseclick",
      left_click: e.button === 0,
      middle_click: e.button === 1,
      right_click: e.button === 2,
      question_number: questionNumber,
      time: new Date(Date.now())
    };

    try {
      await emit(connection, data);
    } catch (err) {
      // TODO(elianiva): replace with proper logging
      console.error(err);
    }
  };
}

// TODO(elianiva): emit position and direction as a single event??
export function mouseMoveHandler(connection: unknown, questionNumber: number) {
  return async (e: MouseEvent) => {
    // only emit if it's actually moving
    const direction = calculateDirection(e);
    if (!direction) return;

    const data: CodingEventMouseMove = {
      // TODO(elianiva): revisit session_id
      session_id: "TBD",
      type: "coding_event_mousemove",
      direction: direction,
      x_position: e.pageX,
      y_position: e.pageY,
      window_width: window.innerWidth,
      window_height: window.innerHeight,
      question_number: questionNumber,
      time: new Date(Date.now())
    };

    try {
      await emit(connection, data);
    } catch (err) {
      // TODO(elianiva): replace with proper logging
      console.error(err);
    }
  };
}
