package jsonschema

// func (t *System) Resolve() error {
// 	for id := range t.IdMap {
// 		if err := t.ResolveId(id); err != nil {
// 			return err
// 		}
// 	}

// 	return nil
// }

// func (t *System) ResolveId(id string) error {
// 	obj := t.ById(id)

// 	if obj.Empty() {
// 		return nil
// 	}

// 	if !obj.Get().Type.Defined() || obj.Get().Type.Get() != objecttype.Object {
// 		return nil
// 	}

// 	p := obj.Get()

// 	if err := t.ResolveSchema(&p); err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (t *System) ResolveSchema(obj *JsonSchema) error {
// 	for name, prop := range obj.Properties {
// 		if prop.Ref.Defined() {
// 			ref := t.ById(prop.Ref.Get())

// 			if ref.Empty() {
// 				return ErrSchemaNotFoundV(prop.Ref.Get())
// 			}

// 			sch := ref.Get()
// 			obj.Properties[name] = &sch
// 		}

// 		if prop.Type.Defined() && prop.Type.Get() == objecttype.Array {
// 			if prop.Items == nil {
// 				return ErrArrayItemMissing
// 			}

// 			if prop.Items.Ref.Defined() {
// 				ref := t.ById(prop.Ref.Get())

// 				if ref.Empty() {
// 					return ErrSchemaNotFoundV(prop.Ref.Get())
// 				}

// 				sch := ref.Get()
// 				prop.Items = &sch
// 			}
// 		}
// 	}

// 	return nil
// }
