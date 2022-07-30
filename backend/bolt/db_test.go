package bolt

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_KVWriteAndRead(t *testing.T) {
	expect := "mars"
	if err := writeToDB("test", "name", []byte(expect)); err != nil {
		panic(err)
	}
	value, err := readFromDB("test", "name")
	if err != nil {
		panic(err)
	}
	assert.Equal(t, expect, string(value))
}

func Test_ValueNotFound(t *testing.T) {
	value, err := readFromDB("test", "name")
	if err != nil {
		panic(err)
	}
	assert.Equal(t, 0, len(value))
}
