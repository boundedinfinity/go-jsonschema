package model

import (
	o "github.com/boundedinfinity/go-commoner/optioner"
	"github.com/boundedinfinity/go-jsonschema/schematype"
)

func NewInteger() *JsonSchemaInteger {
	schema := &JsonSchemaInteger{
		Schema: o.Some(SCHEMA_VERSION_2020_12),
		Type:   o.Some(schematype.Integer),
	}

	return schema
}

type JsonSchemaInteger struct {
	Id               o.Option[IdT]                   `json:"$id" yaml:"$id"`
	Type             o.Option[schematype.SchemaType] `json:"type" yaml:"type"`
	Schema           o.Option[SchemaT]               `json:"$schema" yaml:"$schema"`
	Comment          o.Option[CommentT]              `json:"$comment" yaml:"$comment"`
	Deprecated       o.Option[bool]                  `json:"deprecated" yaml:"deprecated"`
	Description      o.Option[DescriptionT]          `json:"description" yaml:"description"`
	Title            o.Option[TitleT]                `json:"title" yaml:"title"`
	ReadOnly         o.Option[bool]                  `json:"readOnly" yaml:"readOnly"`
	WriteOnly        o.Option[bool]                  `json:"writeOnly" yaml:"writeOnly"`
	ExclusiveMaximum o.Option[int64]                 `json:"exclusiveMaximum" yaml:"exclusiveMaximum"`
	ExclusiveMinimum o.Option[int64]                 `json:"exclusiveMinimum" yaml:"exclusiveMinimum"`
	Maximum          o.Option[int64]                 `json:"maximum" yaml:"maximum"`
	Minimum          o.Option[int64]                 `json:"minimum" yaml:"minimum"`
	MultipleOf       o.Option[int64]                 `json:"multipleOf" yaml:"multipleOf"`
}

func (t JsonSchemaInteger) GetId() o.Option[IdT] {
	return t.Id
}

func (t JsonSchemaInteger) GetRef() o.Option[IdT] {
	return o.None[IdT]()
}

func (t JsonSchemaInteger) IsConcrete() bool {
	return true
}

func (t JsonSchemaInteger) IsRef() bool {
	return false
}

func (t JsonSchemaInteger) Validate() error {
	if err := validateMultipleOf(t.MultipleOf); err != nil {
		return err
	}

	if err := validateMaxMin(t.Maximum, t.Minimum); err != nil {
		return err
	}

	return nil
}