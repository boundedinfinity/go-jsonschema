package json_schema_test

import (
	"encoding/json"
	"testing"

	"github.com/boundedinfinity/go-jsonschema/idiomatic/json_schema"
	"github.com/stretchr/testify/assert"
)

var (
	case1 = &json_schema.JsonSchemaNumber{
		JsonSchemaCore: json_schema.JsonSchemaCore{
			Id: "some-id",
		},
	}
)

func Test_Entity_Marshal_Json(t *testing.T) {
	tcs := []struct {
		name     string
		input    json_schema.JsonSchema
		expected string
		err      error
	}{
		{
			name:  "case 1",
			input: case1,
			err:   nil,
			expected: `{
				"type": "number",
				"$id": "some-id"
			}`,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(tt *testing.T) {
			bs, err := json.MarshalIndent(tc.input, "", "   ")
			actual := string(bs)

			assert.ErrorIs(tt, err, tc.err, actual)
			assert.JSONEqf(tt, tc.expected, actual, actual)
		})
	}
}
