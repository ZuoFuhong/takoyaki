import axios from 'axios'

const config = {
  baseURL: '/',
  timeout: 5 * 1000,
  crossDomain: true
}

// 创建请求实例
const _axios = axios.create(config)

// 添加响应拦截器
_axios.interceptors.response.use(function (res) {
  return res.data
}, function (error) {
  console.log(error)
  // 超出 2xx 范围的状态码都会触发该函数。
  // 对响应错误做点什么
  return Promise.reject(error);
});

export function post(url, data = {}, params = {}) {
  return _axios({
    method: 'post',
    url,
    data,
    params,
  })
}

export function get(url, params = {}) {
  return _axios({
    method: 'get',
    url,
    params,
  })
}

export function put(url, data = {}, params = {}) {
  return _axios({
    method: 'put',
    url,
    params,
    data,
  })
}

export function _delete(url, params = {}) {
  return _axios({
    method: 'delete',
    url,
    params,
  })
}

export default _axios
