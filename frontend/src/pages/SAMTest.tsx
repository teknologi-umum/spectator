import { useState } from "react";
import type { FormEvent, FunctionComponent, SVGProps } from "react";
import { Box, Button, Fade, Flex, Heading, Text } from "@chakra-ui/react";
import Layout from "@/components/Layout";
import "@/styles/samtest.css";

import { ReactComponent as Arousal1 } from "@/images/arousal/arousal-1.svg";
import { ReactComponent as Arousal2 } from "@/images/arousal/arousal-2.svg";
import { ReactComponent as Arousal3 } from "@/images/arousal/arousal-3.svg";
import { ReactComponent as Arousal4 } from "@/images/arousal/arousal-4.svg";
import { ReactComponent as Arousal5 } from "@/images/arousal/arousal-5.svg";
import { ReactComponent as Arousal6 } from "@/images/arousal/arousal-6.svg";
import { ReactComponent as Arousal7 } from "@/images/arousal/arousal-7.svg";
import { ReactComponent as Arousal8 } from "@/images/arousal/arousal-8.svg";
import { ReactComponent as Arousal9 } from "@/images/arousal/arousal-9.svg";

import { ReactComponent as Pleasure1 } from "@/images/pleasure/pleasure-1.svg";
import { ReactComponent as Pleasure2 } from "@/images/pleasure/pleasure-2.svg";
import { ReactComponent as Pleasure3 } from "@/images/pleasure/pleasure-3.svg";
import { ReactComponent as Pleasure4 } from "@/images/pleasure/pleasure-4.svg";
import { ReactComponent as Pleasure5 } from "@/images/pleasure/pleasure-5.svg";
import { ReactComponent as Pleasure6 } from "@/images/pleasure/pleasure-6.svg";
import { ReactComponent as Pleasure7 } from "@/images/pleasure/pleasure-7.svg";
import { ReactComponent as Pleasure8 } from "@/images/pleasure/pleasure-8.svg";
import { ReactComponent as Pleasure9 } from "@/images/pleasure/pleasure-9.svg";

interface theme {
  background: any;
  color: any;
}

function getResponseOptions(
  icons: FunctionComponent<SVGProps<SVGSVGElement>>[],
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
            <Icon />
          </label>
        );
      })}
    </Flex>
  );
}

interface theme {
  background: any;
  color: any;
}

export default function SAMTest({ background, color }: theme) {
  const [arousal, setArousal] = useState(0);
  const [pleasure, setPleasure] = useState(0);
  const [currentPage, setCurrentPage] = useState(0);

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
      <Box
        as="form"
        onSubmit={handleSubmit}
        mt="20"
        p="6"
        rounded="md"
        shadow="lg"
        maxW="1300"
        mx="auto"
        bg={background}
      >
        <Box display="inline-block">
          <Heading size="lg" color="gray.700" textAlign="center" mb="8">
            <Text textAlign="center" color={color}>
              Self Assessment Manikin Test (SAM Test)
            </Text>
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
                  [
                    Arousal1,
                    Arousal2,
                    Arousal3,
                    Arousal4,
                    Arousal5,
                    Arousal6,
                    Arousal7,
                    Arousal8,
                    Arousal9
                  ],
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
                  [
                    Pleasure1,
                    Pleasure2,
                    Pleasure3,
                    Pleasure4,
                    Pleasure5,
                    Pleasure6,
                    Pleasure7,
                    Pleasure8,
                    Pleasure9
                  ],
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
                <Button
                  backgroundColor="blue.400"
                  color="white"
                  variant="solid"
                >
                  Finish
                </Button>
              </>
            ) : (
              <Button
                backgroundColor="blue.400"
                color="white"
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
