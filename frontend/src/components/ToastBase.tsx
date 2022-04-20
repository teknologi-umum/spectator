import React from "react";
import { Box } from "@chakra-ui/react";
import type { ReactNode } from "react";

interface ToastBaseProps {
  bg: string;
  fg: string;
  children: ReactNode;
  borderColor: string;
  onClick: () => void;
}

export default function ToastBase({
  bg,
  fg,
  children,
  borderColor,
  onClick
}: ToastBaseProps) {
  return (
    <Box
      bg={bg}
      color={fg}
      borderLeft="4px"
      borderColor={borderColor}
      p={4}
      borderRadius="md"
      fontSize="md"
      fontWeight="bold"
      textAlign="left"
      shadow="sm"
      onClick={onClick}
      cursor="pointer"
    >
      {children}
    </Box>
  );
}
