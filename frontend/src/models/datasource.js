import {get, _delete, post} from "../utils/axios"
import {Base64} from 'js-base64'

class DataSource {

  async getDataSources(data) {
    let raw = Base64.encode(JSON.stringify(data))
    let query = encodeURIComponent(raw)
    return get(`/api/datasources?query=${query}`)
  }

  async saveDataSource(data) {
    return post('/api/datasource', data)
  }

  async deleteDataSource(source) {
    return _delete(`/api/datasource?source=${source}`)
  }

  async getAllDataSource() {
    return get('/api/datasource_option')
  }
}

export default new DataSource()
