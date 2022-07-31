package defs

import (
	"encoding/base64"
	"encoding/json"
	"log"
)

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

// DecodeDataSourceQuery decode request query param
func DecodeDataSourceQuery(query string) (*DataSourceSearchForm, error) {
	bytes, err := base64.StdEncoding.DecodeString(query)
	if err != nil {
		log.Printf("call base64 decode failed, err: %v\n", err)
		return nil, err
	}
	form := new(DataSourceSearchForm)
	if err := json.Unmarshal(bytes, form); err != nil {
		log.Printf("call json.Unmarshal failed, err: %v\n", err)
		return nil, err
	}
	return form, nil
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

// DecodePageSearchQuery decode request query param
func DecodePageSearchQuery(query string) (*PageSearchForm, error) {
	bytes, err := base64.StdEncoding.DecodeString(query)
	if err != nil {
		log.Printf("call base64 decode failed, err: %v\n", err)
		return nil, err
	}
	form := new(PageSearchForm)
	if err := json.Unmarshal(bytes, form); err != nil {
		log.Printf("call json.Unmarshal failed, err: %v\n", err)
		return nil, err
	}
	return form, nil
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
