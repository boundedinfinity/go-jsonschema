package jsonschema

import (
	"net/url"
	"path"

	"github.com/boundedinfinity/optional"
)

func (t *System) mapSchema(schema JsonSchmea) error {
	id, err := url.Parse(schema.Id.Get())

	if err != nil {
		return err
	}

	for dname, def := range schema.Defs {
		fragment, err := url.Parse(path.Join("/", "$defs", dname))

		if err != nil {
			return err
		}

		qname := id.ResolveReference(fragment).String()

		if _, ok := t.uriMap[qname]; ok {
			return ErrDuplicateDefV(qname)
		}

		t.uriMap[qname] = def
	}

	for name, schema := range schema.Properties {
		fragment, err := url.Parse(name)

		if err != nil {
			return err
		}

		qname := id.ResolveReference(fragment).String()

		if _, ok := t.uriMap[qname]; ok {
			return ErrDuplicateDefV(qname)
		}

		t.uriMap[qname] = schema
	}

	return nil
}

func (t *System) resolveSchema(schema *JsonSchmea) error {
	for oname, obj := range t.uriMap {
		switch obj.Type {
		case JsonSchemaType_Object:
			if err := t.resolveObject(*schema, obj); err != nil {
				return err
			}
		case JsonSchemaType_String:
		case JsonSchemaType_Boolean:
		case JsonSchemaType_Integer:
		case JsonSchemaType_Array:
		case JsonSchemaType_Null:
		case JsonSchemaType_Number:
		default:
			if obj.Ref.IsDefined() {
				if robj, err := t.getRef(schema.Id, obj.Ref); err != nil {
					return err
				} else {
					t.uriMap[oname] = robj
				}
			} else {
				return ErrUnsupportedSchemaTypeV(obj.Type)
			}
		}
	}

	return nil
}

func (t *System) resolveObject(schema JsonSchmea, obj JsonSchmeaObject) error {
	for pname, pobj := range obj.Properties {
		if pobj.Ref.IsDefined() {
			if robj, err := t.getRef(schema.Id, obj.Ref); err != nil {
				return err
			} else {
				pobj.Properties[pname] = robj
			}
		}
	}

	return nil
}

func (t *System) getRef(id optional.StringOptional, ref optional.StringOptional) (JsonSchmeaObject, error) {
	pref, err := url.Parse(ref.Get())

	if err != nil {
		return JsonSchmeaObject{}, err
	}

	var key string

	if pref.Host != "" {
		key = pref.String()
	} else {
		pid, err := url.Parse(id.Get())

		if err != nil {
			return JsonSchmeaObject{}, err
		}

		key = pid.ResolveReference(pref).String()
	}

	if obj, ok := t.uriMap[key]; ok {
		return obj, nil
	}

	return JsonSchmeaObject{}, ErrRefNotFoundV(ref.Get())
}
