package web

import (
	"encoding/json"
	"log"
	"net/http"
	"takoyaki/conf"
	"takoyaki/defs"
	"takoyaki/provider/dao"
	"time"
)

// columnsHandle 查询表结构
func columnsHandle(w http.ResponseWriter, r *http.Request) {
	primaryKey := conf.GetConfig().PrimaryKey
	fields := conf.GetConfig().Fields
	// 隐式传入主键
	fields = append(fields, &conf.Field{
		Column:    primaryKey,
		FieldName: "PrimaryKey",
	})
	SendNormalJsonResponse(w, fields)
}

// selectHandler 查询数据
func selectHandle(w http.ResponseWriter, r *http.Request) {
	// 1.读取请求参数
	decoder := json.NewDecoder(r.Body)
	params := make(map[string]interface{})
	if err := decoder.Decode(&params); err != nil {
		log.Printf("call decoder.Decode failed, err: %v", err)
		SendErrorResponse(w, defs.ErrorInternalFaults)
		return
	}
	// 2.提取分页参数
	// Tips: json 中的数字类型没有对应 int, 解析出来是 float64
	var offset, limit int
	if val, ok := params["offset"]; ok {
		offset = int(val.(float64))
		delete(params, "offset")
	}
	if val, ok := params["limit"]; ok {
		limit = int(val.(float64))
		delete(params, "limit")
	}
	// 3.按配置过滤入参
	fieldMap := getFieldConfigMap()
	conditions := make(map[string]interface{})
	for k, v := range params {
		if _, ok := fieldMap[k]; ok {
			switch vv := v.(type) {
			case float64:
				conditions[k] = int(vv)
			default:
				conditions[k] = v
			}
		}
	}
	// 4.查库
	records, err := dao.Select(conditions, offset, limit)
	if err != nil {
		log.Printf("call dao.Select failed, err: %v", err)
		SendErrorResponse(w, defs.ErrorInternalFaults)
		return
	}
	// 5.按配置过滤回包
	voList := make([]map[string]interface{}, 0)
	for _, record := range records {
		vo := make(map[string]interface{})
		for k, v := range record {
			if _, ok := fieldMap[k]; ok {
				switch vt := v.(type) {
				case time.Time:
					// 格式化日期
					vo[k] = vt.Format("2006-01-02 15:04:05")
				default:
					vo[k] = v
				}
			}
		}
		voList = append(voList, vo)
	}
	// 6.回包
	SendNormalJsonResponse(w, voList)
}

// insertHandle 增加数据
func insertHandle(w http.ResponseWriter, r *http.Request) {
	// 1.读取请求参数
	decoder := json.NewDecoder(r.Body)
	params := make(map[string]interface{})
	if err := decoder.Decode(&params); err != nil {
		log.Printf("call decoder.Decode failed, err: %v", err)
		SendErrorResponse(w, defs.ErrorInternalFaults)
		return
	}
	// 2.按配置过滤字段
	fieldMap := getFieldConfigMap()
	record := make(map[string]interface{})
	for k, v := range params {
		if _, ok := fieldMap[k]; ok {
			switch vv := v.(type) {
			case float64:
				record[k] = int(vv)
			default:
				record[k] = v
			}
		}
	}
	// 3.写库
	if err := dao.Insert(record); err != nil {
		log.Printf("call dao.Insert failed, err: %v", err)
		SendErrorResponse(w, defs.ErrorInternalFaults)
		return
	}
	// 4.回包
	SendNormalResponse(w, "ok", http.StatusOK)
}

// updateHandle 更新数据
func updateHandle(w http.ResponseWriter, r *http.Request) {
	// 1.读取请求参数
	decoder := json.NewDecoder(r.Body)
	params := make(map[string]interface{})
	if err := decoder.Decode(&params); err != nil {
		log.Printf("call decoder.Decode failed, err: %v", err)
		SendErrorResponse(w, defs.ErrorInternalFaults)
		return
	}
	// 2.按配置过滤字段
	primaryKey := conf.GetConfig().PrimaryKey
	fieldMap := getFieldConfigMap()
	record := make(map[string]interface{})
	conditions := make(map[string]interface{})
	for k, v := range params {
		if _, ok := fieldMap[k]; ok {
			switch vv := v.(type) {
			case float64:
				record[k] = int(vv)
			default:
				record[k] = v
			}
		}
		if k == primaryKey {
			switch vv := v.(type) {
			case float64:
				conditions[k] = int(vv)
			default:
				conditions[k] = v
			}
		}
	}
	if len(conditions) == 0 {
		SendErrorResponse(w, defs.ErrorResouceNotFound)
		return
	}
	// 3.检查记录
	records, err := dao.Select(conditions, 0, 1)
	if err != nil {
		log.Printf("call dao.Select failed, err: %v", err)
		SendErrorResponse(w, defs.ErrorInternalFaults)
		return
	}
	if len(records) == 0 {
		SendErrorResponse(w, defs.ErrorResouceNotFound)
		return
	}
	// 4.写库
	if err := dao.Update(conditions, record); err != nil {
		log.Printf("call dao.Update failed, err: %v", err)
		SendErrorResponse(w, defs.ErrorInternalFaults)
		return
	}
	// 5.回包
	SendNormalResponse(w, "ok", http.StatusOK)
}

// deleteHandle 删除数据
func deleteHandle(w http.ResponseWriter, r *http.Request) {
	// 1.读取请求参数
	decoder := json.NewDecoder(r.Body)
	params := make(map[string]interface{})
	if err := decoder.Decode(&params); err != nil {
		log.Printf("call decoder.Decode failed, err: %v", err)
		SendErrorResponse(w, defs.ErrorInternalFaults)
		return
	}
	// 2.按配置过滤字段
	primaryKey := conf.GetConfig().PrimaryKey
	conditions := make(map[string]interface{})
	for k, v := range params {
		if k == primaryKey {
			switch vv := v.(type) {
			case float64:
				conditions[k] = int(vv)
			default:
				conditions[k] = v
			}
			break
		}
	}
	if len(conditions) == 0 {
		SendErrorResponse(w, defs.ErrorResouceNotFound)
		return
	}
	// 3.检查记录
	records, err := dao.Select(conditions, 0, 1)
	if err != nil {
		log.Printf("call dao.Select failed, err: %v", err)
		SendErrorResponse(w, defs.ErrorInternalFaults)
		return
	}
	if len(records) == 0 {
		SendErrorResponse(w, defs.ErrorResouceNotFound)
		return
	}
	// 4.删除
	if err := dao.Delete(conditions); err != nil {
		log.Printf("call dao.Delete failed, err: %v", err)
		SendErrorResponse(w, defs.ErrorInternalFaults)
		return
	}
	// 5.回包
	SendNormalResponse(w, "ok", http.StatusOK)
}

// 配置的字段
func getFieldConfigMap() map[string]string {
	fields := conf.GetConfig().Fields
	fieldMap := make(map[string]string)
	for _, v := range fields {
		fieldMap[v.Column] = v.FieldName
	}
	return fieldMap
}
