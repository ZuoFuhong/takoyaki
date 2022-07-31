import {get, post, put} from "../utils/axios"
import {Base64} from 'js-base64'

class Rest {

  // 加载数据结构
  async columns(pageName) {
    return get(`/api/columns?page_name=${pageName}`);
  }

  // 加载数据源适配器
  async select(data) {
    let raw = Base64.encode(JSON.stringify(data))
    let query = encodeURIComponent(raw)
    return get(`/api/select?query=${query}`)
  }

  // 添加一项
  async insert(data) {
    return post("/api/insert", data)
  }

  // 更新一项
  async update(data) {
    return put("/api/update", data)
  }

  // 删除项
  async delete(data) {
    return put("/api/delete", data)
  }
}

export default new Rest()
