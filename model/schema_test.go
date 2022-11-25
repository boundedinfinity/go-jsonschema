package model_test

import (
	"testing"

	o "github.com/boundedinfinity/go-commoner/optioner"
	"github.com/boundedinfinity/go-jsonschema/model"
	"github.com/boundedinfinity/go-jsonschema/schematype"
	"github.com/stretchr/testify/assert"
)

func Test_Unmarshal_String(t *testing.T) {
	input := `{
				"$id": "some-id", 
				"type": "string", 
				"$schema":"https://json-schema.org/draft/2020-12/schema"
			}`
	expected := &model.JsonSchemaString{
		Id:   o.Some(model.IdT("some-id")),
		Type: o.Some(schematype.String),
	}

	actual, err := model.UnmarshalSchema([]byte(input))

	assert.Nil(t, err)
	assert.IsType(t, expected, actual)

	actual2 := actual.(*model.JsonSchemaString)

	assert.Equal(t, expected.Id.Defined(), actual2.Id.Defined())
	assert.Equal(t, expected.Id.Get(), actual2.Id.Get())
	assert.Equal(t, expected.Type.Defined(), actual2.Type.Defined())
	assert.Equal(t, expected.Type.Get(), actual2.Type.Get())
}

func Test_Unmarshal_Array(t *testing.T) {
	input := `{
				"$id": "some-id", 
				"type": "array", 
				"$schema":"https://json-schema.org/draft/2020-12/schema",
				"items": {
					"type": "string"
				}
			}`

	expected := model.NewArray()
	expected.Id = o.Some(model.IdT("some-id"))
	expected.Items = o.Some(model.NewString())

	actual, err := model.UnmarshalSchema([]byte(input))

	assert.Nil(t, err)
	assert.IsType(t, expected, actual)

	assert.Equal(t, expected.Id.Defined(), actual.(*model.JsonSchemaArray).Id.Defined())
	assert.Equal(t, expected.Id.Get(), actual.(*model.JsonSchemaArray).Id.Get())
	assert.Equal(t, expected.Type.Defined(), actual.(*model.JsonSchemaArray).Type.Defined())
	assert.Equal(t, expected.Type.Get(), actual.(*model.JsonSchemaArray).Type.Get())

	assert.Equal(t, expected.Items.Defined(), actual.(*model.JsonSchemaArray).Items.Defined())
	assert.IsType(t, expected.Items.Get(), actual.(*model.JsonSchemaArray).Items.Get())
}

func Test_Unmarshal_Object(t *testing.T) {
	input := `{
				"$id": "some-id", 
				"type": "object", 
				"$schema":"https://json-schema.org/draft/2020-12/schema",
				"properties": {
					"a": { "type": "string" },
					"b": { "type": "integer" }
				}
			}`

	expected := model.NewObject()
	expected.Id = o.Some(model.IdT("some-id"))
	expected.Properties["a"] = model.NewString()
	expected.Properties["b"] = model.NewInteger()

	actual, err := model.UnmarshalSchema([]byte(input))

	assert.Nil(t, err)
	assert.IsType(t, expected, actual)
	assert.Equal(t, expected.Id.Defined(), actual.(*model.JsonSchemaObject).Id.Defined())
	assert.Equal(t, expected.Id.Get(), actual.(*model.JsonSchemaObject).Id.Get())
	assert.Equal(t, expected.Type.Defined(), actual.(*model.JsonSchemaObject).Type.Defined())
	assert.Equal(t, expected.Type.Get(), actual.(*model.JsonSchemaObject).Type.Get())
	assert.True(t, actual.(*model.JsonSchemaObject).Properties.Has("a"))
	assert.True(t, actual.(*model.JsonSchemaObject).Properties.Has("b"))
}

func Test_Unmarshal_Ref(t *testing.T) {
	input := `{
				"$schema":"https://json-schema.org/draft/2020-12/schema",
				"$ref": "some-ref"
			}`

	expected := model.NewRef("some-ref")

	actual, err := model.UnmarshalSchema([]byte(input))

	assert.Nil(t, err)
	assert.IsType(t, expected, actual)
	assert.Equal(t, expected.Id.Defined(), actual.(*model.JsonSchemaRef).Id.Defined())
	assert.Equal(t, expected.Id.Get(), actual.(*model.JsonSchemaRef).Id.Get())
	assert.Equal(t, expected.Ref.Defined(), actual.(*model.JsonSchemaRef).Ref.Defined())
	assert.Equal(t, expected.Ref.Get(), actual.(*model.JsonSchemaRef).Ref.Get())
}
