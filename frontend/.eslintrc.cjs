module.exports = {
  env: {
    browser: true,
    es2021: true
  },
  extends: [
    "plugin:react/recommended",
    "@teknologi-umum/eslint-config-typescript"
  ],
  parser: "@typescript-eslint/parser",
  parserOptions: {
    ecmaFeatures: {
      jsx: true
    },
    ecmaVersion: 12,
    sourceType: "module"
  },
  plugins: ["react", "@typescript-eslint"],
  settings: {
    react: {
      version: "17.0"
    }
  },
  rules: {
    "no-extra-parens": "off"
  }
};
