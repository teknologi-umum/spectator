import * as yup from "yup";

// Labels might come from i18n translations
export const LoginSchema = yup.object().shape({
  password: yup.string().label("Password").ensure().required(),
})