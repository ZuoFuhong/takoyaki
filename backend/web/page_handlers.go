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
	// 2.write to bolt
	_ = bolt.AddPage(form)
	Ok(w, "ok")
}

// ListPage get page list
func ListPage(w http.ResponseWriter, r *http.Request) {
	// 1.parse request params
	form := new(defs.PageSearchForm)
	if err := json.NewDecoder(r.Body).Decode(form); err != nil {
		log.Printf("call decoder.Decode failed, err: %v", err)
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
