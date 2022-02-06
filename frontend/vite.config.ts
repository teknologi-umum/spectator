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
  },
  build: {
    outDir: "../backend/Spectator/wwwroot",
    rollupOptions: {
      output: {
        manualChunks: {
          vendor: ["react", "react-dom", "react-router-dom"],
          codemirror_cpp: ["@codemirror/lang-cpp"],
          codemirror_java: ["@codemirror/lang-java"],
          codemirror_javascript: ["@codemirror/lang-javascript"],
          codemirror_python: ["@codemirror/lang-python"],
          codemirror_php: ["@codemirror/lang-php"],
          codemirror: ["@codemirror/gutter", "@uiw/react-codemirror"],
          chakra: [
            "@chakra-ui/react",
            "@emotion/react",
            "@emotion/styled",
            "framer-motion"
          ]
        }
      }
    }
  }
});
