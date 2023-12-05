package functional

// import (
// 	"github.com/boundedinfinity/go-jsonschema/model"
// 	"github.com/boundedinfinity/go-marshaler"
// 	"github.com/boundedinfinity/go-mimetyper/mime_type"
// 	"github.com/ghodss/yaml"
// )

// func (t *System) LoadPath(paths ...string) error {
// 	for _, path := range paths {
// 		if err := t.loadPath(path); err != nil {
// 			return err
// 		}
// 	}

// 	return nil
// }

// func (t *System) loadPath(root string) error {
// 	m, err := marshaler.ReadFromPath(root)

// 	if err != nil {
// 		return err
// 	}

// 	for source, content := range m {
// 		if t.source2schema.Has(source) {
// 			return model.ErrPathDuplicatev(source)
// 		}

// 		if err := t.LoadSchema(root, source, content.Data, content.MimeType); err != nil {
// 			return err
// 		}
// 	}

// 	return nil
// }

// func (t *System) LoadSchema(root, source string, data []byte, mt mime_type.MimeType) error {
// 	var bs []byte
// 	var err error

// 	switch mt {
// 	case mime_type.ApplicationXYaml:
// 		bs, err = yaml.YAMLToJSON(data)

// 		if err != nil {
// 			return err
// 		}
// 	case mime_type.ApplicationJson:
// 		bs = data
// 	default:
// 		return model.ErrMimeTypeUnsupportedv(mt)
// 	}

// 	schema, err := model.UnmarshalSchema(bs)

// 	if err != nil {
// 		return err
// 	}

// 	if err := t.Register(root, source, schema); err != nil {
// 		return err
// 	}

// 	return nil
// }
