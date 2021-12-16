import { useForm, SubmitHandler } from "react-hook-form";
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

interface FormValues {
  stdNo: string;
  programmingExp: number;
  programmingExercise: number;
  programmingLanguage: string;
}

export default function PersonalInfo() {
  const dispatch = useAppDispatch();
  const personalInfo = useAppSelector((state) => state.personalInfo);
  const navigate = useNavigate();

  const {
    register,
    handleSubmit,
    formState: { errors }
  } = useForm({
    defaultValues: personalInfo,
    resolver: yupResolver(PersonalInfoSchema)
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
      <Box
        as="form"
        onSubmit={handleSubmit(onSubmit)}
        boxShadow="xl"
        p="8"
        rounded="md"
        bg="white"
        maxW="container.sm"
        mx="auto"
      >
        <Heading size="lg" textAlign="center" fontWeight="700" color="gray.700">
          Personal Basic Info
        </Heading>

        <FormControl id="email" mt="6" isRequired>
          <FormLabel>Student Number</FormLabel>
          <Input type="text" {...register("stdNo")} autoComplete="off" />
        </FormControl>

        <FormControl id="email" mt="6" isRequired>
          <FormLabel>How many years have you been doing programming?</FormLabel>
          <Input
            type="number"
            {...register("programmingExp")}
            autoComplete="off"
          />
        </FormControl>

        <FormControl id="email" mt="6" isRequired>
          <FormLabel>
            How many hours in a week do you practice programming?
          </FormLabel>
          <Input
            type="number"
            {...register("programmingExercise")}
            autoComplete="off"
          />
        </FormControl>

        <FormControl id="email" mt="6" isRequired>
          <FormLabel>
            What programming languages are you familiar with (ex: Java, Python,
            C)
          </FormLabel>
          <Input
            type="number"
            {...register("programmingLanguage")}
            autoComplete="off"
          />
        </FormControl>

        <Button
          colorScheme="blue"
          mx="auto"
          mt="6"
          display="block"
          onClick={() => {
            // FIXME: proper navigation logic
            navigate("/instructions");
          }}
        >
          Continue
        </Button>
      </Box>
    </Layout>
  );
}
