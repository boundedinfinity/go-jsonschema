package model

import (
	o "github.com/boundedinfinity/go-commoner/optioner"
	"github.com/boundedinfinity/go-jsonschema/schematype"
	"github.com/boundedinfinity/go-jsonschema/stringformat"
)

type JsonSchemaString struct {
	JsonSchemaCommon
	Format          o.Option[stringformat.StringFormat] `json:"format,omitempty" yaml:"format,omitempty"`
	MaxLength       o.Option[int]                       `json:"maxLength,omitempty" yaml:"maxLength,omitempty"`
	MinLength       o.Option[int]                       `json:"minLength,omitempty" yaml:"minLength,omitempty"`
	Pattern         o.Option[string]                    `json:"pattern,omitempty" yaml:"pattern,omitempty"`
	Enum            o.Option[[]string]                  `json:"enum,omitempty" yaml:"enum,omitempty"`
	EnumDescription o.Option[map[string]string]         `json:"enum-description,omitempty" yaml:"enum-description,omitempty"`
}

func NewString[T ~string](id string) JsonSchemaString {
	return JsonSchemaString{
		JsonSchemaCommon: JsonSchemaCommon{
			Schema: o.Some(SCHEMA_VERSION_2020_12),
			Type:   o.Some(schematype.String),
			Id:     o.Some(id),
		},
	}
}
