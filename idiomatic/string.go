package idiomatic

import (
	"github.com/boundedinfinity/go-jsonschema/model"
)

func NewString() JsonSchema {
	schema := &JsonSchemaString{
		JsonSchemaCommon: JsonSchemaCommon{
			Schema: model.SCHEMA_VERSION_2020_12,
			Type:   "string",
		},
	}

	return schema
}

type JsonSchemaString struct {
	JsonSchemaCommon
	Format    StringFormat `json:"format,omitempty" yaml:"format,omitempty"`
	MaxLength int64        `json:"maxLength,omitempty" yaml:"maxLength,omitempty"`
	MinLength int64        `json:"minLength,omitempty" yaml:"minLength,omitempty"`
	Pattern   string       `json:"pattern,omitempty" yaml:"pattern,omitempty"`
	Enum      []string     `json:"enum,omitempty" yaml:"enum,omitempty"`
}

var _ JsonSchema = &JsonSchemaString{}

func (t *JsonSchemaString) TypeName() string {
	return "string"
}

func (t *JsonSchemaString) Common() *JsonSchemaCommon {
	return &t.JsonSchemaCommon
}

func (t *JsonSchemaString) Merge(other JsonSchemaString) {
	t.Common().Merge(*other.Common())
	t.Format = mergeSimple(other.Format, t.Format)
	t.MaxLength = mergeSimple(other.MaxLength, t.MaxLength)
	t.MinLength = mergeSimple(other.MinLength, t.MinLength)
	t.Pattern = mergeSimple(other.Pattern, t.Pattern)
	t.Enum = append(other.Enum, t.Enum...)
}

func (t JsonSchemaString) Validate() error {
	if err := t.Common().Validate(); err != nil {
		return nil
	}

	if err := validateMaxMin(t.MaxLength, t.MinLength); err != nil {
		return err
	}

	return nil
}
