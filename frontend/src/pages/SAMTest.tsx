import { useEffect, useState } from "react";
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
import { setJwt } from "@/store/slices/jwtSlice";
import { useColorModeValue } from "@/hooks/";
import { withPublic } from "@/hoc";

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

function SAMTest() {
  const navigate = useNavigate();
  const dispatch = useAppDispatch();
  const { isOpen, onOpen, onClose } = useDisclosure();
  const [arousal, setArousal] = useState(0);
  const [pleasure, setPleasure] = useState(0);
  const [currentPage, setCurrentPage] = useState(0);
  const bg = useColorModeValue("white", "gray.700", "gray.800");
  const fg = useColorModeValue("gray.700", "gray.200", "gray.200");
  const fgDarker = useColorModeValue("gray.700", "gray.400", "gray.400");
  
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
                    How aroused are you now?
                  </Text>
                  <Text color={fgDarker} fontSize="lg" mb="4">
                    Arousal refer to how aroused are you generally in the
                    meantime
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
                    How pleased are you now?
                  </Text>
                  <Text color={fgDarker} fontSize="lg">
                    Pleasure refer to how pleased are you generally in the
                    meantime
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
                    Previous
                  </Button>
                  <Button colorScheme="blue" variant="solid" onClick={onOpen}>
                    Finish
                  </Button>
                </>
              ) : (
                <Button
                  colorScheme="blue"
                  variant="solid"
                  onClick={() => goto("next")}
                >
                  Next
                </Button>
              )}
            </Flex>
          </Box>
        </Box>
      </Layout>
      <Modal isOpen={isOpen} onClose={onClose}>
        <ModalOverlay />
        <ModalContent>
          <ModalHeader fontSize="2xl">Confirmation</ModalHeader>
          <ModalCloseButton />
          <ModalBody>
            <Text fontSize="lg" lineHeight="7">
              In this coding test, <b>ALL</b> of your mouse and keyboard
              activity will be recorded for data collecting purpose. We will
              also need the camera permission to record your movement. By
              pressing the confirm button, you&apos;re fully agree with these
              conditions and given us your permissions.
            </Text>
          </ModalBody>

          <ModalFooter>
            <Button
              colorScheme="blue"
              variant="outline"
              mr={3}
              onClick={onClose}
            >
              Cancel
            </Button>
            <Button
              colorScheme="blue"
              onClick={() => {
                const jwt = getJwt();
                dispatch(setJwt(jwt));
                navigate("/coding-test");
              }}
            >
              Confirm
            </Button>
          </ModalFooter>
        </ModalContent>
      </Modal>
    </>
  );
}

export default withPublic(SAMTest);
