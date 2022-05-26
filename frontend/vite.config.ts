import { defineConfig } from "vite";
import react from "@vitejs/plugin-react";
import { resolve } from "path";
import svgr from "@svgr/rollup";

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [
    react({
      jsxRuntime: "classic"
    }),
    svgr({
      memo: true
    })
  ],
  resolve: {
    alias: {
      "@": resolve("src")
    }
  },
  build: {
    ssr: false,
    target: ["es2020"],
    outDir: "../backend/Spectator/wwwroot",
    rollupOptions: {
      output: {
        manualChunks: {
          vendor: ["react", "react-dom", "react-router-dom"],
          codemirror: [
            "@codemirror/state",
            "@codemirror/view",
            "@uiw/react-codemirror",
            "@codemirror/lang-cpp",
            "@codemirror/lang-java",
            "@codemirror/lang-javascript",
            "@codemirror/lang-python",
            "@codemirror/lang-php"
          ],
          chakra: [
            "@chakra-ui/react",
            "@emotion/react",
            "@emotion/styled",
            "framer-motion"
          ],
          react_tour: ["@reactour/tour"]
        }
      }
    }
  }
});
