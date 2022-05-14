import React from "react";
import type { FC, SVGProps } from "react";
import { Flex, useRadioGroup } from "@chakra-ui/react";
import { useColorModeValue } from "@/hooks";
import SAMRadio from "./SAMRadio";

interface SAMRadioGroupProps {
  name: string;
  value: number;
  onChange: (nextValue: string) => void;
  items: {
    value: number;
    Icon: Record<string, FC<SVGProps<SVGSVGElement>>>;
  }[];
}

export default function SAMRadioGroup({
  name,
  value,
  items,
  onChange
}: SAMRadioGroupProps) {
  const fgDarker = useColorModeValue("gray.600", "gray.400", "gray.500");

  const { getRootProps, getRadioProps } = useRadioGroup({
    name,
    value,
    onChange
  });
  const group = getRootProps();

  return (
    <Flex wrap="wrap" gap="3" mt="4" color={fgDarker} {...group}>
      {items.map(({ value, Icon }) => {
        const radio = getRadioProps({ value });
        return (
          <SAMRadio key={value} {...radio}>
            <Icon.ReactComponent />
          </SAMRadio>
        );
      })}
    </Flex>
  );
}
