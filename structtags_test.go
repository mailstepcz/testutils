package testutils

import (
	"reflect"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestJSONTagChecker(t *testing.T) {
	req := require.New(t)

	type person struct {
		FullName   string      `json:"fullName"`
		Age        int         `json:"age"`
		JobID      uuid.UUID   `json:"jobId"`
		AddressIDs []uuid.UUID `json:"addressIds"`
	}

	err := CheckJSONTags(reflect.TypeFor[person](), "id")
	req.NoError(err)
}
