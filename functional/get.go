package functional

// import (
// 	o "github.com/boundedinfinity/go-commoner/optioner"
// 	"github.com/boundedinfinity/go-jsonschema/model"
// 	"github.com/boundedinfinity/go-jsonschema/schematype"
// )

// func (t *System) AllId() []model.JsonSchema {
// 	return t.id2schema.Values().Get()
// }

// func (t *System) AllPath() []model.JsonSchema {
// 	return t.source2schema.Values().Get()
// }

// func (t *System) Id(schema model.JsonSchema) o.Option[string] {
// 	if schema == nil || schema.Base() == nil {
// 		return o.None[string]()
// 	}

// 	return schema.Base().Id
// }

// func (t *System) Ref(schema model.JsonSchema) o.Option[string] {
// 	if schema == nil {
// 		return o.None[string]()
// 	}

// 	switch c := schema.(type) {
// 	case model.JsonSchemaRef:
// 		return c.Ref
// 	case *model.JsonSchemaRef:
// 		return c.Ref
// 	default:
// 		return o.None[string]()
// 	}
// }

// func (t *System) Get(id string) o.Option[model.JsonSchema] {
// 	return o.FirstOf(t.id2schema.Get(id), t.source2schema.Get(id))
// }

// func (t *System) GetSource(id string) o.Option[string] {
// 	s := t.Get(id)

// 	if s.Empty() || s.Get().Base().Id.Empty() {
// 		return o.None[string]()
// 	}

// 	return t.id2source.Get(s.Get().Base().Id.Get())
// }

// func (t *System) GetRoot(id string) o.Option[string] {
// 	s := t.Get(id)

// 	if s.Empty() || s.Get().Base().Id.Empty() {
// 		return o.None[string]()
// 	}

// 	return t.id2root.Get(s.Get().Base().Id.Get())
// }

// func (t *System) GetType(schema model.JsonSchema) o.Option[schematype.SchemaType] {
// 	if schema == nil {
// 		return o.None[schematype.SchemaType]()
// 	}

// 	switch schema.(type) {
// 	case model.JsonSchemaArray:
// 		return o.Some(schematype.Array)
// 	case *model.JsonSchemaArray:
// 		return o.Some(schematype.Array)
// 	case model.JsonSchemaInteger:
// 		return o.Some(schematype.Integer)
// 	case *model.JsonSchemaInteger:
// 		return o.Some(schematype.Integer)
// 	case model.JsonSchemaNumber:
// 		return o.Some(schematype.Number)
// 	case *model.JsonSchemaNumber:
// 		return o.Some(schematype.Number)
// 	case model.JsonSchemaObject:
// 		return o.Some(schematype.Object)
// 	case *model.JsonSchemaObject:
// 		return o.Some(schematype.Object)
// 	case model.JsonSchemaString:
// 		return o.Some(schematype.String)
// 	case *model.JsonSchemaString:
// 		return o.Some(schematype.String)
// 	case model.JsonSchemaBoolean:
// 		return o.Some(schematype.Boolean)
// 	case model.JsonSchemaNull:
// 		return o.Some(schematype.Null)
// 	case *model.JsonSchemaNull:
// 		return o.Some(schematype.Null)
// 	default:
// 		return o.None[schematype.SchemaType]()
// 	}
// }

// func (t *System) IsRef(schema model.JsonSchema) bool {
// 	if schema == nil {
// 		return false
// 	}

// 	switch schema.(type) {
// 	case model.JsonSchemaRef, *model.JsonSchemaRef:
// 		return true
// 	default:
// 		return false
// 	}
// }

// func (t *System) IsConcrete(schema model.JsonSchema) bool {
// 	return !t.IsRef(schema)
// }
