import type { StepType } from "@reactour/tour";
import { TFunction } from "react-i18next";

export type TourStepsBuilder = (t: TFunction) => StepType[];
