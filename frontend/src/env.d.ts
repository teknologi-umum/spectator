/// <reference types="vite/client" />

interface ImportMetaEnv {
  readonly VITE_SESSION_HUB_URL: string
  readonly VITE_EVENT_HUB_URL: string
  readonly VITE_LOGGER_URL: string
}

interface ImportMeta {
  readonly env: ImportMetaEnv
}
