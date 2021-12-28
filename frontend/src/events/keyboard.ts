import { emit } from "@/events/emitter";

const F_KEYS: Record<string, boolean> = {
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
  F12: true
};

export function keystrokeHandler(connection: unknown) {
  return async (e: KeyboardEvent) => {
    // based on:
    // https://github.com/teknologi-umum/spectator/blob/25879b9b599790d45dee83892974a28ff40abd9c/backend/Spectator.RepositoryDALs/Internals/EventMapper.cs#L78-L84
    const payload = {
      keyChar: e.key,
      key: e.keyCode,
      shift: e.shiftKey,
      alt: e.altKey,
      control: e.ctrlKey,
      meta: e.metaKey,
      unrelated: F_KEYS[e.key]
    };

    const data = { event: "keyboard", timestamp: Date.now() };

    // ignore if it's triggered from codemirror because we it has separate
    // listener
    if ((e.target as HTMLDivElement).classList[0] === "cm-content") {
      // everything INSIDE the editor is always related except F-keys
      payload.unrelated ||= false;

      // don't allow pressing F-keys inside the editor
      if (F_KEYS[e.key]) e.preventDefault();

      await emit(connection, { ...data, value: JSON.stringify(payload) });
      return;
    }

    // everything OUTSIDE the editor is always unrelated
    payload.unrelated = true;
    await emit(connection, { ...data, value: JSON.stringify(payload) });

    // don't allow to do anything outside of the code editor
    e.preventDefault();
  };
}
