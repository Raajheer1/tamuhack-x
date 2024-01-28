/** @type {import('tailwindcss').Config} */
export default {
  content: [
    "./index.html",
    "./src/**/*.{vue,js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {
      colors: {
        'aa-blue': '#0064d2',
        'aa-red' : '#CA1E23',
      },
    },
    fontFamily: {
      'outfit': ['Outfit', 'sans-serif'],
      'montserrat': ['Montserrat', 'sans-serif'],
    }
  },
  plugins: [],
}

