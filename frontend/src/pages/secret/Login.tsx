import Layout from "@/components/Layout";
import ThemeButton from "@/components/ThemeButton";
// import type { SubmitHandler } from "react-hook-form";
import { SubmitHandler, useForm } from "react-hook-form";
import type { InitialState as LoginState } from "@/store/slices/loginSlice/types";
import { useColorModeValue } from "@/hooks";
import { 
  Box, 
  Button, 
  FormControl, 
  FormErrorMessage, 
  FormLabel, 
  Heading, 
  Input 
} from "@chakra-ui/react";
import { useAppDispatch, useAppSelector } from "@/store";
import { useEffect } from "react";
import { yupResolver } from "@hookform/resolvers/yup";
import { LoginSchema } from "@/schema/Login";
import { useNavigate } from "react-router-dom";
import { recordPassword } from "@/store/slices/loginSlice";

interface FormValues {
  password: string;
}

export default function Login() {
  const dispatch = useAppDispatch();
  const login = useAppSelector<LoginState>(
    (state) => state.login
  )
  
  const navigate = useNavigate();
  const bg = useColorModeValue("white", "gray.700", "gray.800");
  const fg = useColorModeValue("gray.800", "gray.100", "gray.100");
  const border = useColorModeValue("gray.400", "gray.500", "gray.600");

  const {
    register,
    handleSubmit,
    formState: { errors }
  } = useForm({
    defaultValues: login,
    resolver: yupResolver(LoginSchema),
    reValidateMode: "onBlur"
  })

  const onSubmit: SubmitHandler<FormValues> = (data) => {
    dispatch(recordPassword(data.password));
    navigate("/download");
  }

  useEffect(() => {
    document.title = "Login | Spectator";
  }, [])
  
  useEffect(() => {
    console.log("errors", errors)
  }, [errors])

  return (
    <Layout>
      <ThemeButton position="fixed" />
      <Box
        height="90vh"
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
          maxW="container.sm"
          mx="auto"
          bg={bg}
          color={fg}
        >
          <Heading size="lg" textAlign="center" fontWeight="700">
            Login
          </Heading>

          <Box>
            {/* `eslint` is not happy with `!!foo`, need to use `Boolean` instead */}
            <FormControl
              mt="6"
              // isInvalid={errors.programmingExp !== undefined}
              isInvalid={errors.password !== undefined}
            >
              <FormLabel>
                Password
              </FormLabel>
              <Input
                borderColor={border}
                type="text"
                autoComplete="off"
                {...register("password")}
              />
              <FormErrorMessage>
                Password harus diisi
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
            {/* {t("translation.translations.ui.continue")} */}
            Submit
          </Button>
        </Box>
      </Box>
    </Layout>
  )
}