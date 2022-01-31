import { useAppSelector } from "@/store";

export function useColorModeValue(light: string, dimmed: string, dark: string) {
  const { currentTheme } = useAppSelector((state) => state.app);

  if (currentTheme === "light") return light;
  if (currentTheme === "dimmed") return dimmed;
  if (currentTheme === "dark") return dark;

  return "";
}