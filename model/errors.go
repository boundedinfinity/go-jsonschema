package model

import (
	"github.com/boundedinfinity/go-commoner/errorer"
	"golang.org/x/exp/constraints"
)

var (
	ErrJsonSchemaTypeNotFound             = errorer.Errorf("json schema type not found")
	ErrJsonSchemaTypeInvalid              = errorer.Errorf("json schema type invalid")
	ErrJsonSchemaTypeInvalidv             = errorer.Wrapv(ErrJsonSchemaTypeInvalid)
	ErrJsonSchemaTypeUnsupported          = errorer.Errorf("json schema type unsupported")
	ErrJsonSchemaTypeUnsupportedv         = errorer.Wrapv(ErrJsonSchemaTypeUnsupported)
	ErrJsonSchemaIdNotFound               = errorer.Errorf("json schema ID not found")
	ErrJsonSchemaIdNotFoundv              = errorer.Wrapv(ErrJsonSchemaIdNotFound)
	ErrJsonSchemaCantResolve              = errorer.Errorf("json schema can't resolve")
	ErrJsonSchemaCantResolvev             = errorer.Wrapv(ErrJsonSchemaCantResolve)
	ErrJsonSchemaIdDuplicate              = errorer.Errorf("json schema ID duplicate")
	ErrJsonSchemaIdDuplicatev             = errorer.Wrapv(ErrJsonSchemaIdDuplicate)
	ErrJsonPathDuplicate                  = errorer.Errorf("json schema path duplicate")
	ErrJsonPathDuplicatev                 = errorer.Wrapv(ErrJsonPathDuplicate)
	ErrJsonSchemaNotFound                 = errorer.Errorf("json schema not found")
	ErrJsonSchemaNotFoundv                = errorer.Wrapv(ErrJsonSchemaNotFound)
	ErrJsonSchemaRefEmpty                 = errorer.Errorf("reference is empty")
	ErrJsonSchemaRefNotFound              = errorer.Errorf("reference not found")
	ErrJsonSchemaRefNotFoundv             = errorer.Wrapv(ErrJsonSchemaRefNotFound)
	ErrJsonSchemaObjectPropertiesNotFound = errorer.Errorf("object properties not found")
	ErrJsonSchemaTypeOrRefNotFound        = errorer.Errorf("json schema type or $ref not found missing")
	ErrJsonSchemaMimeTypeUnsupported      = errorer.Errorf("MIME type unsupported")
	ErrJsonSchemaMimeTypeUnsupportedv     = errorer.Wrapv(ErrJsonSchemaMimeTypeUnsupported)
	ErrJsonSchemaNumberMultipleOfInvalid  = errorer.Errorf("multipleOf must be greater than 0")
	ErrJsonSchemaNumberMultipleOfInvalidv = errorer.Wrapv(ErrJsonSchemaNumberMultipleOfInvalid)
	ErrJsonSchemaArrayMaxContainsInvalid  = errorer.Errorf("maxContains must be greater than 0")
	ErrJsonSchemaArrayMaxContainsInvalidv = errorer.Wrapv(ErrJsonSchemaArrayMaxContainsInvalid)
	ErrJsonSchemaArrayMinContainsInvalid  = errorer.Errorf("minContains must be greater than 0")
	ErrJsonSchemaArrayMinContainsInvalidv = errorer.Wrapv(ErrJsonSchemaArrayMinContainsInvalid)
	ErrJsonSchemaArrayInvalidMaxMin       = errorer.Errorf("maxContains must be greater than minContains")
	ErrJsonSchemaArrayItemsEmpty          = errorer.Errorf("array items is empty")
	ErrJsonSchemaNumberInvalidMaxMin      = errorer.Errorf("maximum must be greater than minimum")
	ErrJsonSchemaStringMaxLength          = errorer.Errorf("maxLength must be greater than 0")
	ErrJsonSchemaStringMaxLengthv         = errorer.Wrapv(ErrJsonSchemaStringMaxLength)
	ErrJsonSchemaStringMinLength          = errorer.Errorf("minLength must be greater than 0")
	ErrJsonSchemaStringMinLengthv         = errorer.Wrapv(ErrJsonSchemaStringMinLength)
	ErrJsonSchemaStringMaxMinLength       = errorer.Errorf("maxLength must be greater than minLength")
	ErrJsonSchemaObjectMaxProperties      = errorer.Errorf("maxProperties must be greater than 0")
	ErrJsonSchemaObjectMaxPropertiesv     = errorer.Wrapv(ErrJsonSchemaObjectMaxProperties)
	ErrJsonSchemaObjectMinProperties      = errorer.Errorf("minProperties must be greater than 0")
	ErrJsonSchemaObjectMinPropertiesv     = errorer.Wrapv(ErrJsonSchemaObjectMinProperties)
	ErrJsonSchemaObjectMaxMinProperties   = errorer.Errorf("maxProperties must be greater than minProperties")
	ErrJsonSchemaObjectPropertiesMissing  = errorer.Errorf("properties must be greater than 0")
)

type number interface {
	constraints.Integer | constraints.Float
}

func ErrJsonSchemaNumberInvalidMaxMinv[T number](min, max T) error {
	return errorer.Wrapf(ErrJsonSchemaNumberInvalidMaxMin)("min: %v, max: %v", min, max)
}

func ErrJsonSchemaArrayInvalidMaxMinv[T number](min, max T) error {
	return errorer.Wrapf(ErrJsonSchemaArrayInvalidMaxMin)("min: %v, max: %v", min, max)
}

func ErrJsonSchemaStringMaxMinLengthv[T number](min, max T) error {
	return errorer.Wrapf(ErrJsonSchemaStringMaxMinLength)("min: %v, max: %v", min, max)
}

func ErrJsonSchemaObjectMaxMinPropertiesv[T number](min, max T) error {
	return errorer.Wrapf(ErrJsonSchemaObjectMaxMinProperties)("min: %v, max: %v", min, max)
}
