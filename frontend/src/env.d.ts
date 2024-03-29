/// <reference types="vite/client" />

interface ImportMetaEnv {
  readonly VITE_SESSION_HUB_URL: string
  readonly VITE_EVENT_HUB_URL: string
  readonly VITE_LOGGER_URL: string
  readonly VITE_ADMIN_URL: string
  readonly VITE_VIDEO_STREAM_URL: string;
  readonly VITE_MINIO_URL: string;
}

interface ImportMeta {
  readonly env: ImportMetaEnv
}
