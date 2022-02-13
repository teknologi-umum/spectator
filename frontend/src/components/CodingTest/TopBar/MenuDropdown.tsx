import React from "react";
import type { ReactNode } from "react";
import { ChevronDownIcon } from "@/icons";
import { Menu, Button, MenuButton, MenuList } from "@chakra-ui/react";

interface MenuItemProps {
  buttonWidth?: string;
  bg: string;
  fg: string;
  dropdownWidth: string;
  title: string;
  children: ReactNode;
}

export default function MenuDropdown({
  buttonWidth = "auto",
  bg,
  fg,
  dropdownWidth,
  title,
  children,
  ...rest
}: MenuItemProps) {
  return (
    <Menu >
      <MenuButton
        textTransform="capitalize"
        textAlign="left"
        w={buttonWidth}
        as={Button}
        rightIcon={<ChevronDownIcon />}
        color={fg}
        bg={bg}
        _hover={{ backgroundColor: bg }}
        _active={{ backgroundColor: bg }}
        {...rest}
      >
        {title}
      </MenuButton>
      <MenuList minW={dropdownWidth}>{children}</MenuList>
    </Menu>
  );
}
