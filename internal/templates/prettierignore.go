package templates

// PrettierIgnore returns the .prettierignore contents.
func PrettierIgnore() string {
	return `node_modules
dist
build
*.min.js
`
}

