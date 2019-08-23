const path = require('path');

module.exports = {
  target: 'node',
  entry: './src/app',
  output: {
    path: path.join(__dirname, '/dist'),
    filename: 'app.js',
  },
  resolve: {
    extensions: ['.ts', '.mjs', '.js'],
  },
  module: {
    rules: [
      {
        test: /\.(ts|js)$/,
        exclude: /node_modules/,
        use: {
          loader: 'babel-loader',
        },
      },
    ],
  },
  externals: [''],
  stats: {
    errors: true,
    errorDetails: true,
  },
};
