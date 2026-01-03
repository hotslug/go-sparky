package templates

import (
	"strings"

	"github.com/hotslug/go-sparky/internal/plan"
)

// EslintConfig returns the eslint.config.js template.
func EslintConfig(p plan.Plan) string {
	ignores := []string{
		`"dist"`,
		`"node_modules"`,
		`"eslint.config.ts"`,
		`"eslint.config.js"`,
	}

	if p.IsVite() {
		ignores = append(ignores, `"vite.config.ts"`, `"vite.config.js"`)
	}

	if p.IsBun() {
		ignores = append(ignores, `"bun-env.d.ts"`)
	}

	ignoreBlock := strings.Join(ignores, ",\n      ")

	globalsExtra := ""
	if p.IsBun() {
		globalsExtra = "        ...globals.node,\n"
	}

	tailEntries := []string{
		"  configPrettier",
	}
	if p.IsBun() {
		tailEntries = append(tailEntries, `  {
    files: ["src/index.ts"],
    rules: {
      "@typescript-eslint/no-unused-vars": "off",
    },
  }`,
			`  {
    files: ["src/frontend.tsx"],
    rules: {
      "import/order": "off",
      "import/newline-after-import": "off",
    },
  }`)
	}

	tailBlock := strings.Join(tailEntries, ",\n")

	return `import js from "@eslint/js";
import globals from "globals";
import tsParser from "@typescript-eslint/parser";
import tsPlugin from "@typescript-eslint/eslint-plugin";
import tanstackQuery from "@tanstack/eslint-plugin-query";
import reactPlugin from "eslint-plugin-react";
import reactHooks from "eslint-plugin-react-hooks";
import jsxA11y from "eslint-plugin-jsx-a11y";
import importPlugin from "eslint-plugin-import";
import unicorn from "eslint-plugin-unicorn";
import prettier from "eslint-plugin-prettier";
import configPrettier from "eslint-config-prettier";

export default [
  {
    ignores: [
      ` + ignoreBlock + `,
    ],
  },
  js.configs.recommended,
  importPlugin.flatConfigs.recommended,
  {
    files: ["**/*.{ts,tsx}"],
    languageOptions: {
      parser: tsParser,
      parserOptions: {
        ecmaVersion: "latest",
        sourceType: "module",
        ecmaFeatures: {
          jsx: true,
        },
      },
      globals: {
        ...globals.browser,
        ...globals.es2020,
` + globalsExtra + `      },
    },
    plugins: {
      "@typescript-eslint": tsPlugin,
      "@tanstack/query": tanstackQuery,
      react: reactPlugin,
      "react-hooks": reactHooks,
      "jsx-a11y": jsxA11y,
      unicorn,
      prettier,
    },
    rules: {
      ...tsPlugin.configs.recommended.rules,
      ...reactPlugin.configs.recommended.rules,
      ...reactPlugin.configs["jsx-runtime"].rules,
      ...reactHooks.configs.recommended.rules,
      ...jsxA11y.configs.recommended.rules,
      ...tanstackQuery.configs.recommended.rules,
      ...unicorn.configs.recommended.rules,
      "react/react-in-jsx-scope": "off",
      "react/prop-types": "off",
      "react/no-unescaped-entities": "off",
      "import/no-unresolved": "off",
      "unicorn/filename-case": "off",
      "unicorn/prefer-node-protocol": "off",
      "import/order": [
        "warn",
        {
          "groups": [["builtin", "external"], "internal", ["parent", "sibling", "index"]],
          "newlines-between": "always",
          "alphabetize": { "order": "asc", "caseInsensitive": true },
        },
      ],
      "import/newline-after-import": ["warn", { "count": 1 }],
      "prettier/prettier": "warn",
    },
    settings: {
      react: {
        version: "detect",
      },
      "import/resolver": {
        typescript: {},
        node: {
          extensions: [".js", ".jsx", ".ts", ".tsx"],
        },
      },
    },
  },
` + tailBlock + `
];
`
}

// EslintConfigRelaxed returns a softer eslint.config.js template (no unicorn, import rules relaxed).
func EslintConfigRelaxed(p plan.Plan) string {
	ignores := []string{
		`"dist"`,
		`"node_modules"`,
		`"eslint.config.ts"`,
		`"eslint.config.js"`,
	}

	if p.IsVite() {
		ignores = append(ignores, `"vite.config.ts"`, `"vite.config.js"`)
	}

	if p.IsBun() {
		ignores = append(ignores, `"bun-env.d.ts"`)
	}

	ignoreBlock := strings.Join(ignores, ",\n      ")

	globalsExtra := ""
	if p.IsBun() {
		globalsExtra = "        ...globals.node,\n"
	}

	tailEntries := []string{
		"  configPrettier",
	}
	if p.IsBun() {
		tailEntries = append(tailEntries, `  {
    files: ["src/index.ts"],
    rules: {
      "@typescript-eslint/no-unused-vars": "off",
    },
  }`,
			`  {
    files: ["src/frontend.tsx"],
    rules: {
      "import/order": "off",
      "import/newline-after-import": "off",
    },
  }`)
	}

	tailBlock := strings.Join(tailEntries, ",\n")

	return `import js from "@eslint/js";
import globals from "globals";
import tsParser from "@typescript-eslint/parser";
import tsPlugin from "@typescript-eslint/eslint-plugin";
import tanstackQuery from "@tanstack/eslint-plugin-query";
import reactPlugin from "eslint-plugin-react";
import reactHooks from "eslint-plugin-react-hooks";
import jsxA11y from "eslint-plugin-jsx-a11y";
import importPlugin from "eslint-plugin-import";
import configPrettier from "eslint-config-prettier";

export default [
  {
    ignores: [
      ` + ignoreBlock + `,
    ],
  },
  js.configs.recommended,
  importPlugin.flatConfigs.recommended,
  {
    files: ["**/*.{ts,tsx}"],
    languageOptions: {
      parser: tsParser,
      parserOptions: {
        ecmaVersion: "latest",
        sourceType: "module",
        ecmaFeatures: {
          jsx: true,
        },
      },
      globals: {
        ...globals.browser,
        ...globals.es2020,
` + globalsExtra + `      },
    },
    plugins: {
      "@typescript-eslint": tsPlugin,
      "@tanstack/query": tanstackQuery,
      react: reactPlugin,
      "react-hooks": reactHooks,
      "jsx-a11y": jsxA11y,
    },
    rules: {
      ...tsPlugin.configs.recommended.rules,
      ...reactPlugin.configs.recommended.rules,
      ...reactPlugin.configs["jsx-runtime"].rules,
      ...reactHooks.configs.recommended.rules,
      ...jsxA11y.configs.recommended.rules,
      ...tanstackQuery.configs.recommended.rules,
      "react/react-in-jsx-scope": "off",
      "react/prop-types": "off",
      "react/no-unescaped-entities": "off",
      "import/no-unresolved": "off",
      "import/order": "off",
      "import/newline-after-import": "off",
    },
    settings: {
      react: {
        version: "detect",
      },
      "import/resolver": {
        typescript: {},
        node: {
          extensions: [".js", ".jsx", ".ts", ".tsx"],
        },
      },
    },
  },
` + tailBlock + `
];
`
}
