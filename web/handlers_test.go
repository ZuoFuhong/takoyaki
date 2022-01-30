package web

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

func Test_Map(t *testing.T) {
	tmp := make(map[string]interface{})
	tmp["name"] = "mars"
	tmp["age"] = 25
	tmp["date"] = time.Now()
	for k, v := range tmp {
		if k == "name" {
			tmp["name"] = "dfasdfas"
		}
		fmt.Printf("k = %s, v = %v\n", k, v)
	}

	fmt.Printf("%v\n", tmp)
	switch tmp["date"].(type) {
	case time.Time:
		fmt.Printf("value: %v\n", tmp["date"])
	}

	num := 1.0
	fmt.Printf("num type = %s\n", reflect.TypeOf(num).Name())
	fmt.Printf("type = %v\n", reflect.TypeOf(tmp["date"]).Name())
}
