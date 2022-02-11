import * as yup from "yup";

const transformRequired = (_?: string, origin?: string) => {
  if (origin === undefined || origin === "" || Number.isNaN(origin)) {
    return undefined;
  }
  return Number(origin);
};

// Labels might come from i18n translations
export const PersonalInfoSchema = yup.object().shape({
  studentNumber: yup.string().label("Student Number").ensure().required(),
  yearsOfExperience: yup
    .number()
    .label("Years of Experience")
    .transform(transformRequired)
    .required(),
  hoursOfPractice: yup
    .number()
    .label("Hours of Practice")
    .transform(transformRequired)
    .required(),
  familiarLanguages: yup
    .string()
    .label("Familiar Languages")
    .ensure()
    .required()
});
