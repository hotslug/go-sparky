package templates

// PrettierConfig returns the .prettierrc template.
func PrettierConfig() string {
	return `{
  "printWidth": 120,
  "tabWidth": 2,
  "useTabs": false,
  "trailingComma": "es5",
  "bracketSpacing": true,
  "arrowParens": "avoid",
  "endOfLine": "lf",
  "semi": false,
  "plugins": [
    "prettier-plugin-tailwindcss",
    "@ianvs/prettier-plugin-sort-imports"
  ],
  "importOrder": [
    "^@/(.*)$",
    "^[./]"
  ],
  "importOrderSeparation": true
}
`
}
