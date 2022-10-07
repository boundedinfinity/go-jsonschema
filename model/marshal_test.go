package model_test

import (
	"testing"

	"github.com/boundedinfinity/go-jsonschema/model"
	"github.com/boundedinfinity/go-jsonschema/schematype"
	"github.com/boundedinfinity/mimetyper/mime_type"
	"github.com/stretchr/testify/assert"
)

var (
	jsonString = `{
		"$id": "https://www.boundedinfinity.com/schema/test/string-plain",
		"$schema": "https://json-schema.org/draft/2020-12/schema",
		"type": "string"
	}`
	jsonInteger = `{
		"$id": "https://www.boundedinfinity.com/schema/test/string-plain",
		"$schema": "https://json-schema.org/draft/2020-12/schema",
		"type": "integer"
	}`
	json_String_Integer = "[" + jsonString + "," + jsonInteger + "]"
)

func Test_Unmarshal_JsonString(t *testing.T) {
	input := jsonString

	var schemas []model.JsonSchema

	err := model.Unmarshal([]byte(input), &schemas, mime_type.ApplicationJson)

	assert.Nil(t, err)
	assert.True(t, schemas[0].GetType().Defined())
	assert.Equal(t, schematype.String, schemas[0].GetType().Get())
}

func Test_Unmarshal_JsonInteger(t *testing.T) {
	input := jsonInteger

	var schemas []model.JsonSchema

	err := model.Unmarshal([]byte(input), &schemas, mime_type.ApplicationJson)

	assert.Nil(t, err)
	assert.Equal(t, 1, len(schemas))
	assert.True(t, schemas[0].GetType().Defined())
}

func Test_Unmarshal_Mixed(t *testing.T) {
	input := json_String_Integer

	var schemas []model.JsonSchema

	err := model.Unmarshal([]byte(input), &schemas, mime_type.ApplicationJson)

	assert.Nil(t, err)
	assert.Equal(t, 2, len(schemas))
	assert.True(t, schemas[0].GetType().Defined())
}
