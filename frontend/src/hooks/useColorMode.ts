import { useAppSelector, useAppDispatch } from "@/store";
import { setColorMode as setColorModeState } from "@/store/slices/appSlice";
import type { Theme } from "@/store/slices/appSlice/types";

export function useColorMode() {
  const { colorMode } = useAppSelector((state) => state.app);
  const dispatch = useAppDispatch();

  return { colorMode, setColorMode: (mode:Theme) => dispatch(setColorModeState(mode)) };
}
