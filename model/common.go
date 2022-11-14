package model

import (
	"fmt"

	o "github.com/boundedinfinity/go-commoner/optioner"
	"github.com/boundedinfinity/go-jsonschema/schematype"
)

type JsonSchema interface {
	Copy() JsonSchema
	Merge(JsonSchema) JsonSchema
	GetType() o.Option[schematype.SchemaType]
	GetRef() o.Option[string]
	GetId() o.Option[string]
}

type JsonSchemaCommon struct {
	Id          o.Option[string]                `json:"$id" yaml:"$id"`
	Type        o.Option[schematype.SchemaType] `json:"type" yaml:"type"`
	Ref         o.Option[string]                `json:"$ref" yaml:"$ref"`
	Schema      o.Option[string]                `json:"$schema" yaml:"$schema"`
	Comment     o.Option[string]                `json:"$comment" yaml:"$comment"`
	Deprecated  o.Option[bool]                  `json:"deprecated" yaml:"deprecated"`
	Description o.Option[string]                `json:"description" yaml:"description"`
	Title       o.Option[string]                `json:"title" yaml:"title"`
	ReadOnly    o.Option[bool]                  `json:"readOnly" yaml:"readOnly"`
	WriteOnly   o.Option[bool]                  `json:"writeOnly" yaml:"writeOnly"`
}

func (t JsonSchemaCommon) GetType() o.Option[schematype.SchemaType] {
	return t.Type
}

func (t JsonSchemaCommon) GetRef() o.Option[string] {
	return t.Ref
}

func (t JsonSchemaCommon) GetId() o.Option[string] {
	return t.Id
}

func (t JsonSchemaCommon) Copy() JsonSchema {
	return JsonSchemaCommon{
		Id:          t.Id,
		Type:        t.Type,
		Ref:         t.Ref,
		Schema:      t.Schema,
		Comment:     t.Comment,
		Deprecated:  t.Deprecated,
		Description: t.Description,
		Title:       t.Title,
		ReadOnly:    t.ReadOnly,
		WriteOnly:   t.WriteOnly,
	}
}

func (t JsonSchemaCommon) Merge(other JsonSchema) JsonSchema {
	copy := t.Copy().(JsonSchemaCommon)
	internal := other.Copy().(JsonSchemaCommon)

	if internal.Title.Defined() {
		copy.Title = internal.Title
	}

	if internal.Description.Defined() {
		if t.Description.Defined() {
			copy.Description = o.Some(fmt.Sprintf("%v\n%v",
				internal.Description.Get(),
				t.Description.Get(),
			))
		} else {
			copy.Description = internal.Description
		}
	}

	return copy
}
