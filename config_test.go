package config

import (
	"fmt"
	"testing"
)

func TestConstructEntity(t *testing.T) {

	st := NewStringType()
	p := Property{"test"}

	switch v := st.(type) {
	case StringType:
		fmt.Println("string type")
	}
}
