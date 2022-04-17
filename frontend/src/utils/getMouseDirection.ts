// see: https://codepen.io/ronnygunawan/pen/wvrojXg
// obviously i'm not smart enough to figure these out

import { MouseDirection } from "@/stub/enums";

interface MovementSample {
  t: number;
  x: number;
  y: number;
}

const samples: MovementSample[] = [];
const MAX_SAMPLES = 20;
const TIME_WINDOW = 1000; // 1 second
const EMA_FACTOR = 0.6;

function getDirection(vx: number, vy: number): MouseDirection {
  const isHorizontal = Math.abs(vx) > Math.abs(vy);
  const isRight = vx > 0;
  const isDown = vy > 0;

  /* prettier-ignore */
  return isHorizontal
    ? (isRight ? MouseDirection.RIGHT : MouseDirection.LEFT)
    : (isDown ? MouseDirection.DOWN : MouseDirection.UP);
}

// save direction history
let direction: MouseDirection = MouseDirection.STOP;

/**
 * `calculateDirection` will return mouse move direction.
 * It collects data in 1 second and will return null if it doesn't have enough data to calculate.
 * @param e - Mouse move event from the event handler
 * @return The direction of the mouse
 */
export function calculateDirection(e: MouseEvent): MouseDirection | null {
  const now = Date.now();
  samples.push({ t: now, x: e.pageX, y: e.pageY });

  // discard older entries
  while (now - samples[0].t > TIME_WINDOW || samples.length > MAX_SAMPLES) {
    samples.shift();
  }

  // not enough samples
  const prevDir = direction;
  if (samples.length <= 5) {
    direction = MouseDirection.STOP;
    if (prevDir !== direction) {
      return direction;
    }
    return null;
  }

  // EMA of velocity
  let vx = (samples[1].x - samples[0].x) / (samples[1].t - samples[0].t);
  let vy = (samples[1].y - samples[0].y) / (samples[1].t - samples[0].t);
  for (let i = 2; i < samples.length; i++) {
    vx =
      ((samples[i].x - samples[i - 1].x) / (samples[i].t - samples[i - 1].t)) *
        EMA_FACTOR +
      vx * (1 - EMA_FACTOR);
    vy =
      ((samples[i].y - samples[i - 1].y) / (samples[i].t - samples[i - 1].t)) *
        EMA_FACTOR +
      vy * (1 - EMA_FACTOR);
  }

  switch (direction) {
    case MouseDirection.STOP:
      direction = getDirection(vx, vy);
      return direction;
    case MouseDirection.RIGHT:
      if (vx < 0) {
        if (Math.abs(vy) > Math.abs(vx) * 5) {
          if (vy > 0) {
            direction = MouseDirection.DOWN;
            return direction;
          } else {
            direction = MouseDirection.UP;
            return direction;
          }
        } else {
          direction = MouseDirection.LEFT;
          return direction;
        }
      } else {
        direction = getDirection(vx, vy);
        return getDirection(vx, vy);
      }
    case MouseDirection.LEFT:
      if (vx > 0) {
        if (Math.abs(vy) > Math.abs(vx) * 5) {
          if (vy > 0) {
            direction = MouseDirection.DOWN;
            return direction;
          } else {
            direction = MouseDirection.UP;
            return direction;
          }
        } else {
          direction = MouseDirection.RIGHT;
          return direction;
        }
      } else {
        direction = getDirection(vx, vy);
        return direction;
      }
    case MouseDirection.UP:
      if (vy > 0) {
        if (Math.abs(vx) > Math.abs(vy) * 5) {
          if (vx > 0) {
            direction = MouseDirection.RIGHT;
            return direction;
          } else {
            direction = MouseDirection.LEFT;
            return direction;
          }
        } else {
          direction = MouseDirection.DOWN;
          return direction;
        }
      } else {
        direction = getDirection(vx, vy);
        return direction;
      }
    case MouseDirection.DOWN:
      if (vy < 0) {
        if (Math.abs(vx) > Math.abs(vy) * 5) {
          if (vx > 0) {
            direction = MouseDirection.RIGHT;
            return direction;
          } else {
            direction = MouseDirection.LEFT;
            return direction;
          }
        } else {
          direction = MouseDirection.UP;
          return direction;
        }
      } else {
        direction = getDirection(vx, vy);
        return direction;
      }
    default:
      if (direction !== prevDir) {
        return direction;
      }
  }

  return null;
}
