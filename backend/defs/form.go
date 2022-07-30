package defs

type DataSourceForm struct {
	Name       string `json:"name"`
	Host       string `json:"host"`
	Port       string `json:"port"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Database   string `json:"database"`
	CreateTime string `json:"create_time"`
	UpdateTime string `json:"update_time"`
}

type DataSourceSearchForm struct {
	Name     string `json:"name"`
	Host     string `json:"host"`
	Page     int    `json:"page"`
	PageSize int    `json:"page_size"`
}

type PageForm struct {
	Name       string        `json:"name"`
	DataSource string        `json:"data_source"`
	TableName  string        `json:"table_name"`
	PrimaryKey string        `json:"primary_key"`
	Columns    []*PageColumn `json:"columns"`
	CreateTime string        `json:"create_time"`
	UpdateTime string        `json:"update_time"`
}

type PageColumn struct {
	ColumnName string `json:"column_name"`
	FieldName  string `json:"field_name"`
	Search     bool   `json:"search"`
}

type PageSearchForm struct {
	Name       string `json:"name"`
	DataSource string `json:"data_source"`
	Page       int    `json:"page"`
	PageSize   int    `json:"page_size"`
}

type SelectForm struct {
	Conditions map[string]string `json:"conditions"`
	PageName   string            `json:"page_name"`
	Page       int               `json:"page"`
	PageSize   int               `json:"page_size"`
}

type InsertForm struct {
	PageName string            `json:"page_name"`
	Record   map[string]string `json:"record"`
}

type UpdateForm struct {
	PageName     string            `json:"page_name"`
	PrimaryKeyId string            `json:"primary_key_id"`
	Record       map[string]string `json:"record"`
}

type DeleteForm struct {
	PageName     string `json:"page_name"`
	PrimaryKeyId string `json:"primary_key_id"`
}
