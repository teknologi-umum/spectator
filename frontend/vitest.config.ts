/// <reference types="vitest" />

import { defineConfig } from "vite";
import { resolve } from "path";

export default defineConfig({
  test: {
    global: true,
    environment: "jsdom"
  },
  resolve: {
    alias: {
      "@": resolve("src")
    }
  }
});
