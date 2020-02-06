const path = require('path')
const webpack = require('webpack')
const TerserPlugin = require('terser-webpack-plugin')
const VueLoaderPlugin = require('vue-loader').VueLoaderPlugin
const HtmlWebpackPlugin = require('html-webpack-plugin')
const OptimizeCSSAssetsPlugin = require('optimize-css-assets-webpack-plugin')
const VueSSRClientPlugin = require('vue-server-renderer/client-plugin')
const dotenv = require('dotenv')

const resolve = dir => {
  return path.join(__dirname, dir)
}
const env = dotenv.config({path: resolve('../.env')}).parsed
const mode = process.env.NODE_ENV
const isDev = mode === 'development'

module.exports = {
  mode: mode,
  devtool: 'cheap-eval-source-map',
  stats: {
    errors: true,
    errorDetails: true,
    hash: false,
    version: false,
    builtAt: false,
    chunkModules: false,
    children: false,
    chunks: false,
    entrypoints: false,
    modules: false,
    assets: false,
  },
  plugins: [
    new webpack.DefinePlugin({
      'API_URL': JSON.stringify(env.API_URL),
      'ADMIN_URL': JSON.stringify(env.ADMIN_URL),
      'STATIC_URL': JSON.stringify(env.STATIC_URL),
      // 'FIREBASE_API_KEY': JSON.stringify(''),
      // 'FIREBASE_AUTH_DOMAIN': JSON.stringify('mintoot-minh.firebaseapp.com'),
      'process.env.FIREBASE_DB_URL': JSON.stringify(
        'mintoot-minh.firebaseio.com'),
      // 'FIREBASE_PROJECT_ID': JSON.stringify(''),
      'process.env.FIREBASE_STORAGE_BUCKET': JSON.stringify(
        'mintoot-minh.appspot.com'),
      // 'FIREBASE_MESSAGE_ID': JSON.stringify(''),

    }),
    new webpack.HashedModuleIdsPlugin(), // so that file hashes don't change unexpectedly
    new webpack.HotModuleReplacementPlugin(),
    new VueLoaderPlugin(),
    // new ManifestPlugin({
    //   outputPath: resolve('/'),
    // }),
    new HtmlWebpackPlugin({
      // filename: resolve( 'index.html'),
      template: resolve('index.html'),
      'hash': false,
      'compile': false,
      'favicon': './src/modules/storm/img/favicon.png',
      'showErrors': true,
      'chunks': 'all',
      'minify': true,
      chunksSortMode: 'dependency',
    }),
    // new MiniCssExtractPlugin({
    //   filename: '[name].css',
    //   chunkFilename: '[contenthash].css',
    // }),
    // new FriendlyErrorsPlugin(),
  ],
  devServer: {
    disableHostCheck: true,
    clientLogLevel: 'silent',
    historyApiFallback: true,
    hot: true,
    overlay: {warnings: true, errors: true},
    // compress: true,
    port: 8082,
  },
  entry: {
    'main': './src/modules/storm/main.js',
    // vendors: ['vue', 'vuex', 'core-js', 'vue-router', 'axios'],
  },
  output: {
    publicPath: '/',
    path: resolve('../public/storm'),
    filename: '[name].[hash].js',
    chunkFilename: '[name].[contenthash].js',
  },
  resolve: {
    alias: {
      'vue$': isDev ? 'vue/dist/vue.esm.js' : 'vue/dist/vue.runtime.min.js',
      '_c': resolve('src/components'),
      '_w': resolve('src/widgets'),
      '_u': resolve('src/utils'),
      '_v': resolve('src/views'),
    },
    extensions: ['*', '.js', '.vue', '.json'],
  },
  watchOptions: {
    aggregateTimeout: 300,
    poll: 1000,
  },
  module: {
    rules: [
      {
        test: /\.js$/,
        exclude: /(node_modules)/,
        use: {
          loader: 'babel-loader',
          options: {
            plugins: [
              '@babel/plugin-syntax-dynamic-import',
            ],
          },
        },
      },
      {
        test: /\.css$/,
        use: [
          'vue-style-loader',
          {loader: 'css-loader'},
        ],
      },
      {
        test: /\.vue$/,
        loader: 'vue-loader',
        options: {
          loaders: {},
          hotReload: true,
        },
      },
      {
        test: /\.(svg)(\?.*)?$/,
        use: [
          {
            loader: 'svg-inline-loader',
            options: {
              limit: 10000,
              name: 'assets/images/icons/[name].[hash:7].[ext]',
            },
          },
        ],
      },
      {
        test: /\.(png|jpg|gif)$/,
        use: ['file-loader'],
      },
    ],
  },
}

if (process.env.NODE_ENV === 'production') {
  const BrotliPlugin = require('brotli-webpack-plugin')
  module.exports.plugins.push(
    // new CompressionPlugin({
    //   filename: '[path].gz[query]',
    //   algorithm: 'gzip',
    //   test: /\.(js|css|html|svg)$/,
    //   threshold: 10240,
    //   minRatio: 0.8,
    //   deleteOriginalAssets: false,
    // }),
    new VueSSRClientPlugin({}),
    new BrotliPlugin({
      filename: '[path].br[query]',
      algorithm: 'brotli',
      quality: 9,
      test: /\.(js|css|html|svg)$/,
      threshold: 10240,
      minRatio: 0.8,
      deleteOriginalAssets: false,
    }),
  )
  module.exports.optimization = {
    minimize: true,
    minimizer: [
      new TerserPlugin({
        terserOptions: {
          output: {
            comments: false,
          },
        },
        cache: true,
      }),
      new OptimizeCSSAssetsPlugin({}),
    ],
    splitChunks: {
      cacheGroups: {
        vendor: {
          test: /node_modules/,
          chunks: 'all',
          name: 'vendor',
          priority: 20,
          reuseExistingChunk: true,
        },
        common: {
          name: 'common',
          minChunks: 2,
          chunks: 'async',
          priority: 10,
          reuseExistingChunk: true,
          enforce: true,
        },
      },
    },
  }
}