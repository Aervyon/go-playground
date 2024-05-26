/** @type {import('tailwindcss').Config} */
export default {
  daisyui: {
    themes: [
      {
        default: {
          "primary": "#372F42",
          "secondary": "#2B2533",
          "base-100": "#372F42",
          "neutral": "#FCF4C5",
          "accent": "#76EA7C",
          "success": "#76EA7C",
          "error": "#ED373A"
        }
      }
    ]
  },
  content: [
    './index.html',
    './src/**/*.{html,js,vue,ts}',
    "./src/**/*.{vue,js,ts,jsx,tsx}",
  ],
  theme: {
      colors: {
        bg: '#372F42',
        'bg-darker': '#2B2533',
        text: '#FCF4C5',
        'text-2': '#B2AD8D',
        accent: '#76EA7C',
        'accent-darker': '#448D48',
        'accent-2': '#ED373A',
        'accent-2-darker': '#6D181A',
        link: '#DADADA',
        placeholder: '#777777',
        input: '#2B2533'
      
    },
  },
  plugins: [
    require('daisyui')
  ],
}

