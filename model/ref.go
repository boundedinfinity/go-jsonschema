package model

import (
	o "github.com/boundedinfinity/go-commoner/optioner"
)

func NewRef(ref string) *JsonSchemaRef {
	schema := &JsonSchemaRef{
		Ref: o.OfZ(IdT(ref)),
	}

	return schema
}

type JsonSchemaRef struct {
	Id          o.Option[IdT]          `json:"$id" yaml:"$id"`
	Ref         o.Option[IdT]          `json:"$ref" yaml:"$ref"`
	Schema      o.Option[SchemaT]      `json:"$schema" yaml:"$schema"`
	Title       o.Option[TitleT]       `json:"title" yaml:"title"`
	Description o.Option[DescriptionT] `json:"description" yaml:"description"`
}

var _ = &JsonSchemaRef{}

func (t JsonSchemaRef) GetId() o.Option[IdT] {
	return t.Id
}

func (t JsonSchemaRef) GetRef() o.Option[IdT] {
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
