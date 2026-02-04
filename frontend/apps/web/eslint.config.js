import js from "@eslint/js";
import globals from "globals";
import reactHooks from "eslint-plugin-react-hooks";
import reactRefresh from "eslint-plugin-react-refresh";
import react from "eslint-plugin-react";
import tseslint from "typescript-eslint";

export default tseslint.config(
  // Global ignores
  {
    ignores: ["dist/**", "node_modules/**", ".vite/**", "coverage/**", "*.config.js"],
  },
  
  // Base configs
  js.configs.recommended,
  ...tseslint.configs.recommended,
  
  // React and TypeScript specific config
  {
    files: ["**/*.{ts,tsx}"],
    languageOptions: {
      ecmaVersion: 2020,
      globals: globals.browser,
      parserOptions: {
        ecmaFeatures: {
          jsx: true,
        },
      },
    },
    plugins: {
      react,
      "react-hooks": reactHooks,
      "react-refresh": reactRefresh,
    },
    rules: {
      ...reactHooks.configs.recommended.rules,
      "react-refresh/only-export-components": [
        "warn",
        { allowConstantExport: true },
      ],
      // React specific rules
      "react/jsx-uses-react": "off", // Not needed with React 17+ automatic JSX runtime
      "react/react-in-jsx-scope": "off", // Not needed with Vite's automatic JSX runtime
    },
    settings: {
      react: {
        version: "detect", // Automatically detect React version
      },
    },
  }
);