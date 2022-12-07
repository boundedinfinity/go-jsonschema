package model

import (
	"encoding/json"

	"github.com/boundedinfinity/go-commoner/mapper"
	o "github.com/boundedinfinity/go-commoner/optioner"
	"github.com/boundedinfinity/go-jsonschema/schematype"
)

func NewObject() *JsonSchemaObject {
	schema := &JsonSchemaObject{
		Schema:     o.Some(SCHEMA_VERSION_2020_12),
		Type:       o.Some(schematype.Object),
		Properties: mapper.Mapper[string, JsonSchema]{},
	}

	return schema
}

type JsonSchemaObject struct {
	Id            o.Option[string]                  `json:"$id" yaml:"$id"`
	Type          o.Option[schematype.SchemaType]   `json:"type" yaml:"type"`
	Schema        o.Option[string]                  `json:"$schema" yaml:"$schema"`
	Comment       o.Option[string]                  `json:"$comment" yaml:"$comment"`
	Deprecated    o.Option[bool]                    `json:"deprecated" yaml:"deprecated"`
	Description   o.Option[string]                  `json:"description" yaml:"description"`
	Title         o.Option[string]                  `json:"title" yaml:"title"`
	ReadOnly      o.Option[bool]                    `json:"readOnly" yaml:"readOnly"`
	WriteOnly     o.Option[bool]                    `json:"writeOnly" yaml:"writeOnly"`
	Properties    mapper.Mapper[string, JsonSchema] `json:"properties" yaml:"properties"`
	MaxProperties o.Option[int64]                   `json:"maxProperties" yaml:"maxProperties"`
	MinProperties o.Option[int64]                   `json:"minProperties" yaml:"minProperties"`
}

var _ = &JsonSchemaObject{}

func (t JsonSchemaObject) GetId() o.Option[string] {
	return t.Id
}

func (t JsonSchemaObject) GetRef() o.Option[string] {
	return o.None[string]()
}

func (t JsonSchemaObject) IsConcrete() bool {
	return true
}

func (t JsonSchemaObject) IsRef() bool {
	return false
}

func (t JsonSchemaObject) Validate() error {
	if err := validateMaxMinProperties(t.MaxProperties, t.MinProperties); err != nil {
		return err
	}

	if t.Properties == nil {
		return ErrObjectPropertiesMissing
	}

	for _, property := range t.Properties {
		if err := property.Validate(); err != nil {
			return err
		}
	}

	return nil
}

type jsonSchemaObject struct {
	Id            o.Option[string]                       `json:"$id" yaml:"$id"`
	Type          o.Option[schematype.SchemaType]        `json:"type" yaml:"type"`
	Schema        o.Option[string]                       `json:"$schema" yaml:"$schema"`
	Comment       o.Option[string]                       `json:"$comment" yaml:"$comment"`
	Deprecated    o.Option[bool]                         `json:"deprecated" yaml:"deprecated"`
	Description   o.Option[string]                       `json:"description" yaml:"description"`
	Title         o.Option[string]                       `json:"title" yaml:"title"`
	ReadOnly      o.Option[bool]                         `json:"readOnly" yaml:"readOnly"`
	WriteOnly     o.Option[bool]                         `json:"writeOnly" yaml:"writeOnly"`
	Properties    mapper.Mapper[string, json.RawMessage] `json:"properties" yaml:"properties"`
	MaxProperties o.Option[int64]                        `json:"maxProperties" yaml:"maxProperties"`
	MinProperties o.Option[int64]                        `json:"minProperties" yaml:"minProperties"`
}

func (t *JsonSchemaObject) UnmarshalJSON(data []byte) error {
	var internal jsonSchemaObject

	if err := json.Unmarshal(data, &internal); err != nil {
		return err
	}

	t.Id = internal.Id
	t.Type = internal.Type
	t.Schema = internal.Schema
	t.Comment = internal.Comment
	t.Deprecated = internal.Deprecated
	t.Description = internal.Description
	t.Title = internal.Title
	t.ReadOnly = internal.ReadOnly
	t.WriteOnly = internal.WriteOnly
	t.MaxProperties = internal.MaxProperties
	t.MinProperties = internal.MinProperties

	if t.Properties == nil {
		t.Properties = mapper.Mapper[string, JsonSchema]{}
	}

	for name, raw := range internal.Properties {
		if property, err := UnmarshalSchema(raw); err != nil {
			return err
		} else {
			t.Properties[name] = property
		}
	}

	return nil
}
