/// <reference types="vitest" />

import { defineConfig } from "vite";
import { resolve } from "path";

export default defineConfig({
  test: {
    global: true,
    environment: "jsdom",
    api: true
  },
  resolve: {
    alias: {
      "@": resolve("src")
    }
  }
});
