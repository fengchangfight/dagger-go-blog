const webpack = require('webpack')
module.exports = {
  /*
  ** Headers of the page
  */
  head: {
    title: '???的博客',
    meta: [
      { charset: 'utf-8' },
      { name: 'viewport', content: 'width=device-width, initial-scale=1' },
      { hid: 'description', name: 'description', content: '???' }
    ],
    script: [
      { src: 'https://cdnjs.cloudflare.com/ajax/libs/jquery/3.1.1/jquery.min.js' }
    ],
    link: [
      { rel: 'icon', type: 'image/x-icon', href: '/favicon.ico' }
    ]
  },
  /*
  ** Customize the progress bar color
  */
  loading: { color: '#3B8070' },
  plugins: ['~/plugins/element-ui', '~/plugins/jquery.base64.min',{ src: "~/plugins/wysiwyg.js", ssr: false },{ src: "~/plugins/vue-gallery", ssr: false }],
  css:['vue-wysiwyg/dist/vueWysiwyg.css'],
  /*
  ** Build configuration
  */
 modules: [
  '@nuxtjs/font-awesome',
 ],
  build: {
    /*
    ** Run ESLint on save
    */
   plugins: [
    new webpack.ProvidePlugin({
      '$': 'jquery',
      jQuery: 'jquery'
    })
  ],
    extend (config, { isDev, isClient }) {
      if (isDev && isClient) {
        config.module.rules.push({
          enforce: 'pre',
          test: /\.(js|vue)$/,
          loader: 'eslint-loader',
          exclude: /(node_modules)/
        })
      }
    }
  }
}
