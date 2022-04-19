import React from "react";
import { Box } from "@chakra-ui/react";
import type { ReactNode } from "react";
import { useColorModeValue } from "@/hooks";

interface ToastBaseProps {
  children: ReactNode;
  borderColor: string;
  onClick: () => void;
}

export default function ToastBase({
  children,
  borderColor,
  onClick
}: ToastBaseProps) {
  const toastBg = useColorModeValue("white", "gray.600", "gray.700");
  const toastFg = useColorModeValue("gray.700", "gray.600", "gray.700");
  
  return (
    <Box
      bg={toastBg}
      color={toastFg}
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
