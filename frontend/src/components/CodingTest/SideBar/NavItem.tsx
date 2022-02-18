import React, { useEffect } from "react";
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
import {
  setCurrentQuestion,
  setCurrentQuestionNumber
} from "@/store/slices/editorSlice";
import { useTranslation } from "react-i18next";

interface NavItemProps {
  questionNumber: number;
  title: string;
  icon: ComponentWithAs<"svg", IconProps>;
}

export default function NavItem({ questionNumber, title, icon }: NavItemProps) {
  const dispatch = useAppDispatch();
  const { isCollapsed } = useAppSelector((state) => state.sideBar);
  const { currentQuestionNumber } = useAppSelector((state) => state.editor);
  const bg = useColorModeValue("teal.50", "teal.500");
  const fg = useColorModeValue("teal.700", "teal.200");
  const { t } = useTranslation();

  const isActive = currentQuestionNumber === questionNumber;

  useEffect(() => {
    setCurrentQuestion({
      questionNumber: currentQuestionNumber,
      title: t(`question.questions.${currentQuestionNumber - 1}.title`),
      instruction: t(`question.questions.${currentQuestionNumber - 1}.title`),
      templateByLanguage: t(
        `question.questions.${currentQuestionNumber - 1}.templates`
      )
    });
  }, [currentQuestionNumber]);

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
          onClick={() =>
            dispatch(setCurrentQuestionNumber(questionNumber))
          }
        >
          <MenuButton w="full">
            <Flex gap="5" alignItems="center" whiteSpace="nowrap">
              <Icon as={icon} width="1.25rem" height="1.25rem" />
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
