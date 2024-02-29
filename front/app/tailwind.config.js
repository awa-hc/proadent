/** @type {import('tailwindcss').Config} */
module.exports = {
  darkMode: "class",
  content: ["./src/**/*.{html,ts}"],
  theme: {
    extend: {
      colors: {
        purple: "#AB1C49",
        red: "#FF3562",
        gray: "#1E1A1D",
      },
    },
  },
  plugins: [],
};
