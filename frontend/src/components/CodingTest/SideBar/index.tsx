import React from "react";
import { Box, Divider, Flex, IconButton } from "@chakra-ui/react";
import type { ComponentWithAs, IconProps } from "@chakra-ui/react";
import { useAppDispatch, useAppSelector } from "@/store";
import { toggleSideBar } from "@/store/slices/codingTestSlice";
import NavItem from "./NavItem";
import {
  DiamondIcon,
  EqualIcon,
  HamburgerIcon,
  PyramidIcon,
  StarIcon,
  SumIcon,
  TemperatureIcon
} from "@/icons";

interface SideBarProps {
  bg: string;
  fg: string;
}

const icons = [
  StarIcon,
  TemperatureIcon,
  EqualIcon,
  SumIcon,
  DiamondIcon,
  PyramidIcon
];

export default function SideBar({ bg, fg }: SideBarProps) {
  const dispatch = useAppDispatch();
  const { isCollapsed } = useAppSelector((state) => state.codingTest);

  return (
    <Flex
      position="relative"
      h="full"
      w={isCollapsed ? "65px" : "200px"}
      bg={bg}
      color={fg}
      flexShrink={0}
      boxShadow="md"
      transition="width 300ms ease"
    >
      <Flex
        p="3"
        flexDirection="column"
        alignItems="flex-start"
        gap="5"
        as="nav"
        w="full"
      >
        <IconButton
          aria-label="Toggle SideBar"
          background="none"
          icon={<HamburgerIcon />}
          onClick={() => {
            dispatch(toggleSideBar());
          }}
          data-tour="sidebar-step-1"
        />

        <Divider />

        <Box w="full" data-tour="sidebar-step-2">
          {icons.map((icon, idx) => (
            <NavItem
              key={idx}
              questionNumber={idx + 1}
              icon={icon as ComponentWithAs<"svg", IconProps>}
              title={`Challenge ${idx + 1}`}
            />
          ))}
        </Box>
      </Flex>
    </Flex>
  );
}
