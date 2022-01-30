import {get, post} from "../utils/axios"

class Rest {

  // 加载数据结构
  async columns() {
    return get("/api/columns");
  }

  // 加载数据源适配器
  async select(query) {
    return post("/api/select", query)
  }

  // 添加一项
  async insert(record) {
    return post("/api/insert", record)
  }

  // 编辑一项
  async update(record) {
    return post("/api/update", record)
  }

  // 删除项
  async delete(record) {
    return post("/api/delete", record)
  }
}

export default new Rest()
