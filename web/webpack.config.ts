/* eslint-disable @typescript-eslint/explicit-function-return-type */
/* eslint-disable @typescript-eslint/no-var-requires */
const webpack = require("webpack")
const path = require("path")
const Dotenv = require("dotenv-webpack")
const HtmlWebpackPlugin = require("html-webpack-plugin")
const CopyPlugin = require("copy-webpack-plugin")
const WorkboxPlugin = require("workbox-webpack-plugin")
const CompressionPlugin = require("compression-webpack-plugin")

module.exports = (env: unknown, options: {mode: string | undefined}) => {
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
      ],
    },
    devtool: isDevelopment ? "eval-cheap-module-source-map" : "source-map",
    optimization: {
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
      new CopyPlugin({
        patterns: [
          {
            from: "src/assets",
            to: "assets",
            globOptions: {
              ignore: ["**/original/**"],
            },
          },
          {from: "src/manifest.json", to: "."},
        ],
      }),
      new WorkboxPlugin.InjectManifest({
        swSrc: "./src/service-worker.ts",
        swDest: "sw.js",
        exclude: [/\.map$/, /manifest.json$/, /sw\.$/],
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
    experiments: {
      topLevelAwait: true,
    },
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
    },
    watchOptions: {
      aggregateTimeout: 300,
      ignored: /node_modules/,
      poll: true,
    },
  }
}
