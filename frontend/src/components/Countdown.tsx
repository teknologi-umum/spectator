import { useState, useEffect } from "react";

function toReadableTime(seconds: number): string {
  const s = Math.floor(seconds % 60);
  const m = Math.floor(seconds / 60 % 60);
  const h = Math.floor(seconds / (60 * 60) % 24);
  return [h, m.toString().padStart(2, "0"), s.toString().padStart(2, "0")].join(
    ":"
  );
}

interface CountdownProps {
  duration: number;
}

export default function Countdown({ duration }: CountdownProps) {
  const [time, setTime] = useState(duration);

  useEffect(() => {
    const timer = setInterval(() => setTime((prev) => prev - 1), 1000);
    return () => clearInterval(timer);
  });

  return (
    <div>
      <p>{toReadableTime(time)}</p>
    </div>
  );
}
