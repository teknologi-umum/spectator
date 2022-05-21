import * as yup from "yup";

export const LoginSchema = yup.object().shape({
  password: yup.string().label("Password").ensure().required()
});
