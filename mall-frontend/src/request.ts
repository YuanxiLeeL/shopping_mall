import axios, { AxiosInstance } from "axios";

const myaxios: AxiosInstance = axios.create({
  baseURL: "http://localhost:8080", // 后端 API 地址
  timeout: 1000, // 请求超时时间
  // headers: { "X-Custom-Header": "foobar" }, // 自定义头部
});

// 添加请求拦截器
myaxios.interceptors.request.use(
  function (config) {
    console.log(config);
    // 在发送请求之前，获取 token 并添加到 Authorization 头部
    const token = localStorage.getItem("token"); // 从 localStorage 获取 token

    // 如果 Token 存在且请求不是登录接口，则添加 Token 到请求头
    if (token && !config.url?.includes("/api/auth/login")) {
      config.headers["Authorization"] = `${token}`;
      return config;
    }

    if (token) {
      config.headers["Authorization"] = `${token}`; // 添加 Authorization 头部
    }
    return config;
  },
  function (error) {
    // 请求错误时的处理
    return Promise.reject(error);
  }
);

// 添加响应拦截器
myaxios.interceptors.response.use(
  function (response) {
    // 2xx 范围内的状态码都会触发该函数。
    // 对响应数据做点什么
    console.log(response);
    const { data } = response;
    console.log(data);

    // 如果响应是 401 (未授权)，即 Token 无效或过期
    if (data.code === 40100) {
      // 不是用户信息接口且没有在登录页时，跳转到登录页面
      if (
        !response.request.responseURL.includes("api/auth/current") &&
        !window.location.pathname.includes("api/auth/login")
      ) {
        window.location.href = `api/auth/login?redirect=${window.location.pathname}`;
      }
    }
    return response;
  },
  function (error) {
    // 超出 2xx 范围的状态码都会触发该函数。
    // 对响应错误做点什么
    return Promise.reject(error);
  }
);

export default myaxios;
