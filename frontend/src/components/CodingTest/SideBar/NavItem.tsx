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
  const { isCollapsed } = useAppSelector((state) => state.codingTest);
  const {
    currentQuestionNumber,
    lockedToCurrentQuestion,
    snapshotByQuestionNumber
  } = useAppSelector((state) => state.editor);
  const bg = useColorModeValue("teal.50", "teal.500");
  const fg = useColorModeValue("teal.700", "teal.200");
  const { t } = useTranslation("question", {
    keyPrefix: "questions"
  });

  const isActive = currentQuestionNumber === questionNumber;
  const isLocked =
    lockedToCurrentQuestion || !snapshotByQuestionNumber[0].submissionAccepted;

  useEffect(() => {
    setCurrentQuestion({
      questionNumber: currentQuestionNumber,
      title: t(`${currentQuestionNumber}.title`),
      instruction: t(`${currentQuestionNumber}.title`),
      templateByLanguage: t(`${currentQuestionNumber}.templates`)
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
          css={{ opacity: isLocked ? 0.5 : 1 }}
          _hover={{ textDecoration: "none" }}
          onClick={() => {
            // don't allow the user to navigate to other question
            // when they're submitting/testing their submission for the current question
            if (isLocked) return;
            dispatch(setCurrentQuestionNumber(questionNumber));
          }}
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
