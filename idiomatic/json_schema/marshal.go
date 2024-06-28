package json_schema

type jsonSchemaDescriminator struct {
	Type string `json:"type" yaml:"type"`
	Ref  string `json:"$ref" yaml:"$ref"`
}

// func NewMarshaler() *JsonSchemaMarshaler {
// 	m := &JsonSchemaMarshaler{}

// 	return m
// }

// type JsonSchemaMarshaler struct {
// 	km *marshaler.KindMarshaler[jsonSchemaDescriminator]
// }

// func (t JsonSchemaMarshaler) Unmarshal(data []byte) (JsonSchema, error) {
// 	// var val any
// 	// var err error

// 	// val, err = t.km.Unmarshal(data)

// 	// if err != nil && !errors.Is(err, marshaler.ErrKindMarshalerTypeNotRegistered) {
// 	// 	return nil, err
// 	// }

// 	// schema, ok := val.(JsonSchema)

// 	// if !ok {
// 	// 	return nil, fmt.Errorf("something")
// 	// }

// 	// switch val := schema.(type) {
// 	// case *JsonSchemaRef:

// 	// }

// 	return nil, nil
// }

// // func UnmarshalSchema(data []byte) (JsonSchema, error) {
// // 	var out JsonSchema
// // 	var d jsonSchemaDescriminator

// // 	if err := json.Unmarshal(data, &d); err != nil {
// // 		return nil, err
// // 	}

// // 	if d.Type.Defined() {
// // 		switch d.Type.Get() {
// // 		case schematype.String:
// // 			var v JsonSchemaString

// // 			if err := json.Unmarshal(data, &v); err != nil {
// // 				return nil, err
// // 			}

// // 			out = &v
// // 		case schematype.Integer:
// // 			var v JsonSchemaInteger

// // 			if err := json.Unmarshal(data, &v); err != nil {
// // 				return nil, err
// // 			}

// // 			out = &v
// // 		case schematype.Number:
// // 			var v JsonSchemaNumber

// // 			if err := json.Unmarshal(data, &v); err != nil {
// // 				return nil, err
// // 			}

// // 			out = &v
// // 		case schematype.Array:
// // 			var v JsonSchemaArray

// // 			if err := json.Unmarshal(data, &v); err != nil {
// // 				return nil, err
// // 			}

// // 			out = &v
// // 		case schematype.Object:
// // 			var v JsonSchemaObject

// // 			if err := json.Unmarshal(data, &v); err != nil {
// // 				return nil, err
// // 			}

// // 			for name, raw := range d.Properties {
// // 				if property, err := UnmarshalSchema(raw); err != nil {
// // 					return nil, err
// // 				} else {
// // 					v.Properties[name] = property
// // 				}
// // 			}

// // 			out = &v
// // 		default:
// // 			if d.Ref.Defined() {
// // 				return nil, fmt.Errorf("something went wrong")
// // 			}
// // 		}
// // 	}

// // 	if d.Ref.Defined() {
// // 		var v JsonSchemaRef

// // 		if err := json.Unmarshal(data, &v); err != nil {
// // 			return nil, err
// // 		}

// // 		out = &v
// // 	}

// // 	if out == nil {
// // 		return nil, fmt.Errorf("something went wrong")
// // 	}

// // 	return out, nil
// // }
