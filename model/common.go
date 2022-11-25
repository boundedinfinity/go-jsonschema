package model

import (
	"encoding/json"
	"fmt"

	"github.com/boundedinfinity/go-commoner/caser"
	o "github.com/boundedinfinity/go-commoner/optioner"
	"github.com/boundedinfinity/go-commoner/pather"
	"github.com/boundedinfinity/go-jsonschema/schematype"
)

type IdT string
type TitleT string
type DescriptionT string
type SchemaT string
type CommentT string
type PatternT string
type EnumT string
type EnumDescriptionT string

type JsonSchema interface {
	GetId() o.Option[IdT]
	GetRef() o.Option[IdT]
	IsConcrete() bool
	IsRef() bool
	Validate() error
}

type jsonSchemaDescriminator struct {
	Id         o.Option[IdT]                   `json:"$id" yaml:"$id"`
	Type       o.Option[schematype.SchemaType] `json:"type" yaml:"type"`
	Ref        o.Option[string]                `json:"$ref" yaml:"$ref"`
	Items      json.RawMessage                 `json:"items" yaml:"items"`
	Properties map[string]json.RawMessage      `json:"properties" yaml:"properties"`
}

func mergeDescription(a, b o.Option[DescriptionT]) o.Option[DescriptionT] {
	var d string

	if a.Defined() && b.Defined() {
		d = fmt.Sprintf("%v\n%v", b.Get(), a.Get())
	}

	if b.Defined() {
		d = string(b.Get())
	}

	return o.OfZ(DescriptionT(d))
}

func mergeTitle(a, b o.Option[TitleT], id o.Option[IdT]) o.Option[TitleT] {
	return o.FirstOf(b, a, func() o.Option[TitleT] {
		var n string
		n = pather.Base(id.Get())
		n = caser.KebabToPascal(n)
		return o.OfZ(TitleT(n))
	}())
}
