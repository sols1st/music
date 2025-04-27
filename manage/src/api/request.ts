import axios from "axios";
import { ElMessage } from "element-plus";

const BASE_URL = process.env.NODE_HOST;

axios.defaults.timeout = 5000; // 超时时间设置
axios.defaults.baseURL = BASE_URL;
axios.defaults.headers.post["Content-Type"] = "application/json;charset=UTF-8";

axios.interceptors.response.use(
  (response) => {
    if (response.status === 200) {
      return Promise.resolve(response);
    } else {
      return Promise.reject(response);
    }
  },
  (error) => {
    if (error.response.status) {
      switch (error.response.status) {
      case 403:
        ElMessage("拒绝访问(403)");
        break;
      case 404:
        ElMessage("资源不存在(404)");
        break;
      case 408:
        ElMessage("请求超时(404)");
        break;
      case 500:
        ElMessage("服务器错误(500)");
        break;
      case 501:
        ElMessage("服务未实现(501)");
        break;
      case 502:
        ElMessage("网络错误(502)");
        break;
      case 503:
        ElMessage("服务不可用(503)");
        break;
      case 504:
        ElMessage("网络超时(504)");
        break;
      case 505:
        ElMessage("HTTP版本不受支持(505)");
        break;
      default:
        break;
      }
      return Promise.reject(error.response);
    }
  }
);

export function getBaseURL() {
  return BASE_URL;
}

/**
 * 封装get方法
 * @param url
 * @param data
 * @returns {Promise}
 */
export function get(url, params?: object) {
  return new Promise((resolve, reject) => {
    axios.get(url, params).then(
      (response) => resolve(response.data),
      (error) => reject(error)
    );
  });
}

/**
 * 封装post请求
 * @param url
 * @param data
 * @returns {Promise}
 */
export function post(url, data = {}) {
  return new Promise((resolve, reject) => {
    axios.post(url, data).then(
      (response) => resolve(response.data),
      (error) => reject(error)
    );
  });
}

/**
 * 封装delete请求
 * @param url
 * @param data
 * @returns {Promise}
 */
export function deletes(url, data = {}) {
  return new Promise((resolve, reject) => {
    axios.delete(url, data).then(
      (response) => resolve(response.data),
      (error) => reject(error)
    );
  });
}

/**
 * 封装put请求
 * @param url
 * @param data
 * @returns {Promise}
 */
export function put(url, data = {}) {
  return new Promise((resolve, reject) => {
    axios.put(url, data).then(
      (response) => resolve(response.data),
      (error) => reject(error)
    );
  });
}
