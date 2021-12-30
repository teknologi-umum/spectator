// see: https://codepen.io/ronnygunawan/pen/wvrojXg
// obviously i'm not smart enough to figure these out
export enum Direction {
  STOP = "stop",
  UP = "up",
  LEFT = "left",
  RIGHT = "right",
  DOWN = "down"
}

interface MovementSample {
  t: number;
  x: number;
  y: number;
}

const samples: MovementSample[] = [];
const MAX_SAMPLES = 20;
const TIME_WINDOW = 1000; // 1 second
const EMA_FACTOR = 0.6;

function getDirection(vx: number, vy: number): Direction {
  const isHorizontal = Math.abs(vx) > Math.abs(vy);
  const isRight = vx > 0;
  const isDown = vy > 0;

  /* prettier-ignore */
  return isHorizontal
    ? (isRight ? Direction.RIGHT : Direction.LEFT)
    : (isDown ? Direction.DOWN : Direction.UP);
}

// save direction history
let direction: Direction = Direction.STOP;

/**
 * `calculateDirection` will return mouse move direction.
 * It collects data in 1 second and will return null if it doesn't have enough data to calculate.
 * @param e - Mouse move event from the event handler
 * @return The direction of the mouse
 */
export function calculateDirection(e: MouseEvent): Direction | null {
  const now = Date.now();
  samples.push({ t: now, x: e.pageX, y: e.pageY });

  // discard older entries
  while (now - samples[0].t > TIME_WINDOW || samples.length > MAX_SAMPLES) {
    samples.shift();
  }

  // not enough samples
  const prevDir = direction;
  if (samples.length <= 5) {
    direction = Direction.STOP;
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
    case Direction.STOP:
      direction = getDirection(vx, vy);
      return direction;
    case Direction.RIGHT:
      if (vx < 0) {
        if (Math.abs(vy) > Math.abs(vx) * 5) {
          if (vy > 0) {
            direction = Direction.DOWN;
            return direction;
          } else {
            direction = Direction.UP;
            return direction;
          }
        } else {
          direction = Direction.LEFT;
          return direction;
        }
      } else {
        direction = getDirection(vx, vy);
        return getDirection(vx, vy);
      }
    case Direction.LEFT:
      if (vx > 0) {
        if (Math.abs(vy) > Math.abs(vx) * 5) {
          if (vy > 0) {
            direction = Direction.DOWN;
            return direction;
          } else {
            direction = Direction.UP;
            return direction;
          }
        } else {
          direction = Direction.RIGHT;
          return direction;
        }
      } else {
        direction = getDirection(vx, vy);
        return direction;
      }
    case Direction.UP:
      if (vy > 0) {
        if (Math.abs(vx) > Math.abs(vy) * 5) {
          if (vx > 0) {
            direction = Direction.RIGHT;
            return direction;
          } else {
            direction = Direction.LEFT;
            return direction;
          }
        } else {
          direction = Direction.DOWN;
          return direction;
        }
      } else {
        direction = getDirection(vx, vy);
        return direction;
      }
    case Direction.DOWN:
      if (vy < 0) {
        if (Math.abs(vx) > Math.abs(vy) * 5) {
          if (vx > 0) {
            direction = Direction.RIGHT;
            return direction;
          } else {
            direction = Direction.LEFT;
            return direction;
          }
        } else {
          direction = Direction.UP;
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
