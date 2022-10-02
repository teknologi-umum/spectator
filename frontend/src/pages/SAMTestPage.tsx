import React, { useEffect, useState } from "react";
import type { FC, SVGProps, FormEvent } from "react";
import {
  Box,
  Button,
  Fade,
  Flex,
  Heading,
  Text,
  Modal,
  ModalOverlay,
  ModalContent,
  ModalHeader,
  ModalFooter,
  ModalBody,
  ModalCloseButton,
  useDisclosure,
  HStack
} from "@chakra-ui/react";
import Layout from "@/components/Layout";
import { useNavigate } from "react-router-dom";
import { useAppDispatch, useAppSelector } from "@/store";
import {
  markFirstSAMSubmitted,
  markSecondSAMSubmitted
} from "@/store/slices/sessionSlice";
import { setDeadlineAndQuestions } from "@/store/slices/editorSlice";
import { useColorModeValue } from "@/hooks/";
import { useTranslation } from "react-i18next";
import SAMRadioGroup from "@/components/SAMTest/SAMRadioGroup";
import WithTour from "@/hoc/WithTour";
import { samTestTour } from "@/tours";
import { useTour } from "@reactour/tour";
import { sessionSpoke } from "@/spoke";
import { SettingsDropdown } from "@/components/Settings";

const ICONS = {
  arousal: import.meta.globEager("../images/arousal/arousal-*.svg"),
  pleasure: import.meta.globEager("../images/pleasure/pleasure-*.svg")
};

function getResponseOptions(
  icons: Record<string, FC<SVGProps<SVGSVGElement>>>[]
) {
  return icons.map((Icon, idx) => ({ value: idx + 1, Icon }));
}

enum Page {
  FIRST,
  LAST,
}

