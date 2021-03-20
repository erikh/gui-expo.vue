const path = require("path");

module.exports = (env, options) => ({
  mode: "production",
  entry: "./src",
  plugins: [],
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
