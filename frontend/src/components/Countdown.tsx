import { useState, useEffect } from "react";
import { Box, Text } from "@chakra-ui/react";

function toReadableTime(seconds: number): string {
  const s = Math.floor(seconds % 60);
  const m = Math.floor((seconds / 60) % 60);
  const h = Math.floor((seconds / (60 * 60)) % 24);
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
    <Box
      position="fixed"
      right="1rem"
      top="1rem"
      py="4"
      px="6"
      rounded="md"
      bg="gray.700"
    >
      <Text
        fontWeight="700"
        fontSize="xl"
        letterSpacing="wider"
        fontFamily="Inter"
        color="white"
      >
        {toReadableTime(time)}
      </Text>
    </Box>
  );
}
