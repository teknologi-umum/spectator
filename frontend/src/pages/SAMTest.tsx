import { useState } from "react";
import type { FC, SVGProps, FormEvent } from "react";
import {
  Box,
  Button,
  Fade,
  Flex,
  Heading,
  Text,
  useColorModeValue
} from "@chakra-ui/react";
import Layout from "@/components/Layout";
import "@/styles/samtest.css";
import ThemeButton from "@/components/ThemeButton";

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
  const [arousal, setArousal] = useState(0);
  const [pleasure, setPleasure] = useState(0);
  const [currentPage, setCurrentPage] = useState(0);
  const bg = useColorModeValue("white", "gray.700");
  const fg = useColorModeValue("gray.800", "gray.100");

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

  return (
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
                <Text fontWeight="bold" fontSize="xl" mb="2">
                  How aroused are you now?
                </Text>
                <Text fontSize="lg" mb="4">
                  Arousal refer to how aroused are you generally in the meantime
                </Text>
                {getResponseOptions(
                  Object.values(ICONS.arousal),
                  arousal,
                  setArousal
                )}
              </Box>
            </Fade>
          )}

          {currentPage === 1 && (
            <Fade in={currentPage === 1}>
              <Box>
                <Text fontWeight="bold" fontSize="xl" mb="2">
                  How pleased are you now?
                </Text>
                <Text fontSize="lg">
                  Pleasure refer to how pleased are you generally in the
                  meantime
                </Text>
                {getResponseOptions(
                  Object.values(ICONS.pleasure),
                  pleasure,
                  setPleasure
                )}
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
                <Button colorScheme="blue" variant="solid">
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
  );
}
