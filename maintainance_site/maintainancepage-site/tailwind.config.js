/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ['./src/**/*.{js,ts,jsx,tsx}', './index.html'],
  theme: {
    extend: {
      // Custom animations
      animation: {
        'spin-slow': 'spin 3s linear infinite', // Slower spin
        'spin-fast': 'spin 0.5s linear infinite', // Faster spin
      },
      // Custom keyframes (optional if you'd like more custom animations)
      keyframes: {
        spin: {
          '0%': { transform: 'rotate(0deg)' },
          '100%': { transform: 'rotate(360deg)' },
        },
      },
    },
  },
  plugins: [],
};