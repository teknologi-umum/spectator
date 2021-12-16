import { useNavigate } from "react-router-dom";
import arousal from "./Arousal-copy.png";
import { Heading, Text, Container, Button, Image } from "@chakra-ui/react";

export default function Instructions() {
  const navigate = useNavigate();
  return (
    <Container maxW="container.md" bg="white" p="6" rounded="md" shadow="md">
      <Heading size="lg" textAlign="center" mb="4" color="gray.700">
        General Instructions
      </Heading>
      <Text fontSize="18" lineHeight="8" color="gray.700">
        This experiment contains of two part. The first part is SAM Test and the
        second part is coding test. SAM Test is a self assessment test to
        measure your emotion. This test contains three question that will assess
        your emotion and will take time less than 5 minutes. The second test is
        coding test that consist of six programming questions. In this part you
        have to answer the questions and finish it within 90 minutes.
      </Text>

      <Heading size="md" mt="6" mb="4">
        1. SAM Test
      </Heading>
      <Text fontSize="18" lineHeight="8" color="gray.700">
        In this part there will be three question that you are required to
        answer by clicking one out of nine images available. The images
        represent how do you feel about the question. The image that is located
        on the far left indicates that you are very disagree with the statement
        and the image on the far right indicates that you are very agree with
        the statement meanwhile whe 5th image indicates that you are neutral
        towards the statement. You are required to fill this test two times. The
        first is before you take the test and the second test is after you
        finish programming test. The first SAM Test will be asking your current
        emotion meanwhile the second SAM Test will be asking your emotion during
        programming test.
      </Text>
      <Image src={arousal} alt={arousal} my="2" />
      <Text as="label" fontSize="sm" lineHeight="8" color="gray.600">
        SAM Test Example
      </Text>

      <Heading size="md" mt="6" mb="4">
        2. Programming Test
      </Heading>
      <Text fontSize="18" lineHeight="8" color="gray.700">
        In this part there will be six programming questions that you are
        required to answer within 90 minutes by using java programming language.
        You are not allowed to search the answer somewhere else or get help from
        other people. If you are not able to answer the questions, you are
        allowed to left the page empty and go to the next question. You are
        allowed to go back to previous questions when there is still time left.
      </Text>
      <Text as="label" fontSize="sm" lineHeight="8" color="gray.600">
        P.S: This test will not affect your mark
      </Text>

      <Button
        type="submit"
        colorScheme="blue"
        onClick={() => navigate("/sam-test")}
        display="block"
        mx="auto"
        mt="6"
      >
        Begin Test
      </Button>
    </Container>
  );
}
