package jsonschema

import (
	"github.com/boundedinfinity/optional"
)

func (t *System) mapSchema(schema *JsonSchmea, file string) error {
	id := schema.Id.Get()

	if schema.Id.IsDefined() {
		t.i2l[id] = file
		t.l2i[file] = id
		t.schmeaMap[id] = schema
	}

	for name, obj := range schema.Defs {
		qname, err := urlJoin(id, "$defs", name)

		if err != nil {
			return err
		}

		if _, ok := t.schmeaMap[qname]; ok {
			return ErrDuplicateDefV(qname)
		}

		t.schmeaMap[qname] = obj
	}

	switch schema.Type {
	case JsonSchemaType_String, JsonSchemaType_Number, JsonSchemaType_Integer, JsonSchemaType_Boolean, JsonSchemaType_Null:
		t.schmeaMap[id] = schema
	case JsonSchemaType_Object:
		for name, obj := range schema.Properties {
			qname, err := urlJoin(id, name)

			if err != nil {
				return err
			}

			if _, ok := t.schmeaMap[qname]; ok {
				return ErrDuplicateDefV(qname)
			}

			t.schmeaMap[qname] = obj
		}
	case JsonSchemaType_Array:
	}

	return nil
}

func (t *System) resolveSchema(schema *JsonSchmea) error {
	if schema.Ref.IsDefined() {
		id := schema.Ref.Get()
		if ref, ok := t.schmeaMap[id]; ok {
			*schema = *ref
		} else {
			return ErrRefNotFoundV(schema.Ref.Get())
		}
	} else {
		switch schema.Type {
		case JsonSchemaType_Array:
			if err := t.resolveSchema(schema.Items); err != nil {
				return err
			}
		}
	}

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
