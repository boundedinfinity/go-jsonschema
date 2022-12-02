package model

import (
	"github.com/boundedinfinity/go-commoner/mapper"
	o "github.com/boundedinfinity/go-commoner/optioner"
	"github.com/boundedinfinity/go-jsonschema/schematype"
	"github.com/boundedinfinity/go-jsonschema/stringformat"
)

func NewString() JsonSchema {
	schema := &JsonSchemaString{
		Schema: o.Some(SCHEMA_VERSION_2020_12),
		Type:   o.Some(schematype.String),
	}

	return schema
}

type JsonSchemaString struct {
	Id              o.Option[IdT]                                    `json:"$id" yaml:"$id"`
	Type            o.Option[schematype.SchemaType]                  `json:"type" yaml:"type"`
	Schema          o.Option[SchemaT]                                `json:"$schema" yaml:"$schema"`
	Comment         o.Option[CommentT]                               `json:"$comment" yaml:"$comment"`
	Deprecated      o.Option[bool]                                   `json:"deprecated" yaml:"deprecated"`
	Description     o.Option[DescriptionT]                           `json:"description" yaml:"description"`
	Title           o.Option[TitleT]                                 `json:"title" yaml:"title"`
	ReadOnly        o.Option[bool]                                   `json:"readOnly" yaml:"readOnly"`
	WriteOnly       o.Option[bool]                                   `json:"writeOnly" yaml:"writeOnly"`
	Format          o.Option[stringformat.StringFormat]              `json:"format,omitempty" yaml:"format,omitempty"`
	MaxLength       o.Option[int]                                    `json:"maxLength,omitempty" yaml:"maxLength,omitempty"`
	MinLength       o.Option[int]                                    `json:"minLength,omitempty" yaml:"minLength,omitempty"`
	Pattern         o.Option[PatternT]                               `json:"pattern,omitempty" yaml:"pattern,omitempty"`
	Enum            o.Option[[]EnumT]                                `json:"enum,omitempty" yaml:"enum,omitempty"`
	EnumDescription o.Option[mapper.Mapper[EnumT, EnumDescriptionT]] `json:"enum-description,omitempty" yaml:"enum-description,omitempty"`
}

var _ = &JsonSchemaString{}

func (t JsonSchemaString) GetId() o.Option[IdT] {
	return t.Id
}

func (t JsonSchemaString) GetRef() o.Option[IdT] {
	return o.None[IdT]()
}

func (t JsonSchemaString) IsConcrete() bool {
	return true
}

func (t JsonSchemaString) IsRef() bool {
	return false
}

func (t JsonSchemaString) Validate() error {
	if err := validateMaxMinLength(t.MaxLength, t.MinLength); err != nil {
		return err
	}

	return nil
}
