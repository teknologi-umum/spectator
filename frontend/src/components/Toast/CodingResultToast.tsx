import React from "react";
import { Flex } from "@chakra-ui/react";
import { Box } from "@chakra-ui/react";
import { CheckmarkIcon, CrossIcon } from "@/icons";

interface CodingResultToastProps {
  isCorrect: boolean;
  onClick: () => void;
  bg: string;
  fg: string;
  green: string;
  red: string;
}

export default function CodingResultToast({
  isCorrect,
  onClick,
  bg,
  fg,
  green,
  red
}: CodingResultToastProps) {
  return (
    <Box
      bg={bg}
      color={fg}
      borderLeft="4px"
      borderColor={isCorrect ? green : red}
      p={4}
      borderRadius="md"
      fontSize="md"
      fontWeight="bold"
      textAlign="left"
      shadow="sm"
      onClick={onClick}
      cursor="pointer"
    >
      <Flex align="center" gap="2" color={isCorrect ? green : red}>
        {isCorrect ? (
          <>
            <CheckmarkIcon width="1.25rem" height="1.25rem" /> Correct answer!
          </>
        ) : (
          <>
            <CrossIcon width="1.25rem" height="1.25rem" /> Wrong answer!
          </>
        )}
      </Flex>
    </Box>
  );
}
