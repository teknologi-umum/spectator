import { defineConfig } from "vite";
import react from "@vitejs/plugin-react";
import { resolve } from "path";
import svgr from "@svgr/rollup";

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    react(),
    svgr({
      memo: true
    })
  ],
  resolve: {
    alias: {
      "@": resolve("src")
    }
  }
});
