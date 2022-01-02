import { useAppSelector, useAppDispatch } from "@/store";
import { setColorMode as setColorModeState } from "@/store/slices/themeSlice";
import type { Theme } from "@/store/slices/themeSlice/types";

export function useColorMode() {
  const { currentTheme } = useAppSelector((state) => state.app);
  const dispatch = useAppDispatch();

  return { colorMode: currentTheme, setColorMode: (mode:Theme) => dispatch(setColorModeState(mode)) };
}
