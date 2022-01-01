// my main purpose from creating this file was
// a playground to generate fake calculations and saved bunch of collections
// that emitted from scroll event into memory.
// but it seems handling scroll features is not that easy.

// still a bit confused about these problems below:
// 1. Number Of Scrolls, I thought this one was easy.
//    we can use setTimeout that cleared each time user scrolling.
//    if the user stop scrolling within 1 secs, just pushToCollection.
//    after that just calculate the length of collections member array.
//    The problem is, if the user scrolling on Y axis towards the positive number from its initial point,
//    once the user going to negative, is that count as 2 scrolls?

// 2. Scroll Distance, still no idea

// 3. Scroll Speed, once the problem of number of scrolls solved, this one should be easy
//    calculate a differentiation of time and distance from stop_y - start_y / stop_timestamp - start_timestamo

// fake interface
interface Collection {
  session_id: string;
  start_timestamp: number;
  start_x: number;
  start_y: number;
  stop_timestamp: number;
  stop_x: number;
  stop_y: number;
}

const collections: Collection[] = [];

export function pushToCollection(data: Collection) {
  return collections.push(data);
}

export function getNumberOfScrolls() {
  return null;
}

export function getAverageScrollDistance() {
  return null;
}

export function getAverageScrollSpeed() {
  return null;
}

export default collections;
