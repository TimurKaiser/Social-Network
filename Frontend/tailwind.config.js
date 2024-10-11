/** @type {import('tailwindcss').Config} */
export default {
  content: ["./src/**/*.{html,js,svelte,ts,svg}"],
  theme: {
    extend: {
      scale: {
        103: "1.03",
      },
      colors: {
        customGray: "#D9D9D9",
        customPP: "#6C757D",
      },
    },
  },
  plugins: [],
};
