import path from 'path';
import HtmlWebpackPlugin from 'html-webpack-plugin';
import Dotenv from 'dotenv-webpack';

module.exports = {
  target: 'web',
  entry: './src/index',
  output: {
    path: path.join(__dirname, '/dist'),
    filename: 'main.js',
  },
  resolve: {
    extensions: ['.ts', '.tsx', '.js'],
  },
  module: {
    rules: [
      {
        test: /\.(ts|js)x?$/,
        exclude: /node_modules/,
        use: {
          loader: 'babel-loader',
        },
      },
    ],
  },
  plugins: [
    new HtmlWebpackPlugin({
      template: './src/index.html',
    }),
    new Dotenv(),
  ],
  stats: {
    errors: true,
    errorDetails: true,
  },
  devServer: {
    contentBase: path.join(__dirname, 'dist'),
    port: 8080,
    open: 'Google Chrome',
  },
};
