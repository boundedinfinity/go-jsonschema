package model

import (
	o "github.com/boundedinfinity/go-commoner/optioner"
	"github.com/boundedinfinity/go-jsonschema/schematype"
)

func NewBoolean() *JsonSchemaBoolean {
	schema := &JsonSchemaBoolean{
		JsonSchemaBase: JsonSchemaBase{
			Schema: o.Some(SCHEMA_VERSION_2020_12),
			Type:   o.Some(schematype.Boolean),
		},
	}

	return schema
}

type JsonSchemaBoolean struct {
	JsonSchemaBase
}

var _ = &JsonSchemaBoolean{}

func (t JsonSchemaBoolean) GetId() o.Option[string] {
	return t.Id
}

func (t JsonSchemaBoolean) GetRef() o.Option[string] {
	return o.None[string]()
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
