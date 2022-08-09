package model

import (
	"github.com/boundedinfinity/jsonschema/schematype"
	"github.com/boundedinfinity/jsonschema/stringformat"
	o "github.com/boundedinfinity/optioner"
)

type JsonSchemaString[T ~string] struct {
	JsonSchemaCommon
	Format          o.Option[stringformat.StringFormat] `json:"format,omitempty" yaml:"format,omitempty"`
	Enum            o.Option[[]T]                       `json:"enum,omitempty" yaml:"enum,omitempty"`
	MaxLength       o.Option[int]                       `json:"maxLength,omitempty" yaml:"maxLength,omitempty"`
	MinLength       o.Option[int]                       `json:"minLength,omitempty" yaml:"minLength,omitempty"`
	Pattern         o.Option[T]                         `json:"pattern,omitempty" yaml:"pattern,omitempty"`
	EnumDescription o.Option[map[T]string]              `json:"enum-description,omitempty" yaml:"enum-description,omitempty"`
}

func NewString[T ~string](id string) JsonSchemaString[T] {
	return JsonSchemaString[T]{
		JsonSchemaCommon: JsonSchemaCommon{
			Schema: o.Some(SCHEMA_VERSION_2020_12),
			Type:   o.Some(schematype.String),
			Id:     o.Some(id),
		},
	}
}
