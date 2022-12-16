package model

import (
	"encoding/json"
	"fmt"

	"github.com/boundedinfinity/go-commoner/caser"
	o "github.com/boundedinfinity/go-commoner/optioner"
	"github.com/boundedinfinity/go-commoner/pather"
	"github.com/boundedinfinity/go-jsonschema/schematype"
)

type JsonSchema interface {
	Base() *JsonSchemaBase
	Validate() error
}

type JsonSchemaBase struct {
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

func (t JsonSchemaBase) Base() *JsonSchemaBase {
	return &t
}

type jsonSchemaDescriminator struct {
	Id         o.Option[string]                `json:"$id" yaml:"$id"`
	Type       o.Option[schematype.SchemaType] `json:"type" yaml:"type"`
	Ref        o.Option[string]                `json:"$ref" yaml:"$ref"`
	Items      json.RawMessage                 `json:"items" yaml:"items"`
	Properties map[string]json.RawMessage      `json:"properties" yaml:"properties"`
}

func mergeDescription(a, b o.Option[string]) o.Option[string] {
	var d string

	if a.Defined() && b.Defined() {
		d = fmt.Sprintf("%v\n%v", b.Get(), a.Get())
	}

	if b.Defined() {
		d = string(b.Get())
	}

	return o.OfZ(d)
}

func mergeTitle(a, b o.Option[string], id o.Option[string]) o.Option[string] {
	return o.FirstOf(b, a, func() o.Option[string] {
		var n string
		n = pather.Base(id.Get())
		n = caser.KebabToPascal(n)
		return o.OfZ(n)
	}())
}
