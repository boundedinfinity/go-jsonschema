package jsonschema

import (
	"github.com/boundedinfinity/optional"
)

func (t *System) mapSchema(schema JsonSchmea, file string) error {
	if schema.Id.IsDefined() {
		t.i2l[schema.Id.Get()] = file
		t.l2i[file] = schema.Id.Get()
		t.schmeaMap[schema.Id.Get()] = &schema
	}

	for name, obj := range schema.Defs {
		qname, err := urlJoin(schema.Id.Get(), "$defs", name)

		if err != nil {
			return err
		}

		if _, ok := t.schmeaMap[qname]; ok {
			return ErrDuplicateDefV(qname)
		}

		t.schmeaMap[qname] = &obj
	}

	switch schema.Type {
	case JsonSchemaType_String, JsonSchemaType_Number, JsonSchemaType_Integer, JsonSchemaType_Boolean, JsonSchemaType_Null:
		t.schmeaMap[schema.Id.Get()] = &schema
	case JsonSchemaType_Object:
		for name, obj := range schema.Properties {
			qname, err := urlJoin(schema.Id.Get(), name)

			if err != nil {
				return err
			}

			if _, ok := t.schmeaMap[qname]; ok {
				return ErrDuplicateDefV(qname)
			}

			t.schmeaMap[qname] = &obj
		}
	case JsonSchemaType_Array:
	}

	return nil
}

func (t *System) resolveSchema(schema *JsonSchmea) error {
	if schema.Ref.IsDefined() {
		id := schema.Ref.Get()
		if ref, ok := t.schmeaMap[id]; ok {
			schema = ref
		} else {
			return ErrRefNotFoundV(schema.Ref.Get())
		}
	}
	// switch obj.Type {
	// case JsonSchemaType_Object:

	// case JsonSchemaType_String:
	// case JsonSchemaType_Boolean:
	// case JsonSchemaType_Integer:
	// case JsonSchemaType_Array:
	// 	// if err := t.resolveSchema(obj.Items); err != nil {
	// 	// 	return err
	// 	// }
	// case JsonSchemaType_Null:
	// case JsonSchemaType_Number:
	// default:
	// 	if obj.Ref.IsDefined() {
	// 		if robj, err := t.getRef(schema.Id, obj.Ref); err != nil {
	// 			return err
	// 		} else {
	// 			t.typeMap[oname] = robj
	// 		}
	// 	} else {
	// 		return ErrUnsupportedSchemaTypeV(obj.Type)
	// 	}
	// }

	return nil
}

func (t *System) resolveObject(schema JsonSchmea, obj JsonSchmea) error {
	// for pname, pobj := range obj.Properties {
	// 	if pobj.Ref.IsDefined() {
	// 		if robj, err := t.getRef(schema.Id, obj.Ref); err != nil {
	// 			return err
	// 		} else {
	// 			pobj.Properties[pname] = *robj
	// 		}
	// 	}
	// }

	return nil
}

func (t *System) getRef(id optional.StringOptional, ref optional.StringOptional) (*JsonSchmea, error) {
	// var key string

	// if strings.HasPrefix(ref.String(), "#") {
	// 	idUrl, err := url.Parse(id.Get())

	// 	if err != nil {
	// 		return &JsonSchmea{}, err
	// 	}

	// 	idUrl.Path = path.Join(idUrl.Path, strings.TrimPrefix(ref.Get(), "#"))
	// 	key = idUrl.String()
	// } else {
	// 	key = ref.String()
	// }

	// if obj, ok := t.typeMap[key]; ok {
	// 	return obj, nil
	// }

	return &JsonSchmea{}, ErrRefNotFoundV(ref.Get())
}
