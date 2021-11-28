package stringformat

import (
	"encoding/json"
	"fmt"
)

// StringFormatOption contains a initialized or empty copy of the StringFormat type.
type StringFormatOption struct {
	v *StringFormat
}

// NewStringFormatValue creates a new StringFormatOption type with an initialized value.
func NewStringFormatValue(v StringFormat) StringFormatOption {
	return StringFormatOption{
		v: &v,
	}
}

// NewStringFormatValue creates a new StringFormatOption type with the given value.
func NewStringFormatPtr(v *StringFormat) StringFormatOption {
	return StringFormatOption{
		v: v,
	}
}

// StringFormat2Ptr create a pointer from the given type.
func StringFormat2Ptr(v StringFormat) *StringFormat {
	return &v
}

// NewStringFormatValue creates a new StringFormatOption type with an empty value.
func NewStringFormatEmpty() StringFormatOption {
	return StringFormatOption{}
}

// NewStringFormatEmptyIfZeroValue creates a new initialized StringFormatOption type if the input StringFormat doesn't equal the StringFormat's zero value, or an empty StringFormatOption otherwise.
func NewStringFormatEmptyIfZeroValue(v StringFormat) StringFormatOption {
	var e StringFormat

	if v == e {
		return NewStringFormatEmpty()
	}

	return NewStringFormatValue(v)
}

//IsEmpty returns true if the contained StringFormat value is empty, false otherwise.
func (t StringFormatOption) IsEmpty() bool {
	return t.v == nil
}

//IsDefined returns true if the contained StringFormat value is initialized, false otherwise.
func (t StringFormatOption) IsDefined() bool {
	return !t.IsEmpty()
}

//Get returns the contained StringFormat value.
//NOTE: If the value is empty, this will return the StringFormat zero value.
func (t StringFormatOption) Get() StringFormat {
	var v StringFormat

	if t.IsEmpty() {
		return v
	}

	return *t.v
}

//Get returns the contained StringFormat value if the contained value is initialized or the input value is the value is not initialized.
func (t StringFormatOption) GetOrElse(v StringFormat) StringFormat {
	if t.IsDefined() {
		return t.Get()
	}

	return v
}

//MarshalJSON marshals the StringFormatOption type into the JSON representation.
func (t StringFormatOption) MarshalJSON() ([]byte, error) {
	if t.IsDefined() {
		return json.Marshal(*t.v)
	}

	return json.Marshal(nil)
}

//UnmarshalJSON unmarshals the JSON representation to the StringFormatOption type.
func (t *StringFormatOption) UnmarshalJSON(data []byte) error {
	if data == nil {
		return nil
	}

	if string(data) == "null" {
		return nil
	}

	var v StringFormat

	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	t.v = &v

	return nil
}

//MarshalYAML marshals the StringFormatOption type into the YAML representation.
func (t StringFormatOption) MarshalYAML() (interface{}, error) {
	if t.IsDefined() {
		return *t.v, nil
	}

	return nil, nil
}

//UnmarshalYAML unmarshals the YAML representation to the StringFormatOption type.
func (t *StringFormatOption) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var v StringFormat

	if err := unmarshal(&v); err != nil {
		return err
	}

	t.v = &v

	return nil
}

func (t *StringFormatOption) String() string {
	return fmt.Sprintf("%v", t.Get())
}
