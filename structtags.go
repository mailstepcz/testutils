package testutils

import (
	"fmt"
	"reflect"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/fealsamh/go-utils/nocopy"
	"github.com/fealsamh/go-utils/testutils"
	"github.com/fealsamh/go-utils/textutils"
)

// CheckJSONTags checks a struct's JSON tags.
// The terms argument specifies abbreviations that ought not to be split, e.g. "id" or "url".
func CheckJSONTags(typ reflect.Type, terms ...string) error {
	for _, field := range testutils.JSONFields(typ) {
		name := nocopy.Bytes(field.Name)
		r, _ := utf8.DecodeRune(name)
		utf8.EncodeRune(name, unicode.ToLower(r))
		comps := textutils.SplitCamelcasedString(nocopy.String(name), terms...)
		var sb strings.Builder
		for i, comp := range comps {
			comp := nocopy.Bytes(strings.ToLower(comp))
			if i > 0 {
				r, _ := utf8.DecodeRune(comp)
				utf8.EncodeRune(comp, unicode.ToUpper(r))
			}
			sb.Write(comp)
		}
		if sb.String() != field.Tag {
			return fmt.Errorf("non-matching field '%s' ('%s' <> '%s')", field.Name, sb.String(), field.Tag)
		}
	}

	return nil
}
