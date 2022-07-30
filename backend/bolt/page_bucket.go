package bolt

import (
	"encoding/json"
	"errors"
	"log"
	"takoyaki/defs"
	"takoyaki/utils"
)

var NotFoundRecord = errors.New("not found record")

const pageBucket = "page_bucket"

// AddPage create/update page config
func AddPage(form *defs.PageForm) error {
	pageName := form.Name
	formBytes, err := json.Marshal(form)
	if err != nil {
		return err
	}
	if err := writeToDB(pageBucket, pageName, formBytes); err != nil {
		log.Printf("call bolt.WriteToDB failed, err: %v\n", err)
	}
	return err
}

// GetPage query page config
func GetPage(pageName string) (*defs.PageForm, error) {
	bytes, err := readFromDB(pageBucket, pageName)
	if err != nil {
		log.Printf("call bolt.ReadFromDB failed, err: %v\n", err)
		return nil, err
	}
	if len(bytes) == 0 {
		return nil, NotFoundRecord
	}
	form := new(defs.PageForm)
	if err := json.Unmarshal(bytes, form); err != nil {
		return nil, err
	}
	return form, nil
}

// ListPage get page list
func ListPage(form *defs.PageSearchForm) ([]*defs.PageForm, int, error) {
	values, err := readAllFromDB(pageBucket)
	if err != nil {
		return nil, 0, err
	}
	dfList := make([]*defs.PageForm, 0)
	for _, value := range values {
		df := new(defs.PageForm)
		if err := json.Unmarshal(value, form); err != nil {
			continue
		}
		if utils.Contains(df.Name, form.Name) && utils.Contains(df.DataSource, form.DataSource) {
			dfList = append(dfList, df)
		}
	}
	// data paging
	total := len(dfList)
	offset, endpos := utils.PagePos(total, form.Page, form.PageSize)
	return dfList[offset:endpos], total, nil
}

// DeletePage delete a key from the bucket
func DeletePage(pageName string) error {
	return writeToDB(dsBucket, pageName, []byte(""))
}
