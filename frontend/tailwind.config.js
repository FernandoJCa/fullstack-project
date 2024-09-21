/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./index.html",
    "./src/**/*.{vue,js,ts,jsx,tsx}",
  ],
  darkMode: 'class',
  theme: {
    extend: {},
  },
  plugins: [],
  safelist: [
    {
      pattern: /bg-(red|green|blue|gray)-(100|500|600|700)/,
    },
    {
      pattern: /text-(red|green|gray)-(100|700)/,
    },
  ],
}