import axios from "axios";
import {ElMessage}  from "element-plus";

// 创建实例
const instance = axios.create({
  baseURL: "http://192.168.1.71:8080/api",
  timeout: 10000,
});

// 请求拦截器
instance.interceptors.request.use(
  (config) => {
      const token = window.sessionStorage.getItem('token');
      if(token) {
          config.headers = {
              'Authorization': 'token ' + token
          }
      }
      return config
  },
  (error) => {
    return Promise.reject(error);
  }
);

// 响应拦截器
instance.interceptors.response.use(
  (res) => {
    // 统一处理非200状态错误提示
    if (res.data.code != 200) {
      ElMessage.error(res.data.msg);
    }
    return res;
  },
  (error) => {
    console.log(error);
    ElMessage.error("请求服务接口错误！");
    return Promise.reject(error);
  }
);

export default instance;
