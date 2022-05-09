/// <reference types="vitest" />

import { defineConfig } from "vite";
import { resolve } from "path";

export default defineConfig({
  test: {
    environment: "jsdom",
    transformMode: {
      web: [/\.([jt]sx?)$/]
    }
  },
  resolve: {
    alias: {
      "@": resolve("src")
    }
  }
});
