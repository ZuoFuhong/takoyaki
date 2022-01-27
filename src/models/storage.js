class StorageAPI {

  // Todo: 加载数据源适配器
  async select(query) {
    console.log("select: ", query)
    return [
      {
        isbn: '978-7-5063-6869-1',
        name: '《天幕红尘》',
        price: '48.00',
        author: '豆豆',
        status: '1',
        edition: '2013 年 6 月第 1 版',
        press: '作家出版社',
        address: '北京农展馆南里 10 号'
      },
      {
        isbn: '978-7-5063-6869-2',
        name: '《天幕红尘》',
        price: '48.00',
        author: '豆豆',
        status: '1',
        edition: '2013 年 6 月第 1 版',
        press: '作家出版社',
        address: '北京农展馆南里 10 号'
      },
      {
        isbn: '978-7-5063-6869-3',
        name: '《天幕红尘》',
        price: '48.00',
        author: '豆豆',
        status: '1',
        edition: '2013 年 6 月第 1 版',
        press: '作家出版社',
        address: '北京农展馆南里 10 号'
      },
      {
        isbn: '978-7-5063-6869-4',
        name: '《天幕红尘》',
        price: '48.00',
        author: '豆豆',
        status: '1',
        edition: '2013 年 6 月第 1 版',
        press: '作家出版社',
        address: '北京农展馆南里 10 号'
      },
      {
        isbn: '978-7-5063-6869-5',
        name: '《天幕红尘》',
        price: '48.00',
        author: '豆豆',
        status: '1',
        edition: '2013 年 6 月第 1 版',
        press: '作家出版社',
        address: '北京农展馆南里 10 号'
      }
    ]
  }

  // 添加一项
  async insert(record) {
    console.log("insert record: ", record)
  }

  // 编辑一项
  async update(record) {
    console.log("update record: ", record)
  }

  // 删除项
  async delete(record) {
    console.log("delete record: ", record)
  }
}

export default new StorageAPI()
