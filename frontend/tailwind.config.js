/** @type {import('tailwindcss').Config} */
export default {
  darkMode: "class",
  content: ['./src/**/*.{html,js,svelte,ts}'],
  theme: {
    extend: {
      colors: {
       "lightest": "var(--color-lightest)",
        "lighter": "var(--color-lighter)",
        "light": "var(--color-light)",
        "strong": "var(--color-strong)",
        "stronger": "var(--color-stronger)",
        "strongest": "var(--color-strongest)",
      }
    },
  },
  plugins: [],
}

