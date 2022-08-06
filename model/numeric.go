package model

import o "github.com/boundedinfinity/optioner"

type JsonSchemaNumber[T ~int64 | ~float64] struct {
	JsonSchemaGeneric[T]
	ExclusiveMaximum o.Option[T] `json:"exclusiveMaximum" yaml:"exclusiveMaximum"`
	ExclusiveMinimum o.Option[T] `json:"exclusiveMinimum" yaml:"exclusiveMinimum"`
	Maximum          o.Option[T] `json:"maximum" yaml:"maximum"`
	Minimum          o.Option[T] `json:"minimum" yaml:"minimum"`
	MultipleOf       o.Option[T] `json:"multipleOf" yaml:"multipleOf"`
}

type JsonSchemaInteger[T ~int64] struct {
	JsonSchemaNumber[T]
}
