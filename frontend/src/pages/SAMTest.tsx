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
  useDisclosure
} from "@chakra-ui/react";
import Layout from "@/components/Layout";
import { ThemeButton } from "@/components/CodingTest";
import { useNavigate } from "react-router-dom";
import { useAppDispatch, useAppSelector } from "@/store";
import {
  markFirstSAMSubmitted,
  markSecondSAMSubmitted
} from "@/store/slices/sessionSlice";
import { setDeadlineAndQuestions } from "@/store/slices/editorSlice";
import { useColorModeValue } from "@/hooks/";
import { useTranslation } from "react-i18next";
import SAMRadioGroup from "@/components/SAMRadioGroup";

const ICONS = {
  arousal: import.meta.globEager("../images/arousal/arousal-*.svg"),
  pleasure: import.meta.globEager("../images/pleasure/pleasure-*.svg")
};

function getResponseOptions(
  icons: Record<string, FC<SVGProps<SVGSVGElement>>>[]
) {
  return icons.map((Icon, idx) => ({ value: idx + 1, Icon }));
}

export default function SAMTest() {
  const navigate = useNavigate();
  const dispatch = useAppDispatch();
  const { isOpen, onOpen, onClose } = useDisclosure();
  const [arousal, setArousal] = useState(1);
  const [pleasure, setPleasure] = useState(1);
  const [currentPage, setCurrentPage] = useState(0);
  const bg = useColorModeValue("white", "gray.700", "gray.800");
  const fg = useColorModeValue("gray.700", "gray.200", "gray.200");
  const fgDarker = useColorModeValue("gray.700", "gray.400", "gray.400");
  const { t } = useTranslation();
  const { firstSAMSubmitted } = useAppSelector((state) => state.session);

  function goto(kind: "next" | "prev") {
    if (kind === "prev") {
      setCurrentPage((prev) => (currentPage <= 0 ? prev : prev - 1));
    }
    if (kind === "next") {
      setCurrentPage((prev) => (currentPage >= 1 ? prev : prev + 1));
    }
  }

  function handleSubmit(e: FormEvent) {
    e.preventDefault();
  }

  function finishSAMTest() {
    if (firstSAMSubmitted) {
      dispatch(markSecondSAMSubmitted());
    } else {
      // TODO(elianiva): we should get the deadline and questions from the
      //                 server
      dispatch(
        setDeadlineAndQuestions({
          // 3 hours from now
          deadlineUtc: new Date(
            Date.now() + 3 * 60 * 60 * 1000
          ).getUTCMilliseconds(),
          questions: []
        })
      );
      dispatch(markFirstSAMSubmitted());
    }
    navigate("/coding-test");
  }

  useEffect(() => {
    document.title = "SAM Test | Spectator";
  }, []);

  return (
    <>
      <Layout>
        <Box position="fixed" left={4} top={4}>
          <ThemeButton bg={bg} fg={fg} />
        </Box>
        <Box
          as="form"
          onSubmit={handleSubmit}
          mt="20"
          p="6"
          rounded="md"
          shadow="lg"
          maxW="1300"
          mx="auto"
          bg={bg}
        >
          <Box display="inline-block">
            <Heading size="lg" color={fg} textAlign="center" mb="8">
              Self Assessment Manikin Test (SAM Test)
            </Heading>

            {currentPage === 0 && (
              <Fade in={currentPage === 0}>
                <Box>
                  <Text fontWeight="bold" color={fg} fontSize="xl" mb="2">
                    {t("translation.translations.sam_test.aroused_title")}
                  </Text>
                  <Text color={fgDarker} fontSize="lg" mb="4">
                    {t("translation.translations.sam_test.aroused_body")}
                  </Text>
                  <SAMRadioGroup
                    value={arousal}
                    onChange={(v) => setArousal(parseInt(v))}
                    name="arousal"
                    items={getResponseOptions(Object.values(ICONS.arousal))}
                  />
                </Box>
              </Fade>
            )}

            {currentPage === 1 && (
              <Fade in={currentPage === 1}>
                <Box>
                  <Text fontWeight="bold" color={fg} fontSize="xl" mb="2">
                    {t("translation.translations.sam_test.pleasure_title")}
                  </Text>
                  <Text color={fgDarker} fontSize="lg">
                    {t("translation.translations.sam_test.pleasure_body")}
                  </Text>
                  <SAMRadioGroup
                    value={pleasure}
                    onChange={(v) => setPleasure(parseInt(v))}
                    name="pleasure"
                    items={getResponseOptions(Object.values(ICONS.pleasure))}
                  />
                </Box>
              </Fade>
            )}

            <Flex justifyContent="end" mt="4" gap="4">
              {currentPage === 1 ? (
                <>
                  <Button
                    colorScheme="blue"
                    variant="outline"
                    onClick={() => goto("prev")}
                  >
                    {t("translation.translations.ui.previous")}
                  </Button>
                  <Button colorScheme="blue" variant="solid" onClick={onOpen}>
                    {t("translation.translations.ui.finish")}
                  </Button>
                </>
              ) : (
                <Button
                  colorScheme="blue"
                  variant="solid"
                  onClick={() => goto("next")}
                >
                  {t("translation.translations.ui.next")}
                </Button>
              )}
            </Flex>
          </Box>
        </Box>
      </Layout>
      <Modal isOpen={isOpen} onClose={onClose}>
        <ModalOverlay />
        <ModalContent bg={bg} color={fg}>
          <ModalHeader fontSize="2xl">
            {t("translation.translations.confirmation.title")}
          </ModalHeader>
          <ModalCloseButton />
          <ModalBody>
            <Text fontSize="lg" lineHeight="7">
              {t("translation.translations.confirmation.body")}
            </Text>
          </ModalBody>

          <ModalFooter>
            <Button
              colorScheme="blue"
              variant="outline"
              mr={3}
              onClick={onClose}
            >
              {t("translation.translations.ui.cancel")}
            </Button>
            <Button colorScheme="blue" onClick={finishSAMTest}>
              {t("translation.translations.ui.confirm")}
            </Button>
          </ModalFooter>
        </ModalContent>
      </Modal>
    </>
  );
}
