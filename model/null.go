package model

import (
	o "github.com/boundedinfinity/go-commoner/optioner"
	"github.com/boundedinfinity/go-jsonschema/schematype"
)

func NewNull() *JsonSchemaNull {
	schema := &JsonSchemaNull{
		Schema: o.Some(SCHEMA_VERSION_2020_12),
		Type:   o.Some(schematype.Null),
	}

	return schema
}

type JsonSchemaNull struct {
	Id          o.Option[IdT]                   `json:"$id" yaml:"$id"`
	Type        o.Option[schematype.SchemaType] `json:"type" yaml:"type"`
	Schema      o.Option[SchemaT]               `json:"$schema" yaml:"$schema"`
	Comment     o.Option[CommentT]              `json:"$comment" yaml:"$comment"`
	Deprecated  o.Option[bool]                  `json:"deprecated" yaml:"deprecated"`
	Description o.Option[DescriptionT]          `json:"description" yaml:"description"`
	Title       o.Option[TitleT]                `json:"title" yaml:"title"`
	ReadOnly    o.Option[bool]                  `json:"readOnly" yaml:"readOnly"`
	WriteOnly   o.Option[bool]                  `json:"writeOnly" yaml:"writeOnly"`
}

var _ = &JsonSchemaNull{}

func (t JsonSchemaNull) GetId() o.Option[IdT] {
	return t.Id
}

func (t JsonSchemaNull) GetRef() o.Option[IdT] {
	return o.None[IdT]()
}

func (t JsonSchemaNull) IsConcrete() bool {
	return true
}

func (t JsonSchemaNull) IsRef() bool {
	return false
}

func (t JsonSchemaNull) Validate() error {

	return nil
}
