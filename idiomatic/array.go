package idiomatic

import (
	"github.com/boundedinfinity/go-jsonschema/model"
)

func NewArray() *JsonSchemaArray {
	return &JsonSchemaArray{
		JsonSchemaCommon: JsonSchemaCommon{
			Schema: model.SCHEMA_VERSION_2020_12,
			Type:   "array",
		},
	}
}

var _ JsonSchema = &JsonSchemaArray{}

type JsonSchemaArray struct {
	JsonSchemaCommon
	Items       JsonSchema `json:"items" yaml:"items"`
	MinContains int64      `json:"minContains" yaml:"minContains"`
	MaxContains int64      `json:"manContains" yaml:"manContains"`
}

func (t JsonSchemaArray) TypeName() string {
	return "boolean"
}

func (t *JsonSchemaArray) Common() *JsonSchemaCommon {
	return &t.JsonSchemaCommon
}

func (t JsonSchemaArray) Validate() error {
	if err := t.Common().Validate(); err != nil {
		return nil
	}

	if t.Items == nil {
		return model.ErrArrayItemsEmpty
	}

	if err := t.Items.Validate(); err != nil {
		return err
	}

	if err := validateMaxMinContains(t.MaxContains, t.MinContains); err != nil {
		return err
	}

	return nil
}
