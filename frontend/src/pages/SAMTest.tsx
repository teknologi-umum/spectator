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
import "@/styles/samtest.css";
import ThemeButton from "@/components/ThemeButton";
import { useNavigate } from "react-router-dom";
import { getJwt } from "@/utils/generateFakeJwt";
import { useAppDispatch } from "@/store";
import { setAccessToken } from "@/store/slices/sessionSlice";
import { useColorModeValue } from "@/hooks/";
import { useTranslation } from "react-i18next";

const ICONS = {
  arousal: import.meta.globEager("../images/arousal/arousal-*.svg"),
  pleasure: import.meta.globEager("../images/pleasure/pleasure-*.svg")
};

function getResponseOptions(
  icons: Record<string, FC<SVGProps<SVGSVGElement>>>[],
  state: number,
  setState: React.Dispatch<React.SetStateAction<number>>
) {
  return (
    <Flex wrap="wrap" gap="4" mt="4">
      {icons.map((Icon, idx) => {
        return (
          <label key={idx + 1}>
            <input
              style={{
                opacity: "initial",
                pointerEvents: "all"
              }}
              type="radio"
              value={idx + 1}
              onChange={() => setState(idx + 1)}
              checked={state === idx + 1}
            />
            <Icon.ReactComponent />
          </label>
        );
      })}
    </Flex>
  );
}

export default function SAMTest() {
  const navigate = useNavigate();
  const dispatch = useAppDispatch();
  const { isOpen, onOpen, onClose } = useDisclosure();
  const [arousal, setArousal] = useState(0);
  const [pleasure, setPleasure] = useState(0);
  const [currentPage, setCurrentPage] = useState(0);
  const bg = useColorModeValue("white", "gray.700", "gray.800");
  const fg = useColorModeValue("gray.700", "gray.200", "gray.200");
  const fgDarker = useColorModeValue("gray.700", "gray.400", "gray.400");
  const { t } = useTranslation();
  
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


  useEffect(() => {
    document.title = "SAM Test | Spectator";
  }, []);

  return (
    <>
      <Layout>
        <ThemeButton position="fixed" />
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
                  <Box color={fgDarker}>
                    {getResponseOptions(
                      Object.values(ICONS.arousal),
                      arousal,
                      setArousal
                    )}
                  </Box>
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
                  <Box color={fgDarker}>
                    {getResponseOptions(
                      Object.values(ICONS.pleasure),
                      pleasure,
                      setPleasure
                    )}
                  </Box>
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
          <ModalHeader fontSize="2xl">{t("translation.translations.confirmation.title")}</ModalHeader>
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
            <Button
              colorScheme="blue"
              onClick={() => {
                const jwt = getJwt();
                dispatch(setAccessToken(jwt));
                navigate("/coding-test");
              }}
            >
              {t("translation.translations.ui.confirm")}
            </Button>
          </ModalFooter>
        </ModalContent>
      </Modal>
    </>
  );
}
