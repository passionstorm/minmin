const path = require('path');
const webpack = require('webpack');
const TerserPlugin = require('terser-webpack-plugin');
const VueLoaderPlugin = require('vue-loader').VueLoaderPlugin;
const HtmlWebpackPlugin = require('html-webpack-plugin');
const FriendlyErrorsPlugin = require('friendly-errors-webpack-plugin');
const MiniCssExtractPlugin = require('mini-css-extract-plugin');
const OptimizeCSSAssetsPlugin = require('optimize-css-assets-webpack-plugin');
const CompressionPlugin = require('compression-webpack-plugin');
const BrotliPlugin = require('brotli-webpack-plugin');
const ManifestPlugin = require('./src/plugins/manifest-plugin');

const mode = process.env.NODE_ENV;
const isDev = mode === 'development';

// process.env.ASSET_PATH = isDev ? '/' :'/'

const ASSET_PATH = process.env.ASSET_PATH || '/';
const resolve = dir => {
  return path.join(__dirname, dir);
};
const fs = require('fs');
const iconFiles = fs.readdirSync(resolve('src/assets/images/icons'));
const icons = iconFiles.map(e => e.replace(/\.[^/.]+$/, ''));

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
    new webpack.HashedModuleIdsPlugin(), // so that file hashes don't change unexpectedly
    new webpack.HotModuleReplacementPlugin(),
    new VueLoaderPlugin(),
    new ManifestPlugin({
      outputPath: resolve('/'),
    }),
    new HtmlWebpackPlugin({
      // filename: resolve( 'index.html'),
      template: resolve('index.html'),
      'hash': false,
      'compile': false,
      'favicon': false,
      'showErrors': true,
      'chunks': 'all',
      'minify': true,
      chunksSortMode: 'dependency',
    }),
    new MiniCssExtractPlugin({
      filename: '[name].css',
      chunkFilename: '[contenthash].css',
    }),
    new FriendlyErrorsPlugin(),
    new webpack.DefinePlugin({
      'ICON_LIST': JSON.stringify(icons),
      'BASE_URL': 'http://localhost:9000/api',
    }),
  ],
  devServer: {
    disableHostCheck: true,
    clientLogLevel: 'silent',
    historyApiFallback: true,
    hot: true,
    overlay: {warnings: true, errors: true},
    compress: true,
    port: 8080,
  },
  entry: {
    'admin': './src/main.js',
    // vendors: ['vue', 'vuex', 'core-js', 'vue-router', 'axios', 'lodash-es'],
  },
  output: {
    publicPath: ASSET_PATH,
    path: resolve('../public/dist/admin'),
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
  optimization: {
    moduleIds: 'hashed',
    runtimeChunk: 'single',
    splitChunks: {
      cacheGroups: {
        vendor: {
          test: /[\\/]node_modules[\\/]/,
          name: 'vendors',
          chunks: 'all',
        },
      },
    },
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
          {
            loader: MiniCssExtractPlugin.loader,
            options: {
              publicPath: path.resolve(__dirname, '../public/css'),
              // only enable hot in development
              hmr: isDev,
              reloadAll: true,
            },
          },
          'css-loader',
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
    ],
  },
};

if (process.env.NODE_ENV === 'production') {
  module.exports.plugins.push(
      new CompressionPlugin({
        filename: '[path].gz[query]',
        algorithm: 'gzip',
        test: /\.(js|css|html|svg)$/,
        threshold: 10240,
        minRatio: 0.8,
        deleteOriginalAssets: false,
      }),
      new BrotliPlugin({
        filename: '[path].br[query]',
        algorithm: 'brotli',
        quality: 9,
        test: /\.(js|css|html|svg)$/,
        threshold: 10240,
        minRatio: 0.8,
        deleteOriginalAssets: false,
      }),
  );
  module.exports.optimization = {
    minimize: true,
    moduleIds: 'hashed',
    runtimeChunk: 'single',
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
      chunks: 'all',
      maxInitialRequests: Infinity,
      minSize: 0,
      cacheGroups: {
        default: false,
        vendors: false,
        vendor: {
          name: 'vendor',
          chunks: 'initial',
          test: /[\\/]node_modules[\\/]/,
          priority: -10,
          minChunks: 2,
          reuseExistingChunk: true,
        },
        common: {
          name: 'common',
          minChunks: 2,
          chunks: 'initial',
          priority: 10,
          reuseExistingChunk: true,
        },
      },
    },
  };
}