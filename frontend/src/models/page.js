import {get, _delete, post} from "../utils/axios"
import {Base64} from 'js-base64'

class Page {

  async getPages(data) {
    let raw = Base64.encode(JSON.stringify(data))
    let query = encodeURIComponent(raw)
    return get(`/api/pages?query=${query}`)
  }

  async savePage(data) {
    return post('/api/page', data)
  }

  async deletePage(pageName) {
    return _delete(`/api/page?page_name=${pageName}`)
  }
}

export default new Page()
