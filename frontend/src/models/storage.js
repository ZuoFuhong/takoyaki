class StorageAPI {

  // 表格字段定义【推荐】
  // 1. `column` 字段名不能重复，要去除前后置空格，使用字母+数字的组合，为了布局美观，长度控制在的 10 个字符内。
  // 2. `fieldName` 字段名由 BMP 字符组成，为了布局美观，中文字符长度控制在 5 个字符以内。
  // 3. 字段的排序按照字段定义在列表中的顺序。
  // 4. 字段的值类型，目前仅支持字符串，若定义其它类型（如：数字、布尔）会产生非预期结果。
  columns() {
    return [
      {
        "column": "isbn",       // 字段
        "fieldName": "ISBN",    // 字段名
        "search": true,         // 支持搜索，默认 false
        "searchMode": "equal"   // 搜索模式：like-模糊匹配 equal-精准匹配, 默认 equal
      },
      {
        "column": "name",
        "fieldName": "书名",
        "search": true,
        "searchMode": "like"
      },
      {
        "column": "price",
        "fieldName": "定价",
      },
      {
        "column": "author",
        "fieldName": "作者",
      },
      {
        "column": "edition",
        "fieldName": "版次",
      },
      {
        "column": "press",
        "fieldName": "出版社",
      },
      {
        "column": "address",
        "fieldName": "社址"
      }
    ]
  }

  // 加载数据源适配器
  async select(query) {
    console.log("select: ", query)
    return [
      {
        isbn: '978-7-5063-6869-1',
        name: '《天幕红尘》',
        price: '48.00',
        author: '豆豆',
        edition: '2013 年 6 月第 1 版',
        press: '作家出版社',
        address: '北京农展馆南里 10 号'
      },
      {
        isbn: '978-7-5063-6869-2',
        name: '《天幕红尘》',
        price: '48.00',
        author: '豆豆',
        edition: '2013 年 6 月第 1 版',
        press: '作家出版社',
        address: '北京农展馆南里 10 号'
      },
      {
        isbn: '978-7-5063-6869-3',
        name: '《天幕红尘》',
        price: '48.00',
        author: '豆豆',
        edition: '2013 年 6 月第 1 版',
        press: '作家出版社',
        address: '北京农展馆南里 10 号'
      },
      {
        isbn: '978-7-5063-6869-4',
        name: '《天幕红尘》',
        price: '48.00',
        author: '豆豆',
        edition: '2013 年 6 月第 1 版',
        press: '作家出版社',
        address: '北京农展馆南里 10 号'
      },
      {
        isbn: '978-7-5063-6869-5',
        name: '《天幕红尘》',
        price: '48.00',
        author: '豆豆',
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
