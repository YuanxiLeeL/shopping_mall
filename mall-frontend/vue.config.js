const { defineConfig } = require("@vue/cli-service");
module.exports = defineConfig({
  transpileDependencies: true,
});

module.exports = {
  devServer: {
    // proxy: {
    //   "/api": {
    //     target: "http://localhost:8080/", // 后端服务器地址
    //     changeOrigin: true, // 允许跨域
    //     secure: false, // 如果是 https，需要设置为 true
    //     // pathRewrite: { "^/api": "" }, // 重写路径，去掉 /api 前缀
    //   },
    // },
  },
};
