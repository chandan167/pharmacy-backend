// @ts-check

import eslint from '@eslint/js';
import tseslint from 'typescript-eslint';
import stylistic from '@stylistic/eslint-plugin'
import eslintConfigPrettier from "eslint-config-prettier";

export default tseslint.config({
    languageOptions: {
        parserOptions: {
            project: true,
            tsconfigRootDir: import.meta.dirname,
        },
    },
    files: ['**/*.ts'],
    extends: [
        eslint.configs.recommended,
        ...tseslint.configs.recommended,
        stylistic.configs.recommended,
        eslintConfigPrettier
    ],
    rules: {
        "no-console": "error",
        quotes: ['error', 'single', { allowTemplateLiterals: true }],
        '@stylistic/semi': ["error", "always"],
        '@stylistic/indent': ['error', 4],
        '@typescript-eslint/no-unused-vars': ["error", { "argsIgnorePattern": "^_" }]
    }
});