const { defineConfig } = require('@vue/cli-service')
module.exports = defineConfig({
  transpileDependencies: [
    'vuetify'
  ],
  devServer: {
    hot: true,
    host: "0.0.0.0" ,
    port: 8080,
  },
})