package functional

// import (
// 	"encoding/json"
// 	"fmt"

// 	"github.com/boundedinfinity/go-jsonschema/schematype"
// )

// // Reference
// // https://medium.com/@nate510/dynamic-json-umarshalling-in-go-88095561d6a0

// func UnmarshalSchema(data []byte) (JsonSchema, error) {
// 	var out JsonSchema
// 	var d jsonSchemaDescriminator

// 	if err := json.Unmarshal(data, &d); err != nil {
// 		return nil, err
// 	}

// 	if d.Type.Defined() {
// 		switch d.Type.Get() {
// 		case schematype.String:
// 			var v JsonSchemaString

// 			if err := json.Unmarshal(data, &v); err != nil {
// 				return nil, err
// 			}

// 			out = &v
// 		case schematype.Integer:
// 			var v JsonSchemaInteger

// 			if err := json.Unmarshal(data, &v); err != nil {
// 				return nil, err
// 			}

// 			out = &v
// 		case schematype.Number:
// 			var v JsonSchemaNumber

// 			if err := json.Unmarshal(data, &v); err != nil {
// 				return nil, err
// 			}

// 			out = &v
// 		case schematype.Array:
// 			var v JsonSchemaArray

// 			if err := json.Unmarshal(data, &v); err != nil {
// 				return nil, err
// 			}

// 			out = &v
// 		case schematype.Object:
// 			var v JsonSchemaObject

// 			if err := json.Unmarshal(data, &v); err != nil {
// 				return nil, err
// 			}

// 			for name, raw := range d.Properties {
// 				if property, err := UnmarshalSchema(raw); err != nil {
// 					return nil, err
// 				} else {
// 					v.Properties[name] = property
// 				}
// 			}

// 			out = &v
// 		default:
// 			if d.Ref.Defined() {
// 				return nil, fmt.Errorf("something went wrong")
// 			}
// 		}
// 	}

// 	if d.Ref.Defined() {
// 		var v JsonSchemaRef

// 		if err := json.Unmarshal(data, &v); err != nil {
// 			return nil, err
// 		}

// 		out = &v
// 	}

// 	if out == nil {
// 		return nil, fmt.Errorf("something went wrong")
// 	}

// 	return out, nil
// }
