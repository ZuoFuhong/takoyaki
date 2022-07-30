package bolt

import (
	"encoding/json"
	"log"
	"takoyaki/defs"
	"takoyaki/utils"
)

const dsBucket = "datasource_bucket"

// AddDataSource create/update data source config
func AddDataSource(form *defs.DataSourceForm) error {
	source := form.Name
	formBytes, err := json.Marshal(form)
	if err != nil {
		return err
	}
	if err := writeToDB(dsBucket, source, formBytes); err != nil {
		log.Printf("call bolt.WriteToDB failed, err: %v\n", err)
	}
	return err
}

// GetDataSource query data source config
func GetDataSource(source string) (*defs.DataSourceForm, error) {
	bytes, err := readFromDB(dsBucket, source)
	if err != nil {
		log.Printf("call bolt.ReadFromDB failed, err: %v\n", err)
		return nil, err
	}
	if len(bytes) == 0 {
		return nil, NotFoundRecord
	}
	form := new(defs.DataSourceForm)
	if err := json.Unmarshal(bytes, form); err != nil {
		return nil, err
	}
	return form, nil
}

// ListDataSource get data source list
func ListDataSource(form *defs.DataSourceSearchForm) ([]*defs.DataSourceForm, int, error) {
	values, err := readAllFromDB(dsBucket)
	if err != nil {
		return nil, 0, err
	}
	dfList := make([]*defs.DataSourceForm, 0)
	for _, value := range values {
		df := new(defs.DataSourceForm)
		if err := json.Unmarshal(value, form); err != nil {
			continue
		}
		// filter by keyword
		if utils.Contains(df.Name, form.Name) && utils.Contains(df.Host, form.Host) {
			dfList = append(dfList, df)
		}
	}
	// slice paging
	total := len(dfList)
	offset, endpos := utils.PagePos(total, form.Page, form.PageSize)
	return dfList[offset:endpos], total, nil
}
