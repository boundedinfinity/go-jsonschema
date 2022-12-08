package model

import (
	o "github.com/boundedinfinity/go-commoner/optioner"
	"github.com/boundedinfinity/go-jsonschema/schematype"
)

func NewNull() *JsonSchemaNull {
	schema := &JsonSchemaNull{
		JsonSchemaBase: JsonSchemaBase{
			Schema: o.Some(SCHEMA_VERSION_2020_12),
			Type:   o.Some(schematype.Null),
		},
	}

	return schema
}

type JsonSchemaNull struct {
	JsonSchemaBase
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
