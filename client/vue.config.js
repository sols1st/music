const { defineConfig } = require('@vue/cli-service');

module.exports = defineConfig({
  transpileDependencies: true,
  // 添加代理配置
  devServer: {
    proxy: {
      '/api': {
        target: 'http://localhost:3456', // Gin 后端地址
        changeOrigin: true, // 更改请求头中的 Origin 为目标地址
        pathRewrite: { '^/api': '' }, // 将 /api 前缀移除
      },
    },
  },
  chainWebpack: config => {
    config.plugin('define').tap(definitions => {
      Object.assign(definitions[0]['process.env'], {
        NODE_HOST: '"http://localhost:3456"', // 保持后端地址
        // NODE_HOST: '"http://localhost:8888"',
      });
      return definitions;
    });
  },
});