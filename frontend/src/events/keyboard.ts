import { emit } from "@/events/emitter";

const UNRELATED_KEYS: Record<string, boolean> = {
  F1: true,
  F2: true,
  F3: true,
  F4: true,
  F5: true,
  F6: true,
  F7: true,
  F8: true,
  F9: true,
  F10: true,
  F11: true,
  F12: true,
  Meta: true,
  Alt: true,
  Ctrl: true
};

const SHORTCUT_KEYS: Record<string, boolean> = {
  // save dialog
  s: true,
  // print dialog
  p: true,
  // bookmark
  d: true,
  // close window
  w: true,
  // copy and paste
  c: true,
  v: true
};

export function keystrokeHandler(connection: unknown) {
  return async (e: KeyboardEvent) => {
    const isPressingModifier = e.ctrlKey || e.altKey || e.shiftKey || e.metaKey;

    const data = {
      event: "keyboard",
      value: JSON.stringify({
        key: isPressingModifier
          ? // TODO(elianiva): is there any better way to do this?
            [
              e.ctrlKey && "Ctrl",
              e.altKey && "Alt",
              e.shiftKey && "Shift",
              e.metaKey && "MetaKey",
              // exclude modifier because we already include them above
              !["Control", "Meta", "Shift", "Alt"].includes(e.key) && e.key
            ].filter(Boolean)
          : e.key,
        unrelated: UNRELATED_KEYS[e.key] || isPressingModifier
      }),
      timestamp: Date.now()
    };

    await emit(connection, data);

    if (UNRELATED_KEYS[e.key] || (e.ctrlKey && SHORTCUT_KEYS[e.key])) {
      e.preventDefault();
    }
  };
}
