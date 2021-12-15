import { useForm, SubmitHandler } from "react-hook-form";
import { yupResolver } from "@hookform/resolvers/yup";
import { PersonalInfoSchema } from "../../schema";
import { useAppDispatch, useAppSelector } from "../../store";
import { recordPersonalInfo } from "../../store/slices/personalInfoSlice";
import { useNavigate } from "react-router-dom";

interface FormValues {
  stdNo: string;
  programmingExp: number;
  programmingExercise: number;
  programmingLanguage: string;
}

const PersonalInfo2 = () => {
  const dispatch = useAppDispatch();
  const personalInfo = useAppSelector((state) => state.personalInfo);
  const navigate = useNavigate();

  const {
    register,
    handleSubmit,
    formState: { errors }
  } = useForm({
    defaultValues: personalInfo,
    resolver: yupResolver(PersonalInfoSchema)
  });

  const onSubmit: SubmitHandler<FormValues> = (data) => {
    dispatch(recordPersonalInfo(data));
    navigate("/instructions");
  };

  console.log("errors", errors);

  return (
    <div className="container">
      <form onSubmit={handleSubmit(onSubmit)}>
        <h5 className="grey-text text-darken-3">Personal Basic Info</h5>
        <div>
          <label style={{ fontSize: "16px", color: "black" }}>
            Student Number
          </label>
          <input type="text" {...register("stdNo")} />
        </div>

        <div style={{ marginTop: "15px" }}>
          <label style={{ fontSize: "16px", color: "black" }}>
            How many years have you been doing programming
          </label>
          <input type="number" {...register("programmingExp")} />
        </div>

        <div style={{ marginTop: "15px" }}>
          <label style={{ fontSize: "16px", color: "black" }}>
            How many hours in a week do you practice programming
          </label>
          <input type="number" {...register("programmingExercise")} />
        </div>

        <div style={{ marginTop: "15px" }}>
          <label style={{ fontSize: "16px", color: "black" }}>
            What programming languages are you familiar with (example : java,
            python, c)
          </label>
          <input type="text" {...register("programmingLanguage")} />
        </div>

        <div className="input-field">
          <button className="btn btn-primary lighten-1 z-depth-0">
            Submit
          </button>
        </div>
      </form>
    </div>
  );
};

export default PersonalInfo2;
