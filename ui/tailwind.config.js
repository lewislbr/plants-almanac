module.exports = {
  purge: ['./src/**/*.html', './src/**/*.tsx'],
  theme: {
    extend: {
      colors: {
        backdrop: 'hsl(0, 0%, 0%, 0.75)',
      },
      spacing: {
        '5pc': '5%',
      },
    },
  },
  variants: {},
  plugins: [],
};
