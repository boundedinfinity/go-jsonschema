package model

import (
	"github.com/boundedinfinity/jsonschema/schematype"
	o "github.com/boundedinfinity/optioner"
)

type JsonSchema interface {
	String() string
	GetId() o.Option[string]
	GetType() o.Option[schematype.SchemaType]
	GetRef() o.Option[string]
	Validate() error
}

type JsonSchemaCommon struct {
	Id      o.Option[string]                `json:"$id" yaml:"$id"`
	Schema  o.Option[string]                `json:"$schema" yaml:"$schema"`
	Type    o.Option[schematype.SchemaType] `json:"type" yaml:"type"`
	Ref     o.Option[string]                `json:"$ref" yaml:"$ref"`
	Comment o.Option[string]                `json:"$comment" yaml:"$comment"`
	// Default     o.Option[T]                     `json:"default" yaml:"default"`
	Deprecated  o.Option[bool]   `json:"deprecated" yaml:"deprecated"`
	Description o.Option[string] `json:"description" yaml:"description"`
	// Examples    o.Option[[]T]                   `json:"examples" yaml:"examples"`
	Title     o.Option[string] `json:"title" yaml:"title"`
	ReadOnly  o.Option[bool]   `json:"readOnly" yaml:"readOnly"`
	WriteOnly o.Option[bool]   `json:"writeOnly" yaml:"writeOnly"`
}

func (t JsonSchemaCommon) GetId() o.Option[string] {
	return t.Id
}

func (t JsonSchemaCommon) GetType() o.Option[schematype.SchemaType] {
	return t.Type
}

func (t JsonSchemaCommon) GetRef() o.Option[string] {
	return t.Ref
}