function SAMTest() {
  const navigate = useNavigate();
  const dispatch = useAppDispatch();
  const { t } = useTranslation("translation", {
    keyPrefix: "translations"
  });
  const { isOpen, onOpen, onClose } = useDisclosure();
  const { setIsOpen, setCurrentStep } = useTour();

  const bg = useColorModeValue("white", "gray.700", "gray.800");
  const fg = useColorModeValue("gray.700", "gray.200", "gray.200");
  const fgDarker = useColorModeValue("gray.700", "gray.400", "gray.400");

  const [currentPage, setCurrentPage] = useState(Page.FIRST);
  const [arousedLevel, setArousedLevel] = useState(1);
  const [pleasedLevel, setPleasedLevel] = useState(1);

  const { accessToken, firstSAMSubmitted, tourCompleted } = useAppSelector(
    (state) => state.session
  );
  const { examResult } = useAppSelector((state) => state.examResult);

  const samTranslationKey = firstSAMSubmitted
    ? "sam_test_after"
    : "sam_test_before";

  function goto(kind: "next" | "prev") {
    if (kind === "prev") {
      setCurrentPage((prev) => (currentPage <= 0 ? prev : prev - 1));
    }
    if (kind === "next") {
      setCurrentPage((prev) => (currentPage >= 1 ? prev : prev + 1));
    }
  }

  async function finishSAMTest() {
    if (accessToken === null) return;

    // submit SAM test before starting the exam
    if (!firstSAMSubmitted) {
      await sessionSpoke.submitBeforeExamSAM({
        accessToken,
        arousedLevel,
        pleasedLevel
      });
      const exam = await sessionSpoke.startExam({ accessToken });
      dispatch(
        setDeadlineAndQuestions({
          deadlineUtc: Number(exam.deadline),
          questions: []
        })
      );
      dispatch(markFirstSAMSubmitted());
      navigate("/video-test");
      return;
    }

    // submit SAM test after finishing the exam
    if (examResult !== null) {
      await sessionSpoke.submitAfterExamSAM({
        accessToken,
        arousedLevel,
        pleasedLevel
      });
      dispatch(markSecondSAMSubmitted());
      navigate("/fun-fact");
      return;
    }

    // submit SAM test after each question
    await sessionSpoke.submitSolutionSAM({
      accessToken,
      arousedLevel,
      pleasedLevel
    });
    navigate("/coding-test");
  }

  useEffect(() => {
    document.title = "SAM Test | Spectator";
    if (tourCompleted.samTest) return;
    setIsOpen(true);
  }, []);

  return (
    <>
      <Layout display="flex">
        <SettingsDropdown />
        <Box
          as="form"
          onSubmit={(e: FormEvent) => e.preventDefault()}
          mt="20"
          p="6"
          rounded="md"
          shadow="lg"
          maxW="1300"
          height="full"
          mx="auto"
          bg={bg}
        >
          <Box display="inline-block">
            <Heading size="lg" color={fg} textAlign="center" mb="8">
              Self Assessment Manikin Test (SAM Test)
            </Heading>

            {currentPage === Page.FIRST && (
              <Fade in={currentPage === 0}>
                <Box data-tour="step-1">
                  <Text fontWeight="bold" color={fg} fontSize="xl" mb="2">
                    {t(`${samTranslationKey}.aroused_title`)}
                  </Text>
                  <Text
                    color={fgDarker}
                    fontSize="lg"
                    mb="4"
                    dangerouslySetInnerHTML={{
                      __html: t(`${samTranslationKey}.aroused_body`)
                    }}
                  ></Text>
                  <SAMRadioGroup
                    value={arousedLevel}
                    onChange={(v) => setArousedLevel(parseInt(v))}
                    name="arousedLevel"
                    items={getResponseOptions(Object.values(ICONS.arousal))}
                  />
                </Box>
              </Fade>
            )}

            {currentPage === Page.LAST && (
              <Fade in={currentPage === Page.LAST}>
                <Box>
                  <Text fontWeight="bold" color={fg} fontSize="xl" mb="2">
                    {t(`${samTranslationKey}.pleasure_title`)}
                  </Text>
                  <Text
                    color={fgDarker}
                    fontSize="lg"
                    dangerouslySetInnerHTML={{
                      __html: t(`${samTranslationKey}.pleasure_body`)
                    }}
                  ></Text>
                  <SAMRadioGroup
                    value={pleasedLevel}
                    onChange={(v) => setPleasedLevel(parseInt(v))}
                    name="pleasedLevel"
                    items={getResponseOptions(Object.values(ICONS.pleasure))}
                  />
                </Box>
              </Fade>
            )}

            <Flex justifyContent="end" mt="4" gap="4">
              {currentPage === Page.LAST ? (
                <>
                  <Button
                    colorScheme="blue"
                    variant="outline"
                    onClick={() => goto("prev")}
                  >
                    {t("ui.previous")}
                  </Button>
                  <Button colorScheme="blue" variant="solid" onClick={onOpen}>
                    {t("ui.finish")}
                  </Button>
                </>
              ) : (
                <Button
                  colorScheme="blue"
                  variant="solid"
                  onClick={() => {
                    goto("next");
                    setCurrentStep(2);
                    setIsOpen(true);
                  }}
                  data-tour="step-2"
                >
                  {t("ui.next")}
                </Button>
              )}
            </Flex>
          </Box>
        </Box>
      </Layout>
      <Modal isOpen={isOpen} onClose={onClose}>
        <ModalOverlay />
        <ModalContent bg={bg} color={fg}>
          <ModalHeader fontSize="2xl">{t("confirmation.title")}</ModalHeader>
          <ModalCloseButton />
          <ModalBody>
            <Text fontSize="lg" lineHeight="7">
              {t(
                firstSAMSubmitted
                  ? "confirmation.body-alt"
                  : "confirmation.body"
              )}
            </Text>
          </ModalBody>

          <ModalFooter>
            <HStack spacing={3}>
              <Button colorScheme="blue" variant="outline" onClick={onClose}>
                {t("ui.cancel")}
              </Button>
              <Button
                colorScheme="blue"
                onClick={finishSAMTest}
                data-tour="step-2"
              >
                {t("ui.confirm")}
              </Button>
            </HStack>
          </ModalFooter>
        </ModalContent>
      </Modal>
    </>
  );
}

export default WithTour(SAMTest, samTestTour);
