// vite.config.ts
import { defineConfig } from "file:///D:/Workspace/GoSpace/project/pangchat/source/app/node_modules/vite/dist/node/index.js";
import react from "file:///D:/Workspace/GoSpace/project/pangchat/source/app/node_modules/@vitejs/plugin-react/dist/index.mjs";
import postcsspxtoviewport from "file:///D:/Workspace/GoSpace/project/pangchat/source/app/node_modules/postcss-px-to-viewport/index.js";
import { resolve } from "path";
import { createSvgIconsPlugin } from "file:///D:/Workspace/GoSpace/project/pangchat/source/app/node_modules/vite-plugin-svg-icons/dist/index.mjs";
var __vite_injected_original_dirname = "D:\\Workspace\\GoSpace\\project\\pangchat\\source\\app";
var vite_config_default = defineConfig({
  plugins: [
    react(),
    createSvgIconsPlugin({
      // 指定需要缓存的图标文件夹
      iconDirs: [resolve(__vite_injected_original_dirname, "src/assets/icons")],
      // 指定symbolId格式
      symbolId: "icon-[dir]-[name]"
    })
  ],
  css: {
    postcss: {
      plugins: [
        postcsspxtoviewport({
          unitToConvert: "px",
          // 要转化的单位
          viewportWidth: 375,
          // UI设计稿的宽度，一般写 320
          // 下面的不常用，上面的常用
          unitPrecision: 6,
          // 转换后的精度，即小数点位数
          propList: ["*"],
          // 指定转换的css属性的单位，*代表全部css属性的单位都进行转换
          viewportUnit: "vw",
          // 指定需要转换成的视窗单位，默认vw
          fontViewportUnit: "vw",
          // 指定字体需要转换成的视窗单位，默认vw
          selectorBlackList: ["ignore-"],
          // 指定不转换为视窗单位的类名，
          minPixelValue: 1,
          // 默认值1，小于或等于1px则不进行转换
          mediaQuery: true,
          // 是否在媒体查询的css代码中也进行转换，默认false
          replace: true,
          // 是否转换后直接更换属性值
          landscape: false
          // 是否处理横屏情况
        })
      ]
    },
    preprocessorOptions: {
      less: {
        javascriptEnabled: true
      }
    }
  },
  resolve: {
    alias: {
      "@": resolve(__vite_injected_original_dirname, "src"),
      "@api": resolve(__vite_injected_original_dirname, "src/api"),
      "@img": resolve(__vite_injected_original_dirname, "src/assets/img")
    }
  },
  base: "./"
});
export {
  vite_config_default as default
};
//# sourceMappingURL=data:application/json;base64,ewogICJ2ZXJzaW9uIjogMywKICAic291cmNlcyI6IFsidml0ZS5jb25maWcudHMiXSwKICAic291cmNlc0NvbnRlbnQiOiBbImNvbnN0IF9fdml0ZV9pbmplY3RlZF9vcmlnaW5hbF9kaXJuYW1lID0gXCJEOlxcXFxXb3Jrc3BhY2VcXFxcR29TcGFjZVxcXFxwcm9qZWN0XFxcXHBhbmdjaGF0XFxcXHNvdXJjZVxcXFxhcHBcIjtjb25zdCBfX3ZpdGVfaW5qZWN0ZWRfb3JpZ2luYWxfZmlsZW5hbWUgPSBcIkQ6XFxcXFdvcmtzcGFjZVxcXFxHb1NwYWNlXFxcXHByb2plY3RcXFxccGFuZ2NoYXRcXFxcc291cmNlXFxcXGFwcFxcXFx2aXRlLmNvbmZpZy50c1wiO2NvbnN0IF9fdml0ZV9pbmplY3RlZF9vcmlnaW5hbF9pbXBvcnRfbWV0YV91cmwgPSBcImZpbGU6Ly8vRDovV29ya3NwYWNlL0dvU3BhY2UvcHJvamVjdC9wYW5nY2hhdC9zb3VyY2UvYXBwL3ZpdGUuY29uZmlnLnRzXCI7aW1wb3J0IHtkZWZpbmVDb25maWd9IGZyb20gJ3ZpdGUnXG5pbXBvcnQgcmVhY3QgZnJvbSAnQHZpdGVqcy9wbHVnaW4tcmVhY3QnXG5pbXBvcnQgcG9zdGNzc3B4dG92aWV3cG9ydCBmcm9tICdwb3N0Y3NzLXB4LXRvLXZpZXdwb3J0J1xuaW1wb3J0IHtyZXNvbHZlfSBmcm9tICdwYXRoJ1xuaW1wb3J0IHtjcmVhdGVTdmdJY29uc1BsdWdpbn0gZnJvbSBcInZpdGUtcGx1Z2luLXN2Zy1pY29uc1wiO1xuXG4vLyBodHRwczovL3ZpdGVqcy5kZXYvY29uZmlnL1xuZXhwb3J0IGRlZmF1bHQgZGVmaW5lQ29uZmlnKHtcbiAgICBwbHVnaW5zOiBbXG4gICAgICAgIHJlYWN0KCksXG4gICAgICAgIGNyZWF0ZVN2Z0ljb25zUGx1Z2luKHtcbiAgICAgICAgICAgIC8vIFx1NjMwN1x1NUI5QVx1OTcwMFx1ODk4MVx1N0YxM1x1NUI1OFx1NzY4NFx1NTZGRVx1NjgwN1x1NjU4N1x1NEVGNlx1NTkzOVxuICAgICAgICAgICAgaWNvbkRpcnM6IFtyZXNvbHZlKF9fZGlybmFtZSwgJ3NyYy9hc3NldHMvaWNvbnMnKV0sXG4gICAgICAgICAgICAvLyBcdTYzMDdcdTVCOUFzeW1ib2xJZFx1NjgzQ1x1NUYwRlxuICAgICAgICAgICAgc3ltYm9sSWQ6ICdpY29uLVtkaXJdLVtuYW1lXScsXG4gICAgICAgIH0pLFxuICAgIF0sXG4gICAgY3NzOiB7XG4gICAgICAgIHBvc3Rjc3M6IHtcbiAgICAgICAgICAgIHBsdWdpbnM6IFtcbiAgICAgICAgICAgICAgICBwb3N0Y3NzcHh0b3ZpZXdwb3J0KHtcbiAgICAgICAgICAgICAgICAgICAgdW5pdFRvQ29udmVydDogJ3B4JywgLy8gXHU4OTgxXHU4RjZDXHU1MzE2XHU3Njg0XHU1MzU1XHU0RjREXG4gICAgICAgICAgICAgICAgICAgIHZpZXdwb3J0V2lkdGg6IDM3NSwgLy8gVUlcdThCQkVcdThCQTFcdTdBM0ZcdTc2ODRcdTVCQkRcdTVFQTZcdUZGMENcdTRFMDBcdTgyMkNcdTUxOTkgMzIwXG4gICAgICAgICAgICAgICAgICAgIC8vIFx1NEUwQlx1OTc2Mlx1NzY4NFx1NEUwRFx1NUUzOFx1NzUyOFx1RkYwQ1x1NEUwQVx1OTc2Mlx1NzY4NFx1NUUzOFx1NzUyOFxuICAgICAgICAgICAgICAgICAgICB1bml0UHJlY2lzaW9uOiA2LCAvLyBcdThGNkNcdTYzNjJcdTU0MEVcdTc2ODRcdTdDQkVcdTVFQTZcdUZGMENcdTUzNzNcdTVDMEZcdTY1NzBcdTcwQjlcdTRGNERcdTY1NzBcbiAgICAgICAgICAgICAgICAgICAgcHJvcExpc3Q6IFsnKiddLCAvLyBcdTYzMDdcdTVCOUFcdThGNkNcdTYzNjJcdTc2ODRjc3NcdTVDNUVcdTYwMjdcdTc2ODRcdTUzNTVcdTRGNERcdUZGMEMqXHU0RUUzXHU4ODY4XHU1MTY4XHU5MEU4Y3NzXHU1QzVFXHU2MDI3XHU3Njg0XHU1MzU1XHU0RjREXHU5MEZEXHU4RkRCXHU4ODRDXHU4RjZDXHU2MzYyXG4gICAgICAgICAgICAgICAgICAgIHZpZXdwb3J0VW5pdDogJ3Z3JywgLy8gXHU2MzA3XHU1QjlBXHU5NzAwXHU4OTgxXHU4RjZDXHU2MzYyXHU2MjEwXHU3Njg0XHU4OUM2XHU3QTk3XHU1MzU1XHU0RjREXHVGRjBDXHU5RUQ4XHU4QkE0dndcbiAgICAgICAgICAgICAgICAgICAgZm9udFZpZXdwb3J0VW5pdDogJ3Z3JywgLy8gXHU2MzA3XHU1QjlBXHU1QjU3XHU0RjUzXHU5NzAwXHU4OTgxXHU4RjZDXHU2MzYyXHU2MjEwXHU3Njg0XHU4OUM2XHU3QTk3XHU1MzU1XHU0RjREXHVGRjBDXHU5RUQ4XHU4QkE0dndcbiAgICAgICAgICAgICAgICAgICAgc2VsZWN0b3JCbGFja0xpc3Q6IFsnaWdub3JlLSddLCAvLyBcdTYzMDdcdTVCOUFcdTRFMERcdThGNkNcdTYzNjJcdTRFM0FcdTg5QzZcdTdBOTdcdTUzNTVcdTRGNERcdTc2ODRcdTdDN0JcdTU0MERcdUZGMENcbiAgICAgICAgICAgICAgICAgICAgbWluUGl4ZWxWYWx1ZTogMSwgLy8gXHU5RUQ4XHU4QkE0XHU1MDNDMVx1RkYwQ1x1NUMwRlx1NEU4RVx1NjIxNlx1N0I0OVx1NEU4RTFweFx1NTIxOVx1NEUwRFx1OEZEQlx1ODg0Q1x1OEY2Q1x1NjM2MlxuICAgICAgICAgICAgICAgICAgICBtZWRpYVF1ZXJ5OiB0cnVlLCAvLyBcdTY2MkZcdTU0MjZcdTU3MjhcdTVBOTJcdTRGNTNcdTY3RTVcdThCRTJcdTc2ODRjc3NcdTRFRTNcdTc4MDFcdTRFMkRcdTRFNUZcdThGREJcdTg4NENcdThGNkNcdTYzNjJcdUZGMENcdTlFRDhcdThCQTRmYWxzZVxuICAgICAgICAgICAgICAgICAgICByZXBsYWNlOiB0cnVlLCAvLyBcdTY2MkZcdTU0MjZcdThGNkNcdTYzNjJcdTU0MEVcdTc2RjRcdTYzQTVcdTY2RjRcdTYzNjJcdTVDNUVcdTYwMjdcdTUwM0NcbiAgICAgICAgICAgICAgICAgICAgbGFuZHNjYXBlOiBmYWxzZSAvLyBcdTY2MkZcdTU0MjZcdTU5MDRcdTc0MDZcdTZBMkFcdTVDNEZcdTYwQzVcdTUxQjVcbiAgICAgICAgICAgICAgICB9KSxcbiAgICAgICAgICAgIF1cbiAgICAgICAgfSxcbiAgICAgICAgcHJlcHJvY2Vzc29yT3B0aW9uczoge1xuICAgICAgICAgICAgbGVzczoge1xuICAgICAgICAgICAgICAgIGphdmFzY3JpcHRFbmFibGVkOiB0cnVlLFxuICAgICAgICAgICAgfVxuICAgICAgICB9XG4gICAgfSxcbiAgICByZXNvbHZlOiB7XG4gICAgICAgIGFsaWFzOiB7XG4gICAgICAgICAgICAnQCc6IHJlc29sdmUoX19kaXJuYW1lLCAnc3JjJyksXG4gICAgICAgICAgICAnQGFwaSc6IHJlc29sdmUoX19kaXJuYW1lLCAnc3JjL2FwaScpLFxuICAgICAgICAgICAgJ0BpbWcnOiByZXNvbHZlKF9fZGlybmFtZSwgJ3NyYy9hc3NldHMvaW1nJyksXG4gICAgICAgIH1cbiAgICB9LFxuICAgIGJhc2U6ICcuLycsXG59KVxuIl0sCiAgIm1hcHBpbmdzIjogIjtBQUFnVixTQUFRLG9CQUFtQjtBQUMzVyxPQUFPLFdBQVc7QUFDbEIsT0FBTyx5QkFBeUI7QUFDaEMsU0FBUSxlQUFjO0FBQ3RCLFNBQVEsNEJBQTJCO0FBSm5DLElBQU0sbUNBQW1DO0FBT3pDLElBQU8sc0JBQVEsYUFBYTtBQUFBLEVBQ3hCLFNBQVM7QUFBQSxJQUNMLE1BQU07QUFBQSxJQUNOLHFCQUFxQjtBQUFBO0FBQUEsTUFFakIsVUFBVSxDQUFDLFFBQVEsa0NBQVcsa0JBQWtCLENBQUM7QUFBQTtBQUFBLE1BRWpELFVBQVU7QUFBQSxJQUNkLENBQUM7QUFBQSxFQUNMO0FBQUEsRUFDQSxLQUFLO0FBQUEsSUFDRCxTQUFTO0FBQUEsTUFDTCxTQUFTO0FBQUEsUUFDTCxvQkFBb0I7QUFBQSxVQUNoQixlQUFlO0FBQUE7QUFBQSxVQUNmLGVBQWU7QUFBQTtBQUFBO0FBQUEsVUFFZixlQUFlO0FBQUE7QUFBQSxVQUNmLFVBQVUsQ0FBQyxHQUFHO0FBQUE7QUFBQSxVQUNkLGNBQWM7QUFBQTtBQUFBLFVBQ2Qsa0JBQWtCO0FBQUE7QUFBQSxVQUNsQixtQkFBbUIsQ0FBQyxTQUFTO0FBQUE7QUFBQSxVQUM3QixlQUFlO0FBQUE7QUFBQSxVQUNmLFlBQVk7QUFBQTtBQUFBLFVBQ1osU0FBUztBQUFBO0FBQUEsVUFDVCxXQUFXO0FBQUE7QUFBQSxRQUNmLENBQUM7QUFBQSxNQUNMO0FBQUEsSUFDSjtBQUFBLElBQ0EscUJBQXFCO0FBQUEsTUFDakIsTUFBTTtBQUFBLFFBQ0YsbUJBQW1CO0FBQUEsTUFDdkI7QUFBQSxJQUNKO0FBQUEsRUFDSjtBQUFBLEVBQ0EsU0FBUztBQUFBLElBQ0wsT0FBTztBQUFBLE1BQ0gsS0FBSyxRQUFRLGtDQUFXLEtBQUs7QUFBQSxNQUM3QixRQUFRLFFBQVEsa0NBQVcsU0FBUztBQUFBLE1BQ3BDLFFBQVEsUUFBUSxrQ0FBVyxnQkFBZ0I7QUFBQSxJQUMvQztBQUFBLEVBQ0o7QUFBQSxFQUNBLE1BQU07QUFDVixDQUFDOyIsCiAgIm5hbWVzIjogW10KfQo=
