import { useAppSelector } from "@/store";
import type { InitialState as AppState } from "@/store/slices/appSlice/types";

export function useColorModeValue(light: string, dimmed: string, dark: string) {
  const { currentTheme } = useAppSelector<AppState>((state) => state.app);

  if (currentTheme === "light") return light;
  if (currentTheme === "dimmed") return dimmed;
  if (currentTheme === "dark") return dark;

  return "";
}