package model

import (
	"github.com/boundedinfinity/go-commoner/mapper"
	o "github.com/boundedinfinity/go-commoner/optioner"
	"github.com/boundedinfinity/go-jsonschema/schematype"
	"github.com/boundedinfinity/go-jsonschema/stringformat"
)

type JsonSchemaString struct {
	JsonSchemaCommon `yaml:",inline"`
	Format           o.Option[stringformat.StringFormat]     `json:"format,omitempty" yaml:"format,omitempty"`
	MaxLength        o.Option[int]                           `json:"maxLength,omitempty" yaml:"maxLength,omitempty"`
	MinLength        o.Option[int]                           `json:"minLength,omitempty" yaml:"minLength,omitempty"`
	Pattern          o.Option[string]                        `json:"pattern,omitempty" yaml:"pattern,omitempty"`
	Enum             o.Option[[]string]                      `json:"enum,omitempty" yaml:"enum,omitempty"`
	EnumDescription  o.Option[mapper.Mapper[string, string]] `json:"enum-description,omitempty" yaml:"enum-description,omitempty"`
}

func NewString(id string) JsonSchemaString {
	return JsonSchemaString{
		JsonSchemaCommon: JsonSchemaCommon{
			Schema: o.Some(SCHEMA_VERSION_2020_12),
			Type:   o.Some(schematype.String),
			Id:     o.Some(id),
		},
	}
}

func (t JsonSchemaString) Copy() JsonSchema {
	return JsonSchemaString{
		JsonSchemaCommon: t.JsonSchemaCommon.Copy().(JsonSchemaCommon),
		Format:           t.Format,
		MaxLength:        t.MaxLength,
		MinLength:        t.MinLength,
		Pattern:          t.Pattern,
		Enum:             t.Enum,
		EnumDescription:  t.EnumDescription,
	}
}

func (t JsonSchemaString) Merge(other JsonSchema) JsonSchema {
	internal := other.(JsonSchemaString)

	return JsonSchemaString{
		JsonSchemaCommon: t.JsonSchemaCommon.Merge(internal.JsonSchemaCommon).(JsonSchemaCommon),
		Format:           internal.Format.OrFirst(t.Format),
		MaxLength:        internal.MaxLength.OrFirst(t.MaxLength),
		MinLength:        internal.MinLength.OrFirst(t.MinLength),
		Pattern:          internal.Pattern.OrFirst(t.Pattern),
		Enum:             internal.Enum.OrFirst(t.Enum),
		EnumDescription:  internal.EnumDescription.OrFirst(t.EnumDescription),
	}
}
