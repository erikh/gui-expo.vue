const path = require("path");
const { VueLoaderPlugin } = require("vue-loader");

module.exports = (env, options) => ({
  mode: process.env.NODE_ENV || "production",
  entry: "./src",
  module: {
    rules: [
      {
        test: /\.vue$/,
        loader: "vue-loader",
      },
      {
        test: /\.css$/,
        use: ["vue-style-loader", "css-loader"],
      },
    ],
  },
  plugins: [new VueLoaderPlugin()],
  output: {
    filename: "bundle.js",
    path: path.resolve(__dirname, "build"),
    clean: true,
  },
  resolve: {
    alias: {
      // this embeds the runtime compiler and avoids making us use the vue
      // compiler. it's just to keep things flowing.
      vue: "vue/dist/vue.esm-bundler.js",
    },
  },
});
