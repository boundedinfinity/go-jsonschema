package model

import (
	o "github.com/boundedinfinity/go-commoner/optioner"
)

func NewRef(ref string) *JsonSchemaRef {
	schema := &JsonSchemaRef{
		Ref: o.OfZ(ref),
	}

	return schema
}

type JsonSchemaRef struct {
	Id          o.Option[string] `json:"$id" yaml:"$id"`
	Ref         o.Option[string] `json:"$ref" yaml:"$ref"`
	Schema      o.Option[string] `json:"$schema" yaml:"$schema"`
	Title       o.Option[string] `json:"title" yaml:"title"`
	Description o.Option[string] `json:"description" yaml:"description"`
}

var _ = &JsonSchemaRef{}

func (t JsonSchemaRef) GetId() o.Option[string] {
	return t.Id
}

func (t JsonSchemaRef) GetRef() o.Option[string] {
	return t.Ref
}

func (t JsonSchemaRef) IsConcrete() bool {
	return false
}

func (t JsonSchemaRef) IsRef() bool {
	return true
}

func (t JsonSchemaRef) Validate() error {
	return nil
}
