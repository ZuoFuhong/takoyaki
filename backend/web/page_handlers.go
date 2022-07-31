package web

import (
	"encoding/json"
	"log"
	"net/http"
	"takoyaki/bolt"
	"takoyaki/defs"
	"takoyaki/errcode"
)

// AddPage add page
func AddPage(w http.ResponseWriter, r *http.Request) {
	// 1.parse request params
	form := new(defs.PageForm)
	if err := json.NewDecoder(r.Body).Decode(form); err != nil {
		log.Printf("call decoder.Decode failed, err: %v", err)
		Error(w, errcode.BadRequestParam, "bad request")
		return
	}
	// 2.check the data source is valid
	_, err := bolt.GetDataSource(form.DataSource)
	if err != nil {
		Error(w, errcode.InvalidDataSource, "invalid data source")
		return
	}
	// 3.write to bolt
	_ = bolt.AddPage(form)
	Ok(w, "ok")
}

// ListPage get page list
func ListPage(w http.ResponseWriter, r *http.Request) {
	// 1.parse request params
	query := r.URL.Query().Get("query")
	form, err := defs.DecodePageSearchQuery(query)
	if err != nil {
		Error(w, errcode.BadRequestParam, "bad request")
		return
	}
	// 2.query page from bolt
	list, total, _ := bolt.ListPage(form)
	data := map[string]interface{}{
		"list":  list,
		"total": total,
	}
	Ok(w, data)
}

// DeletePage delete page
func DeletePage(w http.ResponseWriter, r *http.Request) {
	pageName := r.URL.Query().Get("page_name")
	_ = bolt.DeletePage(pageName)
	Ok(w, "ok")
}
