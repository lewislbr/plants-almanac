/* eslint-disable @typescript-eslint/no-var-requires */
const path = require('path');
const NodemonPlugin = require('nodemon-webpack-plugin');
const Dotenv = require('dotenv-webpack');

module.exports = {
  target: 'node',
  entry: './src/index',
  output: {
    path: path.join(__dirname, '/dist'),
    filename: 'index.js',
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
  plugins: [new NodemonPlugin(), new Dotenv()],
  externals: [
    { express: 'commonjs express' },
    { mongoose: 'commonjs mongoose' },
    {
      bufferutil: 'commonjs bufferutil',
      'utf-8-validate': 'commonjs utf-8-validate',
    },
  ],
  stats: {
    assetsSort: '!size',
    builtAt: false,
    children: false,
    entrypoints: false,
    errors: true,
    errorDetails: true,
    hash: false,
    modules: false,
    timings: false,
  },
};
