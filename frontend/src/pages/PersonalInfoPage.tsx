import React, { useEffect } from "react";
import { useForm } from "react-hook-form";
import type { SubmitHandler } from "react-hook-form";
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
  Flex
} from "@chakra-ui/react";
import Layout from "@/components/Layout";
import { LocaleButton, ThemeButton } from "@/components/CodingTest";
import { useColorModeValue, useSignalR } from "@/hooks";
import { useTranslation } from "react-i18next";
import type { PersonalInfo } from "@/models/PersonalInfo";
import { Locale as DtoLocale } from "@/stub/enums";
import { setAccessToken } from "../store/slices/sessionSlice";
import { personalInfoTour } from "@/tours";
import { useTour } from "@reactour/tour";
import WithTour from "@/hoc/WithTour";

function PersonalInfoPage() {
  const { t } = useTranslation();
  const dispatch = useAppDispatch();
  const { locale } = useAppSelector((state) => state.locale);
  const { accessToken } = useAppSelector((state) => state.session);
  const personalInfo = useAppSelector((state) => state.personalInfo);
  const navigate = useNavigate();
  const bg = useColorModeValue("white", "gray.700", "gray.800");
  const fg = useColorModeValue("gray.800", "gray.100", "gray.100");
  const { sessionSpoke } = useSignalR();
  const { setIsOpen} = useTour();

  const {
    register,
    handleSubmit,
    formState: { errors }
  } = useForm({
    defaultValues: personalInfo,
    resolver: yupResolver(PersonalInfoSchema),
    reValidateMode: "onBlur"
  });

  const onSubmit: SubmitHandler<PersonalInfo> = async (data) => {
    if (accessToken === null) {
      console.error("accessToken is null");
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
      console.error(err);
    }
  };

  useEffect(() => {
    document.title = "Personal Info | Spectator";
    setIsOpen(true);
  }, []);

  useEffect(() => {
    // eslint-disable-next-line
    console.log("errors", errors);
  }, [errors]);

  useEffect(() => {
    if (accessToken === null) {
      let dtoLocale: DtoLocale;
      switch (locale) {
      case "EN":
        dtoLocale = DtoLocale.EN;
        break;
      case "ID":
        dtoLocale = DtoLocale.ID;
        break;
      default:
        console.error(`Unknown locale: ${locale}`);
        return;
      }

      sessionSpoke
        .startSession({ locale: dtoLocale })
        .then((sessionReply) => {
          sessionSpoke.setAccessToken(sessionReply.accessToken);
          setAccessToken(sessionReply.accessToken);
        })
        .catch((err) => {
          console.error(`Unable to start session. ${err}`);
        });
    }
  }, []);

  return (
    <Layout>
      <Flex gap={2} position="fixed" left={4} top={4} data-tour="step-1">
        <ThemeButton bg={bg} fg={fg} title={t("translation.translations.ui.theme")} />
        <LocaleButton bg={bg} fg={fg} />
      </Flex>
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
          <FormControl
            id="email"
            mt="6"
            isInvalid={errors.studentNumber !== undefined}
          >
            <FormLabel>
              {t("translation.translations.personal_info.student_number")}
            </FormLabel>
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
            <FormLabel>
              {t("translation.translations.personal_info.programming_years")}
            </FormLabel>
            <Input
              type="number"
              {...register("yearsOfExperience")}
              autoComplete="off"
            />
            <FormErrorMessage>
              {errors?.yearsOfExperience?.message}!
            </FormErrorMessage>
          </FormControl>

          <FormControl mt="6" isInvalid={errors.hoursOfPractice !== undefined}>
            <FormLabel>
              {t("translation.translations.personal_info.programming_practice")}
            </FormLabel>
            <Input
              type="number"
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
            <FormLabel>
              {t(
                "translation.translations.personal_info.programming_experience"
              )}
            </FormLabel>
            <Input
              type="text"
              {...register("familiarLanguages")}
              autoComplete="off"
            />
            <FormErrorMessage>
              {errors?.familiarLanguages?.message}!
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
        >
          {t("translation.translations.ui.continue")}
        </Button>
      </Box>
    </Layout>
  );
}

export default WithTour(PersonalInfoPage, personalInfoTour);
