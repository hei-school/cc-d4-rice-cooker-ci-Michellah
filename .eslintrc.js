module.exports = {
    parser: '@typescript-eslint/parser',
    plugins: ['@typescript-eslint'],
    extends: [
      'eslint:recommended',
      'plugin:@typescript-eslint/recommended'
    ],
    parserOptions: {
      ecmaVersion: 2021,
      sourceType: 'module',
    },
    rules: {
        '@typescript-eslint/explicit-module-boundary-types': ['error', {
            allowArgumentsExplicitlyTypedAsAny: true,
        }],

        '@typescript-eslint/typedef': ['error', {
            variableDeclaration: true,
            parameter: true,
            propertyDeclaration: true,
            arrowParameter: true,
            memberVariableDeclaration: true,
          }],

        '@typescript-eslint/consistent-type-definitions': ['error', 'interface'],

        '@typescript-eslint/prefer-as-const': 'error',

        '@typescript-eslint/no-explicit-any': 'error',

        '@typescript-eslint/naming-convention': [
            'error',
            {
              selector: 'variable',
              format: ['camelCase', 'UPPER_CASE'],
            },
            {
              selector: 'function',
              format: ['camelCase', 'PascalCase'],
            },
            {
              selector: 'class',
              format: ['PascalCase'],
            },
          ],

        '@typescript-eslint/indent': ['error', 2],
    },
  };