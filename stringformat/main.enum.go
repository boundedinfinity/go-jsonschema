//************************************************************************************
//*                                                                                  *
//* ===== DO NOT EDIT =====                                                          *
//* Any change will be overwritten                                                   *
//* Generated by github.com/boundedinfinity/enumer                                   *
//*                                                                                  *
//************************************************************************************

package stringformat

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/boundedinfinity/go-commoner/slicer" // v1.0.15
)

var (
	All = []StringFormat{
		Date,
		DateTime,
		Duration,
		Email,
		Hostname,
		IdnEmail,
		IdnHostname,
		Ipv4,
		Ipv6,
		Iri,
		IriReference,
		JsonPointer,
		Regex,
		RelativeJsonPointer,
		Time,
		Uri,
		UriReference,
		UriTemplate,
		Uuid,
	}
)

func (t StringFormat) String() string {
	return string(t)
}

func Parse(v string) (StringFormat, error) {
	f, ok := slicer.FindFn(All, func(x StringFormat) bool {
		return StringFormat(v) == x
	})

	if !ok {
		return f, ErrorV(v)
	}

	return f, nil
}

func Is(s string) bool {
	return slicer.ContainsFn(All, func(v StringFormat) bool {
		return string(v) == s
	})
}

var ErrInvalid = errors.New("invalid enumeration type")

func ErrorV(v string) error {
	return fmt.Errorf(
		"%w '%v', must be one of %v",
		ErrInvalid, v, slicer.Join(All, ","),
	)
}

func (t StringFormat) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(t))
}

func (t *StringFormat) UnmarshalJSON(data []byte) error {
	var v string

	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	e, err := Parse(v)

	if err != nil {
		return err
	}

	*t = e

	return nil
}

func (t StringFormat) MarshalYAML() (interface{}, error) {
	return string(t), nil
}

func (t *StringFormat) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var v string

	if err := unmarshal(&v); err != nil {
		return err
	}

	e, err := Parse(v)

	if err != nil {
		return err
	}

	*t = e

	return nil
}
