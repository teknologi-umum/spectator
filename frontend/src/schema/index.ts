import * as yup from "yup";

const transformRequired = (_?: string, origin?: string) => {
  if (!origin || Number.isNaN(origin)) {
    return undefined;
  }
  return +origin;
};

export const PersonalInfoSchema = yup.object().shape({
  stdNo: yup.string().ensure().required(),
  programmingExp: yup.number().transform(transformRequired).required(),
  programmingExercise: yup.number().transform(transformRequired).required(),
  programmingLanguage: yup.string().ensure().required(),
});
