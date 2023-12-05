package functional

// import (
// 	o "github.com/boundedinfinity/go-commoner/optioner"
// 	"github.com/boundedinfinity/go-jsonschema/model"
// )

// func (t *System) Resolve(schema model.JsonSchema) o.Option[model.JsonSchema] {
// 	switch schema.(type) {
// 	case model.JsonSchemaRef:
// 		return t.id2schema.Get(schema.Base().Id.Get())
// 	case *model.JsonSchemaRef:
// 		return t.id2schema.Get(schema.Base().Id.Get())
// 	default:
// 		return o.Some(schema)
// 	}
// }

// func (t *System) Validate() error {
// 	for _, schema := range t.All {
// 		resolved := t.Resolve(schema)

// 		if resolved.Empty() {
// 			model.ErrSchemaCantResolvev(schema.Base().Id)
// 		}
// 	}

// 	return nil
// }
