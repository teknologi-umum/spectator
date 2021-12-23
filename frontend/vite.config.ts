import { defineConfig } from "vite";
import react from "@vitejs/plugin-react";
import { resolve } from "path";
import SvgPlugin from 'vite-plugin-react-svg'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    react(),
    SvgPlugin(),
  ],
  resolve: {
    alias: {
      "@": resolve("src")
    }
  }
});
