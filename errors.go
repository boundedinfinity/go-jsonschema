package jsonschema

import "fmt"

var (
	ErrUnknownUriType           = fmt.Errorf("unknown uri type")
	ErrUnknownFileType          = fmt.Errorf("unknown file type")
	ErrDuplicateDef             = fmt.Errorf("duplicate $def")
	ErrUnsupportedSchemaType    = fmt.Errorf("unsupported schema type")
	ErrRefNotFound              = fmt.Errorf("reference not found")
	ErrInvalidMultipleOf        = fmt.Errorf("multipleOf must be greater than 0")
	ErrInvalidMaxContains       = fmt.Errorf("maxContains must be a non-negative integer")
	ErrInvalidMaxLength         = fmt.Errorf("maxLength must be a non-negative integer")
	ErrInvalidMinLength         = fmt.Errorf("minLength must be a non-negative integer")
	ErrInvalidMaxItems          = fmt.Errorf("maxItems must be a non-negative integer")
	ErrInvalidMinItems          = fmt.Errorf("minItems must be a non-negative integer")
	ErrInvalidArrayMissingItems = fmt.Errorf("type array must have items")
	ErrNotNonNegative           = fmt.Errorf("must be a non-negative integer")
	ErrNotInteger               = fmt.Errorf("must be integer")
)

func ErrUnknownUriTypeV(v string) error  { return errV(v, ErrUnknownUriType) }
func ErrUnknownFileTypeV(v string) error { return errV(v, ErrUnknownFileType) }
func ErrDuplicateDefV(v string) error    { return errV(v, ErrDuplicateDef) }
func ErrRefNotFoundV(v string) error     { return errV(v, ErrRefNotFound) }

func ErrUnsupportedSchemaTypeV(v JsonSchemaType) error {
	return errV(v.String(), ErrUnsupportedSchemaType)
}

func errV(v string, err error) error {
	return fmt.Errorf("%v : %w", v, err)
}
