package jsonschema

import (
	"errors"
	"fmt"

	"github.com/boundedinfinity/jsonschema/objecttype"
	"github.com/boundedinfinity/jsonschema/stringformat"
	"github.com/boundedinfinity/mimetyper/mime_type"
)

var (
	ErrIdEmpty                  = errors.New("json schema ID empty")
	ErrSchemaNotFound           = errors.New("json schema not found")
	ErrUriTypeUnsupported       = errors.New("unknown uri type")
	ErrMimeTypeUnsupported      = errors.New("unknown MEME type")
	ErrMimeTypeUndetected       = errors.New("undeteched MIME type")
	ErrDuplicateDef             = errors.New("duplicate $def")
	ErrUnsupportedSchemaType    = errors.New("unsupported schema type")
	ErrRefNotFound              = errors.New("reference not found")
	ErrInvalidMultipleOf        = errors.New("multipleOf must be greater than 0")
	ErrInvalidMaxContains       = errors.New("maxContains must be a non-negative integer")
	ErrInvalidMaxLength         = errors.New("maxLength must be a non-negative integer")
	ErrInvalidMinLength         = errors.New("minLength must be a non-negative integer")
	ErrInvalidMaxItems          = errors.New("maxItems must be a non-negative integer")
	ErrInvalidMinItems          = errors.New("minItems must be a non-negative integer")
	ErrInvalidArrayMissingItems = errors.New("type array must have items")
	ErrNotNonNegative           = errors.New("must be a non-negative integer")
	ErrNotInteger               = errors.New("must be integer")
	ErrMinGreaterThanMax        = errors.New("minLength cannot be greater than maxLength")
	ErrStringInvalidPattern     = errors.New("invalid pattern")
	ErrObjectTypeEmpty          = errors.New("type cannot be empty")
	ErrStringFormatInvalid      = errors.New("invalid string format")
	ErrStringFormatUnsupported  = errors.New("unsupported string format")
)

func ErrUriTypeUnsupportedV(v string) error { return errV(v, ErrUriTypeUnsupported) }

func ErrMimeTypeUnsupportedV(v mime_type.MimeType) error {
	return errV(v.String(), ErrMimeTypeUnsupported)
}

func ErrDuplicateDefV(v string) error { return errV(v, ErrDuplicateDef) }
func ErrRefNotFoundV(v string) error  { return errV(v, ErrRefNotFound) }

func ErrUnsupportedSchemaTypeV(v objecttype.ObjectType) error {
	return errV(v.String(), ErrUnsupportedSchemaType)
}

func ErrSchemaNotFoundV(v string) error { return errV(v, ErrSchemaNotFound) }
func ErrStringInvalidPatternV(v string, err error) error {
	return fmt.Errorf("%v : %w : %w", v, ErrStringInvalidPattern, err)
}

func ErrStringFormatInvalidV(v stringformat.StringFormatOption) error {
	return errV(v.String(), ErrStringFormatInvalid)
}

func ErrStringFormatUnsupportedV(v stringformat.StringFormatOption) error {
	return errV(v.String(), ErrStringFormatUnsupported)
}

func ErrMinGreaterThanMaxV(min, max int) error {
	return fmt.Errorf("%v > %v : %w", min, max, ErrMinGreaterThanMax)
}

func errV(v string, err error) error {
	return fmt.Errorf("%v : %w", v, err)
}
