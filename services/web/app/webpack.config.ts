/* eslint-disable @typescript-eslint/explicit-function-return-type */
/* eslint-disable @typescript-eslint/no-var-requires */
const webpack = require("webpack")
const path = require("path")
const Dotenv = require("dotenv-webpack")
const TerserPlugin = require("terser-webpack-plugin")
const HtmlWebpackPlugin = require("html-webpack-plugin")
const MiniCssExtractPlugin = require("mini-css-extract-plugin")
const CssMinimizerPlugin = require("css-minimizer-webpack-plugin")
const CompressionPlugin = require("compression-webpack-plugin")

module.exports = (env: any, options: {mode: string | undefined}) => {
  const isDevelopment = options.mode !== "production"

  process.env.NODE_ENV = options.mode

  return {
    mode: isDevelopment ? "development" : "production",
    target: "web",
    entry: "./src/index.tsx",
    output: {
      filename: isDevelopment ? "[name].js" : "[name].[contenthash:8].js",
      path: path.join(__dirname, "/dist"),
    },
    resolve: {
      extensions: [".ts", ".tsx", ".js"],
    },
    module: {
      rules: [
        {
          test: /\.ts(x)?$/,
          exclude: /node_modules/,
          use: {
            loader: "babel-loader",
            options: {
              cacheDirectory: true,
            },
          },
        },
        {
          test: /\.css$/,
          use: [
            {loader: MiniCssExtractPlugin.loader},
            {
              loader: "css-loader",
              options: {
                importLoaders: 1,
                modules: false,
                sourceMap: true,
              },
            },
            {
              loader: "postcss-loader",
              options: {
                postcssOptions: {
                  ident: "postcss",
                  plugins: [require("tailwindcss"), require("autoprefixer")],
                },
              },
            },
          ],
        },
        {
          test: /\.(woff|woff2|eot|ttf|otf)$/,
          use: ["file-loader"],
        },
        {
          test: /\.(png|jpe?g|gif|svg)$/,
          use: [
            {
              loader: "file-loader",
              options: {
                outputPath: "images",
              },
            },
            {
              loader: "image-webpack-loader",
              options: {
                mozjpeg: {
                  progressive: true,
                  quality: 65,
                },
                optipng: {
                  enabled: false,
                },
                pngquant: {
                  quality: "65-90",
                  speed: 4,
                },
                gifsicle: {
                  interlaced: false,
                },
                webp: {
                  quality: 75,
                },
              },
            },
          ],
        },
      ],
    },
    devtool: isDevelopment ? "eval-cheap-module-source-map" : "source-map",
    optimization: {
      minimizer: [new TerserPlugin(), new CssMinimizerPlugin()],
      runtimeChunk: {
        name: "runtime",
      },
      splitChunks: {
        chunks: "all",
        cacheGroups: {
          defaultVendors: {
            name: "vendors",
            test: /[\\/]node_modules[\\/]/,
          },
        },
        name: false,
      },
    },
    performance: {
      hints: false,
    },
    plugins: [
      new Dotenv({
        path: ".env",
      }),
      new HtmlWebpackPlugin({
        minify: isDevelopment
          ? false
          : {
              collapseWhitespace: true,
              keepClosingSlash: true,
              minifyCSS: true,
              minifyJS: true,
              minifyURLs: true,
              removeComments: true,
              removeEmptyAttributes: true,
              removeRedundantAttributes: true,
              removeScriptTypeAttributes: true,
              removeStyleLinkTypeAttributes: true,
              useShortDoctype: true,
            },
        template: "./src/index.html",
      }),
      new MiniCssExtractPlugin({
        filename: isDevelopment ? "[name].css" : "[name].[contenthash:8].css",
      }),
      ...(isDevelopment
        ? []
        : [
            new CompressionPlugin({
              algorithm: "brotliCompress",
              compressionOptions: {level: 11},
              filename: "[name][ext].br",
              minRatio: Number.MAX_SAFE_INTEGER,
              test: /\.(html|css|js|svg)$/,
              threshold: 0,
            }),
          ]),
      ...(isDevelopment ? [new webpack.HotModuleReplacementPlugin()] : []),
    ],
    stats: {
      assetsSort: "!size",
      colors: true,
      entrypoints: false,
      errors: true,
      errorDetails: true,
      groupAssetsByChunk: false,
      groupAssetsByExtension: false,
      groupAssetsByInfo: false,
      groupAssetsByPath: false,
      modules: false,
      relatedAssets: true,
      timings: false,
      version: false,
    },
    devServer: {
      contentBase: "dist",
      historyApiFallback: true,
      host: "0.0.0.0",
      hot: true,
      port: process.env.WEB_APP_PORT,
    },
    watchOptions: {
      aggregateTimeout: 300,
      ignored: /node_modules/,
      poll: true,
    },
  }
}
