package web

import (
	"encoding/json"
	"log"
	"net/http"
	"takoyaki/bolt"
	"takoyaki/defs"
	"takoyaki/errcode"
)

// AddDataSource add data source
func AddDataSource(w http.ResponseWriter, r *http.Request) {
	// 1.parse request params
	form := new(defs.DataSourceForm)
	if err := json.NewDecoder(r.Body).Decode(form); err != nil {
		log.Printf("call decoder.Decode failed, err: %v", err)
		Error(w, errcode.BadRequestParam, "bad request")
		return
	}
	// 2.write to bolt
	_ = bolt.AddDataSource(form)
	Ok(w, "ok")
}

// ListDataSource get data source list
func ListDataSource(w http.ResponseWriter, r *http.Request) {
	// 1.parse request params
	form := new(defs.DataSourceSearchForm)
	if err := json.NewDecoder(r.Body).Decode(form); err != nil {
		log.Printf("call decoder.Decode failed, err: %v", err)
		Error(w, errcode.BadRequestParam, "bad request")
		return
	}
	// 2.query source from bolt
	list, total, _ := bolt.ListDataSource(form)
	data := map[string]interface{}{
		"list":  list,
		"total": total,
	}
	Ok(w, data)
}

// DeleteDataSource delete data source
func DeleteDataSource(w http.ResponseWriter, r *http.Request) {
	source := r.URL.Query().Get("source")
	// 1.check page usage
	form := &defs.PageSearchForm{DataSource: source}
	_, total, err := bolt.ListPage(form)
	if err != nil {
		log.Printf("call bolt.ListPage failed, err: %v\n", err)
		Error(w, errcode.InternalServerError, "System busy")
		return
	}
	if total > 0 {
		Error(w, errcode.NotAllowedToDelete, "Not allowed to delete")
		return
	}
	// 2.write to bolt
	_ = bolt.DeleteDataSource(source)
	Ok(w, "ok")
}
