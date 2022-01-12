import React, { useEffect } from "react";
import { useForm } from "react-hook-form";
import type { SubmitHandler } from "react-hook-form";
import { yupResolver } from "@hookform/resolvers/yup";
import { PersonalInfoSchema } from "@/schema";
import { useAppDispatch, useAppSelector } from "@/store";
import { recordPersonalInfo } from "@/store/slices/personalInfoSlice";
import { useNavigate } from "react-router-dom";
import {
  Box,
  FormControl,
  FormLabel,
  Input,
  Heading,
  Button,
  FormErrorMessage
} from "@chakra-ui/react";
import Layout from "@/components/Layout";
import ThemeButton from "@/components/ThemeButton";
import type { InitialState as PersonalInfoState } from "@/store/slices/personalInfoSlice/types";
import { useColorModeValue } from "@/hooks";
import { withPublic } from "@/hoc";
import { useTranslation } from "react-i18next";

interface FormValues {
  studentId: string;
  programmingExp: number;
  programmingExercise: number;
  programmingLanguage: string;
}

function PersonalInfo() {
  const { t } = useTranslation();
  const dispatch = useAppDispatch();
  const personalInfo = useAppSelector<PersonalInfoState>(
    (state) => state.personalInfo
  );
  const navigate = useNavigate();
  const bg = useColorModeValue("white", "gray.700", "gray.800");
  const fg = useColorModeValue("gray.800", "gray.100", "gray.100");
  const border = useColorModeValue("gray.400", "gray.500", "gray.600");
  
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
    document.title = "Personal Info | Spectator";
  }, []);

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
          {/* <Button onClick={() => changeLanguage("en")}>en</Button>
          <Button onClick={() => changeLanguage("id")}>id</Button> */}
          {/* `eslint` is not happy with `!!foo`, need to use `Boolean` instead */}
          <FormControl
            id="email"
            mt="6"
            isInvalid={errors.studentId !== undefined}
          >
            <FormLabel>{t("ns1.translations.personal_info.student_number")}</FormLabel>
            <Input type="text" {...register("studentId")} autoComplete="off" borderColor={border} />
            <FormErrorMessage>{errors?.studentId?.message}!</FormErrorMessage>
          </FormControl>

          <FormControl
            mt="6"
            isInvalid={errors.programmingExp !== undefined}
          >
            <FormLabel>
              {t("ns1.translations.personal_info.programming_years")}
            </FormLabel>
            <Input
              borderColor={border}
              type="number"
              {...register("programmingExp")}
              autoComplete="off"
            />
            <FormErrorMessage>
              {errors?.programmingExp?.message}!
            </FormErrorMessage>
          </FormControl>

          <FormControl
          
            mt="6"
            isInvalid={errors.programmingExercise !== undefined}
          >
            <FormLabel>
              {t("ns1.translations.personal_info.programming_practice")}
            </FormLabel>
            <Input
              borderColor={border}
              type="number"
              {...register("programmingExercise")}
              autoComplete="off"
            />
            <FormErrorMessage>
              {errors?.programmingExercise?.message}!
            </FormErrorMessage>
          </FormControl>

          <FormControl
          
            mt="6"
            isInvalid={errors.programmingLanguage !== undefined}
          >
            <FormLabel>
              {t("ns1.translations.personal_info.programming_experience")}
            </FormLabel>
            <Input
              borderColor={border}
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
          {t("ns1.translations.ui.continue")}
        </Button>
      </Box>
    </Layout>
  );
}

export default withPublic(PersonalInfo);
