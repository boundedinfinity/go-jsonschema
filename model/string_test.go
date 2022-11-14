package model_test

import (
	"encoding/json"
	"testing"

	"github.com/boundedinfinity/go-jsonschema/model"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
)

var schema = "https://json-schema.org/draft/2020-12/schema"
var id = "https://www.boundedinfinity.com/schema/string-1"

func createString() model.JsonSchemaString {
	return model.NewString(id)
}

func Test_String_Unmarshal_Json(t *testing.T) {
	expected := createString()
	var actual model.JsonSchemaString
	input := `{
		"$id": "https://www.boundedinfinity.com/schema/string-1", 
		"type": "string", 
		"$schema":"https://json-schema.org/draft/2020-12/schema"
	}`

	err := json.Unmarshal([]byte(input), &actual)

	assert.Nil(t, err)
	assert.Equal(t, expected.Id.Defined(), actual.Id.Defined())
	assert.Equal(t, expected.Id.Get(), actual.Id.Get())
	assert.Equal(t, expected.Schema.Defined(), actual.Schema.Defined())
	assert.Equal(t, expected.Schema.Get(), actual.Schema.Get())
}

func Test_String_Unmarshal_Yaml(t *testing.T) {
	expected := createString()
	var actual model.JsonSchemaString
	input := `$id: https://www.boundedinfinity.com/schema/string-1
$schema: https://json-schema.org/draft/2020-12/schema
type: string
`

	err := yaml.Unmarshal([]byte(input), &actual)

	assert.Nil(t, err)
	assert.Equal(t, expected.Id.Defined(), actual.Id.Defined())
	assert.Equal(t, expected.Id.Get(), actual.Id.Get())
	assert.Equal(t, expected.Schema.Defined(), actual.Schema.Defined())
	assert.Equal(t, expected.Schema.Get(), actual.Schema.Get())
}

func Test_String_Marshal(t *testing.T) {
	input := createString()

	actual, err := json.Marshal(input)
	expected := `{
		"$id":"https://www.boundedinfinity.com/schema/string-1",
		"$ref":null, 
		"$schema":"https://json-schema.org/draft/2020-12/schema",
		"type":"string",
		"$comment":null,
		"deprecated":null,
		"description":null,
		"title":null,
		"readOnly":null,
		"writeOnly":null,
		"format":null,
		"enum":null,
		"maxLength":null,
		"minLength":null,
		"pattern":null,
		"enum-description":null
	}`

	assert.Nil(t, err)
	assert.JSONEq(t, expected, string(actual))
}
