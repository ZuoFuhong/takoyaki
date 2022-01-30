import axios from 'axios'

const config = {
  baseURL: '/',
  timeout: 5 * 1000,
  crossDomain: true,
  validateStatus(status) {
    return status >= 200 && status < 510
  },
}

// 创建请求实例
const _axios = axios.create(config)

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

export default _axios
