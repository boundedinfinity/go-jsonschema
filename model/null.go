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
	Id          o.Option[string]                `json:"$id" yaml:"$id"`
	Type        o.Option[schematype.SchemaType] `json:"type" yaml:"type"`
	Schema      o.Option[string]                `json:"$schema" yaml:"$schema"`
	Comment     o.Option[string]                `json:"$comment" yaml:"$comment"`
	Deprecated  o.Option[bool]                  `json:"deprecated" yaml:"deprecated"`
	Description o.Option[string]                `json:"description" yaml:"description"`
	Title       o.Option[string]                `json:"title" yaml:"title"`
	ReadOnly    o.Option[bool]                  `json:"readOnly" yaml:"readOnly"`
	WriteOnly   o.Option[bool]                  `json:"writeOnly" yaml:"writeOnly"`
}

var _ = &JsonSchemaNull{}

func (t JsonSchemaNull) GetId() o.Option[string] {
	return t.Id
}

func (t JsonSchemaNull) GetRef() o.Option[string] {
	return o.None[string]()
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
