package templates

// EslintConfig returns the eslint.config.js template.
func EslintConfig() string {
	return `import js from '@eslint/js';
import parser from '@typescript-eslint/parser';
import reactDom from 'eslint-plugin-react-dom';
import reactX from 'eslint-plugin-react-x';

export default [
  {
    ignores: ['dist'],
  },
  js.configs.recommended,
  {
    files: ['**/*.{ts,tsx,js,jsx}'],
    languageOptions: {
      parser,
      parserOptions: {
        ecmaVersion: 'latest',
        sourceType: 'module',
        project: ['./tsconfig.json'],
      },
    },
    plugins: {
      'react-x': reactX,
      'react-dom': reactDom,
    },
    rules: {
      ...reactX.configs['recommended-typescript'].rules,
      ...reactDom.configs.recommended.rules,
    },
  },
];
`
}
