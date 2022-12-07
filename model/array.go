package model

import (
	"encoding/json"

	o "github.com/boundedinfinity/go-commoner/optioner"
	"github.com/boundedinfinity/go-jsonschema/schematype"
)

func NewArray() *JsonSchemaArray {
	return &JsonSchemaArray{
		Schema: o.Some(SCHEMA_VERSION_2020_12),
		Type:   o.Some(schematype.Array),
	}
}

var _ = &JsonSchemaArray{}

type JsonSchemaArray struct {
	Id          o.Option[string]                `json:"$id" yaml:"$id"`
	Type        o.Option[schematype.SchemaType] `json:"type" yaml:"type"`
	Schema      o.Option[string]                `json:"$schema" yaml:"$schema"`
	Comment     o.Option[string]                `json:"$comment" yaml:"$comment"`
	Deprecated  o.Option[bool]                  `json:"deprecated" yaml:"deprecated"`
	Description o.Option[string]                `json:"description" yaml:"description"`
	Title       o.Option[string]                `json:"title" yaml:"title"`
	ReadOnly    o.Option[bool]                  `json:"readOnly" yaml:"readOnly"`
	WriteOnly   o.Option[bool]                  `json:"writeOnly" yaml:"writeOnly"`
	Items       o.Option[JsonSchema]            `json:"items" yaml:"items"`
	MinContains o.Option[int64]                 `json:"minContains" yaml:"minContains"`
	MaxContains o.Option[int64]                 `json:"manContains" yaml:"manContains"`
}

func (t JsonSchemaArray) GetId() o.Option[string] {
	return t.Id
}

func (t JsonSchemaArray) GetRef() o.Option[string] {
	return o.None[string]()
}

func (t JsonSchemaArray) IsConcrete() bool {
	return true
}

func (t JsonSchemaArray) IsRef() bool {
	return false
}

func (t JsonSchemaArray) Validate() error {
	if t.Items.Empty() {
		return ErrArrayItemsEmpty
	}

	if err := t.Items.Get().Validate(); err != nil {
		return err
	}

	if err := validateMaxMinContains(t.MaxContains, t.MinContains); err != nil {
		return err
	}

	return nil
}

type jsonSchemaArray struct {
	Id          o.Option[string]                `json:"$id" yaml:"$id"`
	Type        o.Option[schematype.SchemaType] `json:"type" yaml:"type"`
	Schema      o.Option[string]                `json:"$schema" yaml:"$schema"`
	Comment     o.Option[string]                `json:"$comment" yaml:"$comment"`
	Deprecated  o.Option[bool]                  `json:"deprecated" yaml:"deprecated"`
	Description o.Option[string]                `json:"description" yaml:"description"`
	Title       o.Option[string]                `json:"title" yaml:"title"`
	ReadOnly    o.Option[bool]                  `json:"readOnly" yaml:"readOnly"`
	WriteOnly   o.Option[bool]                  `json:"writeOnly" yaml:"writeOnly"`
	Items       json.RawMessage                 `json:"items" yaml:"items"`
	MinContains o.Option[int64]                 `json:"minContains" yaml:"minContains"`
	MaxContains o.Option[int64]                 `json:"manContains" yaml:"manContains"`
}

func (t *JsonSchemaArray) UnmarshalJSON(data []byte) error {
	var r jsonSchemaArray

	if err := json.Unmarshal(data, &r); err != nil {
		return err
	}

	t.Id = r.Id
	t.Type = r.Type
	t.Schema = r.Schema
	t.Comment = r.Comment
	t.Deprecated = r.Deprecated
	t.Description = r.Description
	t.Title = r.Title
	t.ReadOnly = r.ReadOnly
	t.WriteOnly = r.WriteOnly
	t.MaxContains = r.MaxContains
	t.MinContains = r.MinContains

	items, err := UnmarshalSchema(r.Items)

	if err != nil {
		return err
	}

	t.Items = o.Some(items)

	return nil
}
