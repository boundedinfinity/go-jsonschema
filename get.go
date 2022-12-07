package jsonschema

import (
	o "github.com/boundedinfinity/go-commoner/optioner"
	"github.com/boundedinfinity/go-jsonschema/model"
	"github.com/boundedinfinity/go-jsonschema/schematype"
)

func (t *System) All() []model.JsonSchema {
	return t.idMap.Values().Get()
}

func (t *System) Id(schema model.JsonSchema) o.Option[string] {
	if schema == nil {
		return o.None[string]()
	}

	switch c := schema.(type) {
	case model.JsonSchemaArray:
		return t.Id(c.Items.Get())
	case *model.JsonSchemaArray:
		return t.Id(c.Items.Get())
	case model.JsonSchemaInteger:
		return c.Id
	case *model.JsonSchemaInteger:
		return c.Id
	case model.JsonSchemaNumber:
		return c.Id
	case *model.JsonSchemaNumber:
		return c.Id
	case model.JsonSchemaObject:
		return c.Id
	case *model.JsonSchemaObject:
		return c.Id
	case model.JsonSchemaString:
		return c.Id
	case *model.JsonSchemaString:
		return c.Id
	case model.JsonSchemaBoolean:
		return c.Id
	case model.JsonSchemaNull:
		return c.Id
	case *model.JsonSchemaNull:
		return c.Id
	default:
		return o.None[string]()
	}
}

func (t *System) Ref(schema model.JsonSchema) o.Option[string] {
	if schema == nil {
		return o.None[string]()
	}

	switch c := schema.(type) {
	case model.JsonSchemaRef:
		return c.Ref
	case *model.JsonSchemaRef:
		return c.Ref
	default:
		return o.None[string]()
	}
}

func (t *System) Get(v string) o.Option[model.JsonSchema] {
	if t.pathMap.Has(v) {
		return t.pathMap.Get(v)
	}

	if t.idMap.Has(v) {
		return t.idMap.Get(v)
	}

	return o.None[model.JsonSchema]()
}

func (t *System) GetType(schema model.JsonSchema) o.Option[schematype.SchemaType] {
	if schema == nil {
		return o.None[schematype.SchemaType]()
	}

	switch schema.(type) {
	case model.JsonSchemaArray:
		return o.Some(schematype.Array)
	case *model.JsonSchemaArray:
		return o.Some(schematype.Array)
	case model.JsonSchemaInteger:
		return o.Some(schematype.Integer)
	case *model.JsonSchemaInteger:
		return o.Some(schematype.Integer)
	case model.JsonSchemaNumber:
		return o.Some(schematype.Number)
	case *model.JsonSchemaNumber:
		return o.Some(schematype.Number)
	case model.JsonSchemaObject:
		return o.Some(schematype.Object)
	case *model.JsonSchemaObject:
		return o.Some(schematype.Object)
	case model.JsonSchemaString:
		return o.Some(schematype.String)
	case *model.JsonSchemaString:
		return o.Some(schematype.String)
	case model.JsonSchemaBoolean:
		return o.Some(schematype.Boolean)
	case model.JsonSchemaNull:
		return o.Some(schematype.Null)
	case *model.JsonSchemaNull:
		return o.Some(schematype.Null)
	default:
		return o.None[schematype.SchemaType]()
	}
}

func (t *System) IsRef(schema model.JsonSchema) bool {
	if schema == nil {
		return false
	}

	switch schema.(type) {
	case model.JsonSchemaRef, *model.JsonSchemaRef:
		return true
	default:
		return false
	}
}

func (t *System) IsConcrete(schema model.JsonSchema) bool {
	return !t.IsRef(schema)
}
