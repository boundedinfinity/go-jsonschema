package model

import (
	o "github.com/boundedinfinity/go-commoner/optioner"
	"github.com/boundedinfinity/go-jsonschema/schematype"
)

func NewBoolean() *JsonSchemaBoolean {
	schema := &JsonSchemaBoolean{
		Schema: o.Some(SCHEMA_VERSION_2020_12),
		Type:   o.Some(schematype.Boolean),
	}

	return schema
}

type JsonSchemaBoolean struct {
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

var _ = &JsonSchemaBoolean{}

func (t JsonSchemaBoolean) GetId() o.Option[IdT] {
	return t.Id
}

func (t JsonSchemaBoolean) GetRef() o.Option[IdT] {
	return o.None[IdT]()
}

func (t JsonSchemaBoolean) IsConcrete() bool {
	return true
}

func (t JsonSchemaBoolean) IsRef() bool {
	return false
}

func (t JsonSchemaBoolean) Validate() error {

	return nil
}
