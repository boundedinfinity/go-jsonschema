package model

import (
	"github.com/boundedinfinity/go-commoner/errorer"
)

var (
	ErrSchemaTypeNotFound       = errorer.Errorf("json schema type not found")
	ErrSchemaTypeUnsupported    = errorer.Errorf("json schema type unsupported")
	ErrSchemaTypeUnsupportedv   = errorer.Errorfn(ErrSchemaTypeUnsupported)
	ErrSchemaIdNotFound         = errorer.Errorf("json schema ID not found")
	ErrSchemaIdDuplicate        = errorer.Errorf("json schema ID duplicate")
	ErrSchemaIdDuplicatev       = errorer.Errorfn(ErrSchemaIdDuplicate)
	ErrPathDuplicate            = errorer.Errorf("json schema path duplicate")
	ErrPathDuplicatev           = errorer.Errorfn(ErrPathDuplicate)
	ErrSchemaNotFound           = errorer.Errorf("json schema not found")
	ErrSchemaNotFoundv          = errorer.Errorfn(ErrSchemaNotFound)
	ErrRefEmpty                 = errorer.Errorf("reference is empty")
	ErrRefNotFound              = errorer.Errorf("reference not found")
	ErrRefNotFoundv             = errorer.Errorfn(ErrRefNotFound)
	ErrObjectPropertiesNotFound = errorer.Errorf("object properties not found")
	ErrSchemaTypeOrRefNotFound  = errorer.Errorf("json schema type or $ref not found missing")
	ErrMimeTypeUnsupported      = errorer.Errorf("MIME type unsupported")
	ErrMimeTypeUnsupportedv     = errorer.Errorfn(ErrMimeTypeUnsupported)
	ErrNumberMultipleOfInvalid  = errorer.Errorf("multipleOf must be greater than 0")
	ErrNumberMultipleOfInvalidv = errorer.Errorfn(ErrNumberMultipleOfInvalid)
	ErrArrayMaxContainsInvalid  = errorer.Errorf("maxContains must be greater than 0")
	ErrArrayMaxContainsInvalidv = errorer.Errorfn(ErrArrayMaxContainsInvalid)
	ErrArrayMinContainsInvalid  = errorer.Errorf("minContains must be greater than 0")
	ErrArrayMinContainsInvalidv = errorer.Errorfn(ErrArrayMinContainsInvalid)
	ErrArrayInvalidMaxMin       = errorer.Errorf("maxContains must be greater than minContains")
	ErrArrayInvalidMaxMinv      = errorer.Errorfn(ErrArrayInvalidMaxMin)
	ErrArrayItemsEmpty          = errorer.Errorf("array items is empty")
	ErrNumberInvalidMaxMin      = errorer.Errorf("maximum must be greater than minimum")
	ErrNumberInvalidMaxMinv     = errorer.Errorfn(ErrNumberInvalidMaxMin)
	ErrStringMaxLength          = errorer.Errorf("maxLength must be greater than 0")
	ErrStringMaxLengthv         = errorer.Errorfn(ErrStringMaxLength)
	ErrStringMinLength          = errorer.Errorf("minLength must be greater than 0")
	ErrStringMinLengthv         = errorer.Errorfn(ErrStringMinLength)
	ErrStringMaxMinLength       = errorer.Errorf("maxLength must be greater than minLength")
	ErrStringMaxMinLengthv      = errorer.Errorfn(ErrStringMaxMinLength)
	ErrObjectMaxProperties      = errorer.Errorf("maxProperties must be greater than 0")
	ErrObjectMaxPropertiesv     = errorer.Errorfn(ErrObjectMaxProperties)
	ErrObjectMinProperties      = errorer.Errorf("minProperties must be greater than 0")
	ErrObjectMinPropertiesv     = errorer.Errorfn(ErrObjectMinProperties)
	ErrObjectMaxMinProperties   = errorer.Errorf("maxProperties must be greater than minProperties")
	ErrObjectMaxMinPropertiesv  = errorer.Errorfn(ErrObjectMaxMinProperties)
	ErrObjectPropertiesMissing  = errorer.Errorf("properties must be greater than 0")
)
