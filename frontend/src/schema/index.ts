import * as yup from "yup";

const transformRequired = (_?: string, origin?: string) => {
  if (origin === undefined || origin === "" || Number.isNaN(origin)) {
    return undefined;
  }
  return Number(origin);
};

// Labels might come from i18n translations
export const PersonalInfoSchema = yup.object().shape({
  stdNo: yup.string().label("Student Number").ensure().required(),
  programmingExp: yup
    .number()
    .label("Programming Experience")
    .transform(transformRequired)
    .required(),
  programmingExercise: yup
    .number()
    .label("Programming Exercise")
    .transform(transformRequired)
    .required(),
  programmingLanguage: yup
    .string()
    .label("Programming Language")
    .ensure()
    .required()
});
