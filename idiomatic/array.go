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
	MinContains int        `json:"minContains" yaml:"minContains"`
	MaxContains int        `json:"manContains" yaml:"manContains"`
}

func (t JsonSchemaArray) TypeName() string {
	return "boolean"
}

func (t *JsonSchemaArray) Common() *JsonSchemaCommon {
	return &t.JsonSchemaCommon
}

func (t JsonSchemaArray) Copy() JsonSchema {
	return &JsonSchemaArray{
		JsonSchemaCommon: t.Common().Copy(),
		Items:            t.Copy(),
		MinContains:      t.MinContains,
		MaxContains:      t.MaxContains,
	}
}

func (t JsonSchemaArray) Validate() error {
	if err := t.Common().Validate(); err != nil {
		return nil
	}

	if t.Items == nil {
		return model.ErrJsonSchemaArrayItemsEmpty
	}

	if err := t.Items.Validate(); err != nil {
		return err
	}

	if err := validateMaxMinContains(t.MaxContains, t.MinContains); err != nil {
		return err
	}

	return nil
}
