import { FormEvent, useState } from "react";
//import { recordPretest } from '../../store/actions/pretestAction'
//import Select from 'react-select';
import "./samtest.css";
import { useNavigate } from "react-router-dom";
import { Box, Button, Fade, Flex, Heading, Text } from "@chakra-ui/react";

function getResponseOptions(
  kind: string,
  state: number,
  setState: React.Dispatch<React.SetStateAction<number>>
) {
  return (
    <Flex wrap="wrap" gap="4" mt="4">
      {Array(9)
        .fill(0)
        .map((_, idx) => (
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
            <img
              src={`/sam/${kind}/${kind}-${idx + 1}.png`}
              alt={`${kind}-${idx + 1}`}
            />
          </label>
        ))}
    </Flex>
  );
}

interface SAMTestProps {
  nextQuestion: number;
}
export default function SAMTest({ nextQuestion }: SAMTestProps) {
  const navigate = useNavigate();
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

  function finishQuestions() {
    navigate("/test");
  }

  function handleSubmit(e: FormEvent) {
    e.preventDefault();

    if (nextQuestion <= 6) {
      navigate(`/Q${nextQuestion}`);
    } else {
      finishQuestions();
    }
  }

  return (
    <Box
      as="form"
      onSubmit={handleSubmit}
      mt="20"
      p="6"
      rounded="md"
      shadow="lg"
      maxW="1300"
      mx="auto"
      bg="white"
    >
      <Heading size="lg" color="gray.700" textAlign="center" mb="8">
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
            {getResponseOptions("arousal", arousal, setArousal)}
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
              Pleasure refer to how pleased are you generally in the meantime
            </Text>
            {getResponseOptions("pleasure", pleasure, setPleasure)}
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
  );
}
