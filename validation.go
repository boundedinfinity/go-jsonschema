package jsonschema

import (
	"regexp"

	"github.com/boundedinfinity/jsonschema/objecttype"
	"github.com/boundedinfinity/jsonschema/stringformat"
	"github.com/boundedinfinity/optioner"
)

func (t JsonSchmea) Validate() error {
	if t.Type.IsEmpty() {
		return ErrObjectTypeEmpty
	}

	switch t.Type.Get() {
	case objecttype.Array:
		if t.Items == nil {
			return ErrInvalidArrayMissingItems
		}

		if err := validateNonNegative(t.MinItems); err != nil {
			return err
		}

		if err := validateNonNegative(t.MaxItems); err != nil {
			return err
		}

		if err := validateNonNegative(t.MaxContains); err != nil {
			return err
		}

		if err := t.Items.Validate(); err != nil {
			return err
		}
	case objecttype.Integer:
		if err := t.validateInteger(); err != nil {
			return err
		}

	case objecttype.Number:
		if t.MultipleOf.IsDefined() && t.MultipleOf.Get() <= 0 {
			return ErrInvalidMultipleOf
		}
	case objecttype.String:
		if err := t.validateString(); err != nil {
			return err
		}

	case objecttype.Object:
		for _, property := range t.Properties {
			if err := property.Validate(); err != nil {
				return err
			}
		}
	}

	return nil
}

func (t JsonSchmea) validateInteger() error {
	if t.MultipleOf.IsDefined() {
		if t.MultipleOf.Get() <= 0 {
			return ErrInvalidMultipleOf
		}

		if err := validateInteger(t.MultipleOf); err != nil {
			return err
		}
	}

	if t.Maximum.IsDefined() {
		if err := validateInteger(t.Maximum); err != nil {
			return err
		}
	}

	if t.ExclusiveMaximum.IsDefined() {
		if err := validateInteger(t.ExclusiveMaximum); err != nil {
			return err
		}
	}

	if t.Minimum.IsDefined() {
		if err := validateInteger(t.Minimum); err != nil {
			return err
		}
	}

	if t.ExclusiveMinimum.IsDefined() {
		if err := validateInteger(t.ExclusiveMinimum); err != nil {
			return err
		}
	}
	return nil
}

func (t JsonSchmea) validateString() error {
	if err := validateNonNegative(t.MaxLength); err != nil {
		return err
	}

	if err := validateNonNegative(t.MinLength); err != nil {
		return err
	}

	if err := validateNotGreaterThan(t.MinLength, t.MaxLength); err != nil {
		return err
	}

	if t.Pattern.IsDefined() {
		_, err := regexp.Compile(t.Pattern.Get())

		if err != nil {
			return ErrStringInvalidPatternV(t.Pattern.Get(), err)
		}
	}

	if t.Format.IsDefined() {
		switch t.Format.Get() {
		case stringformat.DateTime, stringformat.Date, stringformat.Time, stringformat.Duration,
			stringformat.Email, stringformat.IdnEmail,
			stringformat.Hostname, stringformat.IdnHostname,
			stringformat.Ipv4, stringformat.Ipv6,
			stringformat.Uuid,
			stringformat.Uri, stringformat.UriReference,
			stringformat.Iri, stringformat.IriReference,
			stringformat.UriTemplate,
			stringformat.JsonPointer, stringformat.RelativeJsonPointer:

		case stringformat.Regex:
			return ErrStringFormatUnsupportedV(t.Format)
		default:
			return ErrStringFormatInvalidV(t.Format)
		}
	}

	return nil
}

func validateNotGreaterThan(x, y optioner.IntOption) error {
	if x.IsEmpty() || y.IsEmpty() {
		return nil
	}

	if x.Get() > y.Get() {
		return ErrMinGreaterThanMaxV(x.Get(), y.Get())
	}
	return nil
}

func validateInteger(v optioner.Float64Option) error {
	if v.IsEmpty() {
		return nil
	}

	x := float64(int64(v.Get()))

	if x != v.Get() {
		return ErrNotInteger
	}

	return nil
}

func validateNonNegative(v optioner.IntOption) error {
	if v.IsEmpty() {
		return nil
	}

	if v.Get() < 0 {
		return ErrNotNonNegative
	}

	return nil
}
