package testutils

import (
	"reflect"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestJSONTagChecker(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		req := require.New(t)

		type person struct {
			FullName   string      `json:"fullName"`
			Age        int         `json:"age"`
			JobID      uuid.UUID   `json:"jobId"`
			AddressIDs []uuid.UUID `json:"addressIds"`
		}

		err := CheckJSONTags(reflect.TypeFor[person](), "id")
		req.NoError(err)
	})

	t.Run("failure", func(t *testing.T) {
		req := require.New(t)

		type person struct {
			FullName   string      `json:"fullname"`
			Age        int         `json:"age"`
			JobID      uuid.UUID   `json:"jobId"`
			AddressIDs []uuid.UUID `json:"addressIds"`
		}

		err := CheckJSONTags(reflect.TypeFor[person](), "id")
		req.NotNil(err)
		req.Contains("non-matching field 'FullName' ('fullName' <> 'fullname')", err.Error())
	})

	t.Run("failure in abbreviation", func(t *testing.T) {
		req := require.New(t)

		type person struct {
			FullName   string      `json:"fullName"`
			Age        int         `json:"age"`
			JobID      uuid.UUID   `json:"jobId"`
			AddressIDs []uuid.UUID `json:"addressIds"`
		}

		err := CheckJSONTags(reflect.TypeFor[person]())
		req.NotNil(err)
		req.Contains("non-matching field 'AddressIDs' ('addressIDs' <> 'addressIds')", err.Error())
	})
}
