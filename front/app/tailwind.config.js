/** @type {import('tailwindcss').Config} */
module.exports = {
  darkMode: "class",
  content: ["./src/**/*.{html,ts}"],
  theme: {
    extend: {
      colors: {
        purple: "#7B4B94",
        red: "#FF3562",
        gray: "#1E1A1D",
      },
    },
  },
  plugins: [],
};
