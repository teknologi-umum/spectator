import React from "react";
import ThemeButton from "@/components/CodingTest/TopBar/ThemeButton";
import { SubmitHandler, useForm } from "react-hook-form";
import { useColorModeValue } from "@/hooks";
import {
  Box,
  Button,
  Flex,
  FormControl,
  FormErrorMessage,
  FormLabel,
  Heading,
  Input,
  InputGroup
} from "@chakra-ui/react";
import { useAppDispatch, useAppSelector } from "@/store";
import { useEffect } from "react";
import { yupResolver } from "@hookform/resolvers/yup";
import { LoginSchema } from "@/schema/SecretPage";
import { useNavigate } from "react-router-dom";
import { recordPassword } from "@/store/slices/loginSlice";
import { LocaleButton } from "@/components/CodingTest";
import { useTranslation } from "react-i18next";

interface FormValues {
  password: string;
}

export default function Login() {
  const dispatch = useAppDispatch();
  const login = useAppSelector((state) => state.login);

  const { t } = useTranslation();
  const navigate = useNavigate();
  const boxBg = useColorModeValue("white", "gray.700", "gray.800");
  const bg = useColorModeValue("gray.100", "gray.800", "gray.900");
  const fg = useColorModeValue("gray.800", "gray.100", "gray.100");

  const {
    register,
    handleSubmit,
    formState: { errors }
  } = useForm({
    defaultValues: login,
    resolver: yupResolver(LoginSchema),
    reValidateMode: "onBlur"
  });

  const onSubmit: SubmitHandler<FormValues> = ({ password }) => {
    // TODO(elianiva): properly authenticate user from backend
    if (password !== "something") {
      return;
    }

    dispatch(recordPassword(password));
    navigate("/secret/download");
  };

  useEffect(() => {
    document.title = "Login | Spectator";
  }, []);

  useEffect(() => {
    console.log("errors", errors);
  }, [errors]);

  return (
    <Flex
      bg={bg}
      alignItems="center"
      justifyContent="center"
      w="full"
      minH="full"
      py="10"
      px="4"
    >
      <Flex gap={2} position="fixed" left={4} top={4} data-tour="step-1">
        <ThemeButton
          bg={boxBg}
          fg={fg}
          title={t("translation.translations.ui.theme")}
        />
        <LocaleButton bg={boxBg} fg={fg} />
      </Flex>
      <Box
        height="full"
        display="flex"
        alignItems="center"
        justifyContent="center"
      >
        <Box
          as="form"
          onSubmit={handleSubmit(onSubmit)}
          boxShadow="xl"
          p="8"
          rounded="md"
          mx="auto"
          bg={boxBg}
          color={fg}
        >
          <Heading size="lg" textAlign="center" fontWeight="700">
            Login
          </Heading>
          <Box>
            <FormControl mt="6" isInvalid={errors.password !== undefined}>
              <Flex mt="6" gap="4">
                <Input
                  type="text"
                  autoComplete="off"
                  placeholder="Type your password..."
                  {...register("password")}
                />
                <Button colorScheme="blue" type="submit">
                  Login
                </Button>
              </Flex>
              <FormErrorMessage>{errors?.password?.message}</FormErrorMessage>
            </FormControl>
          </Box>
        </Box>
      </Box>
    </Flex>
  );
}
