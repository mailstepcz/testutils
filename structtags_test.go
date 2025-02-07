package testutils

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestJSONTagChecker(t *testing.T) {
	req := require.New(t)

	type person struct {
		FullName string `json:"fullName"`
		Age      int    `json:"age"`
	}

	err := CheckJSONTags(reflect.TypeFor[person]())
	req.NoError(err)
}
