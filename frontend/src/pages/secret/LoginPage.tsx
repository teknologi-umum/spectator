import React from "react";
import { SubmitHandler, useForm } from "react-hook-form";
import { useColorModeValue } from "@/hooks";
import {
  Box,
  Button,
  Flex,
  FormControl,
  FormErrorMessage,
  Heading,
  Input
} from "@chakra-ui/react";
import { useEffect } from "react";
import { yupResolver } from "@hookform/resolvers/yup";
import { LoginSchema } from "@/schema/SecretPage";
import { useNavigate } from "react-router-dom";
import { useAppDispatch } from "@/store";
import { setSessionId } from "@/store/slices/sessionSlice";
import { loggerInstance } from "@/spoke/logger";
import { LogLevel } from "@microsoft/signalr";
import { ADMIN_BASE_URL } from "@/constants";
import { SettingsDropdown } from "@/components/Settings";

interface FormValues {
  password: string;
}

export default function Login() {
  const dispatch = useAppDispatch();
  const navigate = useNavigate();
  const boxBg = useColorModeValue("white", "gray.700", "gray.800");
  const bg = useColorModeValue("gray.100", "gray.800", "gray.900");
  const fg = useColorModeValue("gray.800", "gray.100", "gray.100");

  const {
    register,
    handleSubmit,
    formState: { errors, isSubmitting },
    setError
  } = useForm({
    defaultValues: { password: "" },
    resolver: yupResolver(LoginSchema),
    reValidateMode: "onBlur"
  });

  const onSubmit: SubmitHandler<FormValues> = async ({ password }) => {
    try {
      const response = await fetch(ADMIN_BASE_URL + "/login", {
        method: "POST",
        headers: {
          "Content-Type": "application/json"
        },
        body: JSON.stringify({ password })
      });
      const data = await response.json();

      if (response.status !== 200) {
        setError("password", {
          message: data.message
        });
        return;
      }

      dispatch(setSessionId(data.sessionId));
      navigate("/secret/download");
    } catch (err) {
      if (err instanceof Error) {
        loggerInstance.log(LogLevel.Error, err.message);
      }
    }
  };

  useEffect(() => {
    document.title = "Login | Spectator";
  }, []);

  useEffect(() => {
    if (
      errors === null ||
      errors === undefined ||
      Object.keys(errors).length === 0
    ) {
      return;
    }

    if (import.meta.env.DEV) {
      // eslint-disable-next-line no-console
      console.log("errors", errors);
    }
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
      <SettingsDropdown disableLocaleButton />
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
                <Button
                  isLoading={isSubmitting}
                  colorScheme="blue"
                  type="submit"
                >
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
