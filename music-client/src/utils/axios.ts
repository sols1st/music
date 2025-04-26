import axios, { type Method } from "axios";
import { ElMessage } from 'element-plus'

const BASE_URL = process.env.NODE_HOST;

axios.defaults.baseURL = BASE_URL;

axios.defaults.headers.post["Content-Type"] = "application/json;charset=UTF-8";
axios.defaults.timeout = 10000;

// const store = useAccountStore()

// axios.interceptors.request.use(
//   config => {
//     const token = store.token
//     config.headers.authorization = token
//     return config
//   },
//   error => {
//     return Promise.reject(error)
//   }
// )

axios.interceptors.response.use(
  (response) => {
    return response;
  },
  (error) => {
    if (
      error.code === "ECONNABORTED" ||
      error.message.indexOf("timeout") !== -1 ||
      error.message === "Network Error"
    ) {
      ElMessage("网络异常");
      return Promise.reject(error);
    }
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
    return Promise.reject(error);
  }
);

export default function Axios(
  url: string,
  data: any,
  method: Method,
  toast: boolean = true
) {
  return new Promise((resolve, reject) => {
    if (method === "get") {
      axios({
        method,
        url,
        params: data,
      })
        .then((res) => {
          if (res.data.code === 200) {
            if (toast) ElMessage(res.data.msg);
            resolve(res.data.data);
          } else {
            ElMessage(res.data.msg);
            reject(res.data.msg);
          }
        })
        .catch((err) => {
          reject(err);
        });
    } else {
      axios({
        method,
        url,
        data,
      })
        .then((res) => {
          if (res.data.code === 200) {
            if (toast) ElMessage(res.data.msg);
            resolve(res.data.data);
          } else {
            ElMessage(res.data.msg);
            reject(res.data.msg);
          }
        })
        .catch((err) => {
          reject(err);
        });
    }
  });
}
