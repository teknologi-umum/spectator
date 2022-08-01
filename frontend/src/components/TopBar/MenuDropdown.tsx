import React from "react";
import type { ReactNode } from "react";
import { ChevronDownIcon } from "@/icons";
import { Menu, Button, MenuButton, MenuList, MenuButtonProps } from "@chakra-ui/react";
import { useColorModeValue } from "@/hooks";

interface MenuItemProps extends MenuButtonProps {
  buttonWidth?: string;
  bg?: string;
  fg?: string;
  dropdownWidth: string;
  title: string;
  children: ReactNode;
}

export default function MenuDropdown({
  buttonWidth = "auto",
  dropdownWidth,
  title,
  children,
  ...props
}: MenuItemProps) {
  const bg = props.bg ?? useColorModeValue("white", "gray.700", "gray.800");
  const fg = props.fg ?? useColorModeValue("gray.800", "gray.100", "gray.100");

  return (
    <Menu>
      <MenuButton
        textTransform="capitalize"
        textAlign="left"
        w={buttonWidth}
        as={Button}
        rightIcon={<ChevronDownIcon />}
        {...props}
        _hover={{ backgroundColor: bg }}
        _active={{ backgroundColor: bg }}
        color={fg}
        bg={bg}
      >
        {title}
      </MenuButton>
      <MenuList minW={dropdownWidth}>{children}</MenuList>
    </Menu>
  );
}
