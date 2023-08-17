const { defineConfig } = require('@vue/cli-service')

module.exports = defineConfig({
  transpileDependencies: true,

  css: {
    loaderOptions: {
      postcss: {
        postcssOptions: {
          ident: 'postcss',
          plugins: [
            require('tailwindcss'),
            require('autoprefixer')
          ]
        }
      }
    }
  }
})
