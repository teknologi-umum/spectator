import React from "react";
import { Box, useRadio } from "@chakra-ui/react";
import { useColorModeValue } from "@/hooks";
import theme from "@/styles/themes";

// TODO(elianiva): figure out the correct props
// eslint-disable-next-line @typescript-eslint/no-explicit-any
export default function SAMRadio(props: any) {
  const border = useColorModeValue(
    theme.colors.blue[500],
    theme.colors.blue[300],
    theme.colors.blue[300]
  );
  const { getInputProps, getCheckboxProps } = useRadio(props);
  const input = getInputProps();
  const checkbox = getCheckboxProps();

  return (
    <Box as="label">
      <input {...input} />
      <Box
        {...checkbox}
        cursor="pointer"
        borderRadius="4"
        borderWidth="2px"
        borderColor="currentColor"
        _hover={{
          borderColor: border
        }}
        _checked={{
          borderColor: border,
          boxShadow: `inset 0 0 0 2px ${border}, 0 0 4px ${border}aa;`
        }}
      >
        {props.children}
      </Box>
    </Box>
  );
}
