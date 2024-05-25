/** @type {import('tailwindcss').Config} */
export default {
  content: [
    './index.html',
    './src/**/*.{html,js,vue,ts}',
    "./src/**/*.{vue,js,ts,jsx,tsx}",
    'node_modules/preline/dist/*.js'
  ],
  theme: {
    extend: {
      colors: {
        bg: '#372F42',
        'bg-darker': '#2B2533',
        text: '#FCF4C5',
        accent: '#76EA7C',
        'accent-darker': '#448D48',
        'accent-2': '#ED373A',
        'accent-2-darker': '#6D181A',
        link: '#DADADA',
        placeholder: '#777777',
        input: '#2B2533'
      }
      
    },
  },
  plugins: [
    require('preline/plugin')
  ],
}

