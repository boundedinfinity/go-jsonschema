package model_test

import (
	"encoding/json"
	"testing"

	"github.com/boundedinfinity/jsonschema/model"
	"github.com/boundedinfinity/jsonschema/schematype"
	o "github.com/boundedinfinity/optioner"
	"github.com/stretchr/testify/assert"
)

var schema = "https://json-schema.org/draft/2020-12/schema"
var id = "https://www.boundedinfinity.com/schema/string-1"

func createString() model.JsonSchemaString[string] {
	return model.JsonSchemaString[string]{
		JsonSchemaGeneric: model.JsonSchemaGeneric[string]{
			Id:     o.Some(id),
			Schema: o.Some(schema),
			Type:   o.Some(schematype.String),
		},
	}
}

func Test_String(t *testing.T) {
	actual := createString()
	assert.True(t, actual.Id.Defined())
}

func Test_String_Marshal(t *testing.T) {
	input := createString()

	actual, err := json.Marshal(input)
	expected := `{"$id":"https://www.boundedinfinity.com/schema/string-1","$ref":null, "$schema":"https://json-schema.org/draft/2020-12/schema","type":"string","$comment":null,"default":null,"deprecated":null,"description":null,"examples":null,"title":null,"readOnly":null,"writeOnly":null,"format":null,"enum":null,"maxLength":null,"minLength":null,"pattern":null,"enum-description":null}`

	assert.Nil(t, err)
	assert.JSONEq(t, expected, string(actual))
}

func Test_String_Unmarshal(t *testing.T) {
	expected := createString()
	var actual model.JsonSchemaString[string]
	input := `{"$id": "https://www.boundedinfinity.com/schema/string-1", "type": "string", "$schema":"https://json-schema.org/draft/2020-12/schema"}`

	err := json.Unmarshal([]byte(input), &actual)

	assert.Nil(t, err)
	assert.Equal(t, expected.Id.Defined(), actual.Id.Defined())
	assert.Equal(t, expected.Id.Get(), actual.Id.Get())
	assert.Equal(t, expected.Schema.Defined(), actual.Schema.Defined())
	assert.Equal(t, expected.Schema.Get(), actual.Schema.Get())
}
