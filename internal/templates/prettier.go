package templates

// PrettierConfig returns the .prettierrc template.
func PrettierConfig() string {
	return `{
  "singleQuote": true,
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
