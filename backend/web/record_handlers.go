package web

import (
	"encoding/json"
	"log"
	"net/http"
	"takoyaki/bolt"
	"takoyaki/defs"
	"takoyaki/errcode"
	"takoyaki/mysql"
	"time"
)

// columnsHandle get page column config
func columnsHandle(w http.ResponseWriter, r *http.Request) {
	pageName := r.URL.Query().Get("page_name")
	page, err := bolt.GetPage(pageName)
	if err != nil {
		// query degrade, ignore error
		Ok(w, "")
		return
	}
	Ok(w, page.Columns)
}

// selectHandle dynamic select
func selectHandle(w http.ResponseWriter, r *http.Request) {
	// 1.parse request params
	form := new(defs.SelectForm)
	if err := json.NewDecoder(r.Body).Decode(form); err != nil {
		log.Printf("call decoder.Decode failed, err: %v", err)
		Error(w, errcode.BadRequestParam, "bad request")
		return
	}
	// 2.query db record
	records, total, err := mysql.Select(form)
	if err != nil {
		log.Printf("call mysql.Select failed, err: %v", err)
		Error(w, errcode.InternalServerError, "系统繁忙")
		return
	}
	// 3.render value type
	voList := make([]map[string]interface{}, 0)
	for _, record := range records {
		vo := make(map[string]interface{})
		for k, v := range record {
			switch vt := v.(type) {
			case time.Time:
				// 格式化日期
				vo[k] = vt.Format("2006-01-02 15:04:05")
			default:
				vo[k] = v
			}
		}
		voList = append(voList, vo)
	}
	data := map[string]interface{}{
		"list":  voList,
		"total": total,
	}
	Ok(w, data)
}

// insertHandle add once record
func insertHandle(w http.ResponseWriter, r *http.Request) {
	// 1.parse request params
	form := new(defs.InsertForm)
	if err := json.NewDecoder(r.Body).Decode(form); err != nil {
		log.Printf("call decoder.Decode failed, err: %v", err)
		Error(w, errcode.BadRequestParam, "bad request")
		return
	}
	// 2.write db
	if err := mysql.Insert(form); err != nil {
		log.Printf("call mysql.Insert failed, err: %v", err)
		Error(w, errcode.InternalServerError, "系统繁忙")
		return
	}
	Ok(w, "ok")
}

// updateHandle 更新数据
func updateHandle(w http.ResponseWriter, r *http.Request) {
	// 1.读取请求参数
	form := new(defs.UpdateForm)
	if err := json.NewDecoder(r.Body).Decode(form); err != nil {
		log.Printf("call decoder.Decode failed, err: %v", err)
		Error(w, errcode.BadRequestParam, "bad request")
		return
	}
	// 2.write db
	if err := mysql.Update(form); err != nil {
		log.Printf("call mysql.Update failed, err: %v", err)
		Error(w, errcode.InternalServerError, "系统繁忙")
		return
	}
	Ok(w, "ok")
}

// deleteHandle 删除数据
func deleteHandle(w http.ResponseWriter, r *http.Request) {
	// 1.parse request params
	form := new(defs.DeleteForm)
	if err := json.NewDecoder(r.Body).Decode(form); err != nil {
		log.Printf("call decoder.Decode failed, err: %v", err)
		Error(w, errcode.BadRequestParam, "bad request")
		return
	}
	// 2.write db
	if err := mysql.Delete(form); err != nil {
		log.Printf("call mysql.Delete failed, err: %v", err)
		Error(w, errcode.InternalServerError, "系统繁忙")
		return
	}
	Ok(w, "ok")
}
