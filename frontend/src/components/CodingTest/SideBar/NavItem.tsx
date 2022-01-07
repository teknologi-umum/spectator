import React from "react";
import {
  Flex,
  Text,
  Link,
  Menu,
  MenuButton,
  Icon,
  Fade,
  useColorModeValue
} from "@chakra-ui/react";
import type { ComponentWithAs, IconProps } from "@chakra-ui/react";
import { useAppDispatch, useAppSelector } from "@/store";
import { setQuestion } from "@/store/slices/questionSlice";

interface NavItemProps {
  questionNo: number;
  title: string;
  icon: ComponentWithAs<"svg", IconProps>;
}

export default function NavItem({ questionNo, title, icon }: NavItemProps) {
  const dispatch = useAppDispatch();
  const { isCollapsed } = useAppSelector((state) => state.sideBar);
  const { currentQuestion } = useAppSelector((state) => state.question);
  const bg = useColorModeValue("teal.50", "teal.500");
  const fg = useColorModeValue("teal.700", "teal.200");

  const isActive = currentQuestion === questionNo;

  return (
    <Flex w="100%" flexDirection="column" alignItems="flex-start">
      <Menu placement="right">
        <Link
          backgroundColor={isActive ? bg : "transparent"}
          color={isActive ? fg : "inherit"}
          px="3"
          py="2"
          borderRadius="md"
          w="full"
          _hover={{ textDecoration: "none" }}
          onClick={() => dispatch(setQuestion(questionNo))}
        >
          <MenuButton w="full">
            <Flex gap="5" alignItems="center" whiteSpace="nowrap">
              <Icon as={icon} />
              <Fade in={!isCollapsed}>
                <Text
                  fontWeight={isActive ? "bold" : "normal"}
                  whiteSpace="nowrap"
                  textOverflow="ellipsis"
                  overflow="hidden"
                >
                  {title}
                </Text>
              </Fade>
            </Flex>
          </MenuButton>
        </Link>
      </Menu>
    </Flex>
  );
}
