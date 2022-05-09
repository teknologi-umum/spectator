/// <reference types="vitest" />

import { defineConfig } from "vite";
import { resolve } from "path";

export default defineConfig({
  test: {
    environment: "jsdom"
  },
  build: {
    ssr: false
  },
  resolve: {
    alias: {
      "@": resolve("src")
    }
  }
});
