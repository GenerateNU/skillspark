/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./app/**/*.{ts,tsx}", "./components/**/*.{ts,tsx}"],
  presets: [require("nativewind/preset")],
  theme: {
    extend: {
      fontFamily: {
        nunito: ["NunitoSans_400Regular"],
        "nunito-medium": ["NunitoSans_500Medium"],
        "nunito-semibold": ["NunitoSans_600SemiBold"],
        "nunito-bold": ["NunitoSans_700Bold"],
      },
    },
  },
  plugins: [],
};
