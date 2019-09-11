module.exports = {
  root: true,
  env: {
    node: true,
    es6: true
  },
  extends: [
    "plugin:vue/recommended",
    "eslint:recommended",
    "prettier/vue",
  ],
  rules: {
    "prettier/prettier": "error",
    "no-console": process.env.NODE_ENV === "production" ? "error" : "off",
    "no-debugger": process.env.NODE_ENV === "production" ? "error" : "off"
  },
  globals: {
    $nuxt: true
  },
  parserOptions: {
    parser: "babel-eslint",
    ecmaVersion: 2018,
    sourceType: "module"
  },
  plugins: ["prettier"],
};
/**
 * on vscode settings
 * {
 *  "eslint.validate": [
      {
        "language": "vue",
        "autoFix": true
      },
      {
        "language": "javascript",
        "autoFix": true
      },
      {
        "language": "javascriptreact",
        "autoFix": true
      }
    ],
    "eslint.autoFixOnSave": true,
    "editor.formatOnSave": false,
    "vetur.validation.template": false,
 * }
 */