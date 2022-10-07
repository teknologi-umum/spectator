import * as yup from "yup";

const transformRequired = (_?: string, origin?: string) => {
  if (origin === undefined || origin === "" || Number.isNaN(origin)) {
    return undefined;
  }
  return Number(origin);
};

// Labels might come from i18n translations
export const PersonalInfoSchema = yup.object().shape({
  email: yup.string().label("Email").email().ensure().required(),
  gender: yup.string().label("Gender").oneOf(["M", "F"]).required(),
  age: yup.number().label("Age").transform(transformRequired).min(1).required(),
  nationality: yup.string().label("Nationality").oneOf(["indonesia", "malaysia", "other"]).ensure().required(),
  studentNumber: yup
    .string()
    .label("Student Number")
    .matches(
      // based on these conditions:
      // 1. woa180020
      //    the first 3 characters must be alphabet, the rest is number
      // 2. 1202213133
      //    the first 2 character must be the number 12, the rest is number
      // 3. 17188939/1
      //    the first 8 characters must be number, the last 2 character is `/` and a number
      // 4. S2018499/1
      //    the first 1 character must be alphabet, the last 2 character is `/` and a number
      /([a-z]{3}\d{6}|\d{10}|\d{8}\/\d|S20\d{5}\/\d)/,
      "Invalid student number format"
    )
    .ensure()
    .required(),
  yearsOfExperience: yup
    .number()
    .label("Years of Experience")
    .transform(transformRequired)
    .min(0)
    .required(),
  hoursOfPractice: yup
    .number()
    .label("Hours of Practice")
    .transform(transformRequired)
    .min(0)
    .required(),
  familiarLanguages: yup
    .string()
    .label("Familiar Languages")
    .ensure()
    .required(),
  walletNumber: yup
    .string()
    .label("GrabPay Wallet / GoPay Number")
    .ensure()
    .required(),
  walletType: yup
    .string()
    .label("Wallet Type")
    .oneOf(["grabpay", "gopay"])
    .required()
});
