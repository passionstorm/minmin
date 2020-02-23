process.env.VUE_APP_VERSION = require('./package.json').version

const path = require('path')
const fs = require('fs')
const isProduction = process.env.NODE_ENV === 'production'
const isOpenPrerenderSPA = false
const SizePlugin = require('size-plugin')

function resolveRealPath(dir) {
  return path.join(__dirname, dir)
}

const alias = {
  'vue$': 'vue/dist/vue.esm.js',
  '@helper': resolveRealPath('src/helper'),
  '@page': resolveRealPath('src/pages'),
  '@asset': resolveRealPath('src/assets'),
  '@widget': resolveRealPath('src/widgets'),
  '@router': resolveRealPath('src/router'),
  '@view': resolveRealPath('src/views'),
  '@plugin': resolveRealPath('src/plugins'),
}

function loadGlobalStyles() {
  const variables = fs.readFileSync('./src/assets/scss/_variables.scss', 'utf-8')
  const button = fs.readFileSync('./src/assets/scss/_main_sidebar.scss', 'utf-8')
  return variables + button
}

const analyzer = (conf) => {
  if (process.env.npm_config_report) {
    conf.plugin('webpack-bundle-analyzer').use(require('webpack-bundle-analyzer').BundleAnalyzerPlugin)
  }
}

module.exports = {
  publicPath: '/',
  filenameHashing: false,
  outputDir: 'dist',
  productionSourceMap: !isProduction,
  pages: {
    index:{
      entry: './src/main.js',
      template: './public/index.html',
      filename: 'index.html',
    },
    homepage: {
      entry: './src/pages/home/main.js',
      template: './public/index.html',
      // output as dist/index.html
      filename: 'homepage.html',
      // chunks: ['chunk-vendors', 'chunk-common', 'index'],
    },
    mogodb:{
      entry: './src/pages/mgodb/main.js',
      template: './public/index.html',
      // output as dist/index.html
      filename: 'mgopage.html',
    }
  },
  css: {
    loaderOptions: {
      scss: {
        prependData: `@import "./assets/scss/app.scss";`
      }
    }
  },
  chainWebpack: conf => {
    for (const [key, value] of Object.entries(alias)) {
      conf.resolve.alias.set(key, value)
    }
    conf.module.rules.delete('svg')
    conf.module.rule('svg').test(/\.svg$/).use('svg-sprite-loader').loader('svg-sprite-loader').options({
      name: '[name]-[hash:7]',
      prefixize: true,
    })
    // https://github.com/webpack-contrib/webpack-bundle-analyzer
    analyzer()
  },
  configureWebpack: {
    plugins: [
      isProduction && isOpenPrerenderSPA ? new require('prerender-spa-plugin')({
            // Required - The path to the webpack-outputted app to prerender.
            staticDir: path.join(__dirname, 'dist'),
            // Required - Routes to render.
            routes: ['/', '/explore'],
          })
          : () => {
          },
      isProduction ? new SizePlugin() : () => {
      },
    ],
  },
  // configure webpack-dev-server behavior
  devServer: {
    open: process.platform === 'darwin',
    host: '0.0.0.0',
    port: 8080,
    https: false,
    hotOnly: false,
    // See https://github.com/vuejs/vue-cli/blob/dev/docs/cli-service.md#configuring-proxy
    proxy: null, // string | Object
    before: app => {}
  },
}