import { useNavigate } from "react-router-dom";
import { Heading, Text, Container, Button, Image, useColorModeValue } from "@chakra-ui/react";
import Layout from "@/components/Layout";

interface theme {
  background: any,
  color: any
}

export default function Instructions({background, color}: theme) {
  const navigate = useNavigate();

  const textColor = useColorModeValue("gray.600", "gray.400")

  return (
    <Layout>
      <Container maxW="container.md" bg={background} p="6" rounded="md" shadow="md">
        <Heading size="lg" textAlign="center" mb="4" color={color}>
          General Instructions
        </Heading>
        <Text fontSize="18" lineHeight="8" color={textColor}>
          This experiment contains of two part. The first part is SAM Test and
          the second part is coding test. SAM Test is a self assessment test to
          measure your emotion. This test contains three question that will
          assess your emotion and will take time less than 5 minutes. The second
          test is coding test that consist of six programming questions. In this
          part you have to answer the questions and finish it within 90 minutes.
        </Text>

        <Heading size="md" mt="6" mb="4" color={color}>
          1. SAM Test
        </Heading>
        <Text fontSize="18" lineHeight="8" color={textColor}>
          In this part there will be three question that you are required to
          answer by clicking one out of nine images available. The images
          represent how do you feel about the question. The image that is
          located on the far left indicates that you are very disagree with the
          statement and the image on the far right indicates that you are very
          agree with the statement meanwhile whe 5th image indicates that you
          are neutral towards the statement. You are required to fill this test
          two times. The first is before you take the test and the second test
          is after you finish programming test. The first SAM Test will be
          asking your current emotion meanwhile the second SAM Test will be
          asking your emotion during programming test.
        </Text>
        <Image src="/sam/arousal/arousal.svg" alt="arousal" mt="2" />
        <Text as="label" fontSize="sm" lineHeight="8" color={textColor}>
          SAM Test Example
        </Text>

        <Heading size="md" mt="6" mb="4" color={color}>
          2. Programming Test
        </Heading>
        <Text fontSize="18" lineHeight="8" color={textColor}>
          In this part there will be six programming questions that you are
          required to answer within 90 minutes by using java programming
          language. You are not allowed to search the answer somewhere else or
          get help from other people. If you are not able to answer the
          questions, you are allowed to left the page empty and go to the next
          question. You are allowed to go back to previous questions when there
          is still time left.
        </Text>
        <Text as="label" fontSize="sm" lineHeight="8" color={textColor}>
          P.S: This test will not affect your mark
        </Text>

        <Button
          type="submit"
          backgroundColor="blue.400"
          color="white"
          onClick={() => navigate("/sam-test")}
          display="block"
          mx="auto"
          mt="6"
        >
          Begin Test
        </Button>
      </Container>
    </Layout>
  );
}
