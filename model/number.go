package model

import (
	o "github.com/boundedinfinity/go-commoner/optioner"
	"github.com/boundedinfinity/go-jsonschema/schematype"
)

func NewNumber() *JsonSchemaNumber {
	schema := &JsonSchemaNumber{
		Schema: o.Some(SCHEMA_VERSION_2020_12),
		Type:   o.Some(schematype.Number),
	}

	return schema
}

type JsonSchemaNumber struct {
	Id               o.Option[string]                `json:"$id" yaml:"$id"`
	Type             o.Option[schematype.SchemaType] `json:"type" yaml:"type"`
	Schema           o.Option[string]                `json:"$schema" yaml:"$schema"`
	Comment          o.Option[string]                `json:"$comment" yaml:"$comment"`
	Deprecated       o.Option[bool]                  `json:"deprecated" yaml:"deprecated"`
	Description      o.Option[string]                `json:"description" yaml:"description"`
	Title            o.Option[string]                `json:"title" yaml:"title"`
	ReadOnly         o.Option[bool]                  `json:"readOnly" yaml:"readOnly"`
	WriteOnly        o.Option[bool]                  `json:"writeOnly" yaml:"writeOnly"`
	ExclusiveMaximum o.Option[float64]               `json:"exclusiveMaximum" yaml:"exclusiveMaximum"`
	ExclusiveMinimum o.Option[float64]               `json:"exclusiveMinimum" yaml:"exclusiveMinimum"`
	Maximum          o.Option[float64]               `json:"maximum" yaml:"maximum"`
	Minimum          o.Option[float64]               `json:"minimum" yaml:"minimum"`
	MultipleOf       o.Option[float64]               `json:"multipleOf" yaml:"multipleOf"`
}

var _ = &JsonSchemaNumber{}

func (t JsonSchemaNumber) GetId() o.Option[string] {
	return t.Id
}

func (t JsonSchemaNumber) GetRef() o.Option[string] {
	return o.None[string]()
}

func (t JsonSchemaNumber) IsConcrete() bool {
	return true
}

func (t JsonSchemaNumber) IsRef() bool {
	return false
}

func (t JsonSchemaNumber) Validate() error {
	if err := validateMultipleOf(t.MultipleOf); err != nil {
		return err
	}

	if err := validateMaxMin(t.Maximum, t.Minimum); err != nil {
		return err
	}

	return nil
}
