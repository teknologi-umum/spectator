import React, { useEffect } from "react";
import { useForm } from "react-hook-form";
import type { SubmitHandler, SubmitErrorHandler } from "react-hook-form";
import { yupResolver } from "@hookform/resolvers/yup";
import { PersonalInfoSchema } from "@/schema";
import { useAppDispatch, useAppSelector } from "@/store";
import { setPersonalInfo } from "@/store/slices/personalInfoSlice";
import { useNavigate } from "react-router-dom";
import {
  Box,
  FormControl,
  FormLabel,
  Input,
  Heading,
  Button,
  FormErrorMessage,
  InputGroup,
  Select,
  InputRightElement,
  HStack
} from "@chakra-ui/react";
import Layout from "@/components/Layout";
import { useColorModeValue } from "@/hooks";
import { useTranslation } from "react-i18next";
import type { PersonalInfo } from "@/models/PersonalInfo";
import { personalInfoTour } from "@/tours";
import { useTour } from "@reactour/tour";
import WithTour from "@/hoc/WithTour";
import { sessionSpoke } from "@/spoke";
import { loggerInstance } from "@/spoke/logger";
import { LogLevel } from "@microsoft/signalr";
import { SettingsDropdown } from "@/components/Settings";

function PersonalInfoPage() {
  const { t } = useTranslation("translation", {
    keyPrefix: "translations"
  });
  const dispatch = useAppDispatch();
  const { accessToken, tourCompleted } = useAppSelector(
    (state) => state.session
  );
  const personalInfo = useAppSelector((state) => state.personalInfo);
  const navigate = useNavigate();
  const bg = useColorModeValue("white", "gray.700", "gray.800");
  const fg = useColorModeValue("gray.800", "gray.100", "gray.100");
  const darkerBg = useColorModeValue("gray.100", "gray.800", "gray.900");
  const { setIsOpen } = useTour();

  const {
    register,
    handleSubmit,
    formState: { errors, isSubmitting }
  } = useForm({
    defaultValues: personalInfo,
    resolver: yupResolver(PersonalInfoSchema),
    reValidateMode: "onBlur"
  });

  const onSubmit: SubmitHandler<PersonalInfo> = async (data) => {
    if (accessToken === null) {
      loggerInstance.log(
        LogLevel.Error,
        "Access token was empty in Personal Info Page. This should never happen"
      );
      return;
    }

    try {
      await sessionSpoke.submitPersonalInfo({
        ...data,
        accessToken
      });
      dispatch(setPersonalInfo(data));
      navigate("/instructions");
    } catch (err) {
      if (err instanceof Error) {
        loggerInstance.log(LogLevel.Error, err.message);
      } else {
        loggerInstance.log(
          LogLevel.Error,
          "Unkown error occured in Personal Info Page"
        );
      }
    }
  };

  const onError: SubmitErrorHandler<PersonalInfo> = (err) => {
    // only log errors on development
    // this will be noop in production
    if (import.meta.env.DEV) {
      // eslint-disable-next-line
      console.log("errors", err);
    }
  };

  useEffect(() => {
    document.title = "Personal Info | Spectator";
    if (tourCompleted.personalInfo) return;
    setIsOpen(true);
  }, []);

  return (
    <Layout display="flex">
      <SettingsDropdown data-tour="step-1" />
      <Box
        display="flex"
        alignItems="center"
        justifyContent="center"
        height="full"
      >
        <Box
          as="form"
          onSubmit={handleSubmit(onSubmit, onError)}
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
            <FormControl mt="6" isInvalid={errors.email !== undefined}>
              <FormLabel>{t("personal_info.email")}</FormLabel>
              <Input type="text" {...register("email")} autoComplete="off" />
              <FormErrorMessage>{errors?.email?.message}!</FormErrorMessage>
            </FormControl>

            <FormControl mt="6" isInvalid={errors.age !== undefined}>
              <FormLabel>{t("personal_info.age")}</FormLabel>
              <Input type="text" {...register("age")} autoComplete="off" />
              <FormErrorMessage>{errors?.age?.message}!</FormErrorMessage>
            </FormControl>

            <HStack mt="6">
              <FormControl isInvalid={errors.gender !== undefined}>
                <FormLabel>{t("personal_info.gender")}</FormLabel>

                <Select
                  {...register("gender")}
                  size="md"
                  variant="outline"
                  defaultValue="M"
                >
                  <option value="M">Male</option>
                  <option value="F">Female</option>
                </Select>
                <FormErrorMessage>{errors?.gender?.message}!</FormErrorMessage>
              </FormControl>
              <FormControl isInvalid={errors.nationality !== undefined}>
                <FormLabel>{t("personal_info.nationality")}</FormLabel>

                <Select
                  {...register("nationality")}
                  size="md"
                  variant="outline"
                  defaultValue="M"
                >
                  <option value="indonesia">Indonesia</option>
                  <option value="malaysia">Malaysia</option>
                  <option value="other">Other</option>
                </Select>
                <FormErrorMessage>{errors?.gender?.message}!</FormErrorMessage>
              </FormControl>
            </HStack>

            <FormControl mt="6" isInvalid={errors.studentNumber !== undefined}>
              <FormLabel>{t("personal_info.student_number")}</FormLabel>
              <Input
                type="text"
                {...register("studentNumber")}
                autoComplete="off"
              />
              <FormErrorMessage>
                {errors?.studentNumber?.message}!
              </FormErrorMessage>
            </FormControl>

            <FormControl
              mt="6"
              isInvalid={errors.yearsOfExperience !== undefined}
            >
              <FormLabel>{t("personal_info.programming_years")}</FormLabel>
              <Input
                type="text"
                {...register("yearsOfExperience")}
                autoComplete="off"
              />
              <FormErrorMessage>
              </FormErrorMessage>
            </FormControl>

            <FormControl
              mt="6"
              isInvalid={errors.hoursOfPractice !== undefined}
            >
              <FormLabel>{t("personal_info.programming_practice")}</FormLabel>
              <Input
                type="text"
                {...register("hoursOfPractice")}
                autoComplete="off"
              />
              <FormErrorMessage>
                {errors?.hoursOfPractice?.message}!
              </FormErrorMessage>
            </FormControl>

            <FormControl
              mt="6"
              isInvalid={errors.familiarLanguages !== undefined}
            >
              <FormLabel>{t("personal_info.programming_experience")}</FormLabel>
              <Input
                type="text"
                {...register("familiarLanguages")}
                autoComplete="off"
              />
              <FormErrorMessage>
                {errors?.familiarLanguages?.message}!
              </FormErrorMessage>
            </FormControl>

            <FormControl
              mt="6"
              isInvalid={errors.familiarLanguages !== undefined}
            >
              <FormLabel>{t("personal_info.wallet_number")}</FormLabel>
              <InputGroup>
                <Input
                  type="text"
                  {...register("walletNumber")}
                  autoComplete="off"
                  pr="28"
                />
                <InputRightElement width="28">
                  <Select
                    {...register("walletType")}
                    size="sm"
                    variant="filled"
                    mr="1"
                    bg={darkerBg}
                    defaultValue="grabpay"
                  >
                    <option value="grabpay">GrabPay</option>
                    <option value="gopay">GoPay</option>
                  </Select>
                </InputRightElement>
              </InputGroup>
              <FormErrorMessage>
                {errors?.walletNumber?.message}!
              </FormErrorMessage>
            </FormControl>
          </Box>

          <Button
            colorScheme="blue"
            mx="auto"
            mt="6"
            type="submit"
            display="block"
            data-tour="step-2"
            isLoading={isSubmitting}
          >
            {t("ui.continue")}
          </Button>
        </Box>
      </Box>
    </Layout>
  );
}

export default WithTour(PersonalInfoPage, personalInfoTour);
