package conf

import (
	"fmt"
	"testing"
)

func Test_GetConfig(t *testing.T) {
	config := GetConfig()
	fmt.Printf("%v", config)
}
