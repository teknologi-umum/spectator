import { defineConfig } from "vite";
import react from "@vitejs/plugin-react";
import { resolve } from "path";
import svgr from "@svgr/rollup";

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [react({ jsxRuntime: "classic" }), svgr({ memo: true })],
  resolve: {
    alias: {
      "@": resolve("src"),
    },
  },
  build: {
    ssr: false,
    target: ["es2020"],
    outDir: "../backend/Spectator/wwwroot",
  },
});
