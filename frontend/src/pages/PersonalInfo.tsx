import { useForm } from "react-hook-form";
import { FormErrorMessage, useColorModeValue } from "@chakra-ui/react";
import type { SubmitHandler } from "react-hook-form";
import { yupResolver } from "@hookform/resolvers/yup";
import { PersonalInfoSchema } from "@/schema";
import { useAppDispatch, useAppSelector } from "@/store";
import { recordPersonalInfo } from "@/store/slices/personalInfoSlice";
import { useNavigate } from "react-router-dom";
import { useEffect } from "react";
import {
  Box,
  FormControl,
  FormLabel,
  Input,
  Heading,
  Button
} from "@chakra-ui/react";
import Layout from "@/components/Layout";
import ThemeButton from "@/components/ThemeButton";
import type { InitialState as PersonalInfoState } from "@/store/slices/personalInfoSlice/types";

interface FormValues {
  stdNo: string;
  programmingExp: number;
  programmingExercise: number;
  programmingLanguage: string;
}

export default function PersonalInfo() {
  const dispatch = useAppDispatch();
  const personalInfo = useAppSelector<PersonalInfoState>(
    (state) => state.personalInfo
  );
  const navigate = useNavigate();
  const bg = useColorModeValue("white", "gray.700");
  const fg = useColorModeValue("gray.800", "gray.100");

  const {
    register,
    handleSubmit,
    formState: { errors }
  } = useForm({
    defaultValues: personalInfo,
    resolver: yupResolver(PersonalInfoSchema),
    reValidateMode: "onBlur"
  });

  const onSubmit: SubmitHandler<FormValues> = (data) => {
    dispatch(recordPersonalInfo(data));
    navigate("/instructions");
  };

  useEffect(() => {
    // eslint-disable-next-line
    console.log("errors", errors);
  }, [errors]);

  return (
    <Layout>
      <ThemeButton position="fixed" />
      <Box
        as="form"
        onSubmit={handleSubmit(onSubmit)}
        boxShadow="xl"
        p="8"
        rounded="md"
        maxW="container.sm"
        mx="auto"
        bg={bg}
        color={fg}
      >
        <Heading size="lg" textAlign="center" fontWeight="700">
          Personal Basic Info
        </Heading>

        <Box>
          {/* `eslint` is not happy with `!!foo`, need to use `Boolean` instead */}
          <FormControl id="email" mt="6" isInvalid={errors.stdNo !== undefined}>
            <FormLabel>Student Number</FormLabel>
            <Input type="text" {...register("stdNo")} autoComplete="off" />
            <FormErrorMessage>{errors?.stdNo?.message}!</FormErrorMessage>
          </FormControl>

          <FormControl
            id="email"
            mt="6"
            isInvalid={errors.programmingExp !== undefined}
          >
            <FormLabel>
              How many years have you been doing programming?
            </FormLabel>
            <Input
              type="number"
              {...register("programmingExp")}
              autoComplete="off"
            />
            <FormErrorMessage>
              {errors?.programmingExp?.message}!
            </FormErrorMessage>
          </FormControl>

          <FormControl
            id="email"
            mt="6"
            isInvalid={errors.programmingExercise !== undefined}
          >
            <FormLabel>
              How many hours in a week do you practice programming?
            </FormLabel>
            <Input
              type="number"
              {...register("programmingExercise")}
              autoComplete="off"
            />
            <FormErrorMessage>
              {errors?.programmingExercise?.message}!
            </FormErrorMessage>
          </FormControl>

          <FormControl
            id="email"
            mt="6"
            isInvalid={errors.programmingLanguage !== undefined}
          >
            <FormLabel>
              What programming languages are you familiar with (ex: Java,
              Python, C)
            </FormLabel>
            <Input
              type="text"
              {...register("programmingLanguage")}
              autoComplete="off"
            />
            <FormErrorMessage>
              {errors?.programmingLanguage?.message}!
            </FormErrorMessage>
          </FormControl>
        </Box>

        <Button
          colorScheme="blue"
          mx="auto"
          mt="6"
          type="submit"
          display="block"
        >
          Continue
        </Button>
      </Box>
    </Layout>
  );
}
