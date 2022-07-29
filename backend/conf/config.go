package conf

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"sync"
)

var (
	once   sync.Once
	config *Config
)

type Config struct {
	Db         Db       // Db配置
	TableName  string   // 表名
	PrimaryKey string   // 主键
	Fields     []*Field // 展示的字段
}

// Db 配置
type Db struct {
	Addr   string
	User   string
	Pass   string
	Dbname string
}

// Field 表格字段定义【推荐】
// 1. `column` 字段名不能重复，要去除前后置空格，使用字母+数字的组合，为了布局美观，长度控制在的 10 个字符内。
// 2. `fieldName` 字段名由 BMP 字符组成，为了布局美观，中文字符长度控制在 5 个字符以内。
// 3. 字段的排序按照字段定义在列表中的顺序。
// 4. 字段的值类型，目前仅支持字符串，若定义其它类型（如：数字、布尔）会产生非预期结果。
type Field struct {
	Column    string `json:"column"`    // 字段
	FieldName string `json:"fieldName"` // 字段名
	Search    bool   `json:"search"`    // 支持搜索，默认 false
}

// GetConfig 获取配置
func GetConfig() *Config {
	once.Do(func() {
		configStr := loadConfigFile()
		config = new(Config)
		if err := json.Unmarshal([]byte(configStr), config); err != nil {
			log.Panic(err)
		}
	})
	return config
}

// loadConfigFile 加载配置文件
func loadConfigFile() string {
	bytes, err := ioutil.ReadFile("./conf/config.json")
	if err != nil {
		log.Panic(err)
	}
	return string(bytes)
}
