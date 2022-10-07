package model

import (
	"errors"

	"github.com/boundedinfinity/go-commoner/errorer"
)

var (
	ErrSchemaTypeMissing     = errors.New("json schema type missing")
	ErrSchemaTypeUnsupported = errors.New("json schema type unsupported")

	// ErrFileTypeUnsupported = errors.New("unsupported schema type")
	// ErrFileTypeUnsupportedFilef = func(path string, extentions []string) error {
	// 	return fmt.Errorf("%v %w : must be one of %v", path, ErrFileTypeUnsupported, strings.Join(extentions, ", "))
	// }

	// ErrIdEmpty                = errors.New("json schema ID empty")
	// ErrIdEmptyf               = func(key string) error { return fmt.Errorf("%v : %w", key, ErrIdEmpty) }
	// ErrSchemaTypeMissing      = errors.New("missing schema type")
	// ErrSchemaTypeUnsupported  = errors.New("unsupported schema type")
	// ErrSchemaTypeUnsupportedV = func(v schematype.SchemaType) error { return errV(v.String(), ErrSchemaTypeUnsupported) }

	// ErrSchemaNotFound    = errors.New("json schema not found")
	// ErrPropertyNotFound  = errors.New("property not found")
	// ErrPropertyNotFoundV = func(propname string) error { return fmt.Errorf("%v : %w", propname, ErrPropertyNotFound) }

	// ErrUriTypeUnsupported = errors.New("unknown uri type")
	// ErrMimeTypeUndetected = errors.New("undeteched MIME type")
	// ErrDuplicateDef       = errors.New("duplicate $def")

	// ErrRefNotFound              = errors.New("reference not found")
	// ErrInvalidMultipleOf        = errors.New("multipleOf must be greater than 0")
	// ErrInvalidMaxContains       = errors.New("maxContains must be a non-negative integer")
	// ErrInvalidMaxLength         = errors.New("maxLength must be a non-negative integer")
	// ErrInvalidMinLength         = errors.New("minLength must be a non-negative integer")
	// ErrInvalidMaxItems          = errors.New("maxItems must be a non-negative integer")
	// ErrInvalidMinItems          = errors.New("minItems must be a non-negative integer")
	// ErrInvalidArrayMissingItems = errors.New("type array must have items")
	// ErrNotNonNegative           = errors.New("must be a non-negative integer")
	// ErrNotInteger               = errors.New("must be integer")
	// ErrMinGreaterThanMax        = errors.New("minLength cannot be greater than maxLength")
	// ErrStringInvalidPattern     = errors.New("invalid pattern")
	// ErrObjectTypeEmpty          = errors.New("type cannot be empty")
	// ErrStringFormatInvalid      = errors.New("invalid string format")
	// ErrStringFormatUnsupported  = errors.New("unsupported string format")
	// ErrArrayItemMissing         = errors.New("unsupported string format")
)

// Validation Errors

var (
	ErrSchemaTypeUnsupportedV = errorer.Errorfn(ErrSchemaTypeUnsupported)

// ErrNotMultipleOf  = errors.New("not a multiple of")
// ErrNotMultipleOff = func(v, x int) error { return fmt.Errorf("%v %v %v", v, ErrNotMultipleOf, x) }

// ErrIsLessThan  = errors.New("is less than")
// ErrIsLessThanf = func(v, x int) error { return fmt.Errorf("%v %v %v", v, ErrIsLessThan, x) }

// ErrIsLessThanOrEqualTo  = errors.New("is less than or equal to")
// ErrIsLessThanOrEqualTof = func(v, x int) error { return fmt.Errorf("%v %v %v", v, ErrIsLessThanOrEqualTo, x) }

// ErrIsGreaterThan  = errors.New("is greater than")
// ErrIsGreaterThanf = func(v, x int) error { return fmt.Errorf("%v %v %v", v, ErrIsGreaterThan, x) }

// ErrIsGreaterThanOrEqualTo  = errors.New("is greater than or equal to")
// ErrIsGreaterThanOrEqualTof = func(v, x int) error { return fmt.Errorf("%v %v %v", v, ErrIsGreaterThanOrEqualTo, x) }
)

// func ErrUriTypeUnsupportedV(v string) error { return errV(v, ErrUriTypeUnsupported) }

// func ErrDuplicateDefV(v string) error { return errV(v, ErrDuplicateDef) }
// func ErrRefNotFoundV(v string) error  { return errV(v, ErrRefNotFound) }

// func ErrSchemaNotFoundV(v string) error { return errV(v, ErrSchemaNotFound) }
// func ErrStringInvalidPatternV(v string, err error) error {
// 	return fmt.Errorf("%v : %v : %w", v, ErrStringInvalidPattern, err)
// }

// func ErrStringFormatInvalidV(v o.Option[stringformat.StringFormat]) error {
// 	return errV(v.Get().String(), ErrStringFormatInvalid)
// }

// func ErrStringFormatUnsupportedV(v o.Option[stringformat.StringFormat]) error {
// 	return errV(v.Get().String(), ErrStringFormatUnsupported)
// }

// func ErrMinGreaterThanMaxV(min, max int) error {
// 	return fmt.Errorf("%v > %v : %w", min, max, ErrMinGreaterThanMax)
// }

// func errV(v string, err error) error {
// 	return fmt.Errorf("%v : %w", v, err)
// }
