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
	"fmt"
	"strings"
)

type StringFormat string
type StringFormats []StringFormat

const (
	Date         StringFormat = "date"
	DateTime     StringFormat = "date-time"
	Duration     StringFormat = "duration"
	Email        StringFormat = "email"
	Hostname     StringFormat = "hostname"
	IdnEmail     StringFormat = "idn-email"
	IdnHostname  StringFormat = "idn-hostname"
	Ipv4         StringFormat = "ipv4"
	Ipv6         StringFormat = "ipv6"
	Iri          StringFormat = "iri"
	IriReference StringFormat = "iri-reference"
	JsonPointer  StringFormat = "json-pointer"
	Regex        StringFormat = "regex"
	Time         StringFormat = "time"
	Uri          StringFormat = "uri"
	UriReference StringFormat = "uri-reference"
	UriTemplate  StringFormat = "uri-template"
	Uuid         StringFormat = "uuid"
)

var (
	All = StringFormats{
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
		Time,
		Uri,
		UriReference,
		UriTemplate,
		Uuid,
	}
)

func Is(v string) bool {
	return All.Is(v)
}

func Parse(v string) (StringFormat, error) {
	return All.Parse(v)
}

func Strings() []string {
	return All.Strings()
}

func (t StringFormat) String() string {
	return string(t)
}

func errStringFormatInvalid(vs StringFormats, v string) error {
	return fmt.Errorf(
		"invalid enumeration type '%v', must be one of %v",
		v, strings.Join(vs.Strings(), ","),
	)
}

func (t StringFormats) Strings() []string {
	var ss []string

	for _, v := range t {
		ss = append(ss, v.String())
	}

	return ss
}

func (t StringFormats) Parse(v string) (StringFormat, error) {
	var o StringFormat
	var f bool
	n := strings.ToLower(v)

	for _, e := range t {
		if strings.ToLower(e.String()) == n {
			o = e
			f = true
			break
		}
	}

	if !f {
		return o, errStringFormatInvalid(t, v)
	}

	return o, nil
}

func (t StringFormats) Is(v string) bool {
	var f bool

	for _, e := range t {
		if string(e) == v {
			f = true
			break
		}
	}

	return f
}

func (t StringFormats) Contains(v StringFormat) bool {
	for _, e := range t {
		if e == v {
			return true
		}
	}

	return false
}

func (t StringFormat) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(t))
}

func (t *StringFormat) UnmarshalJSON(data []byte) error {
	var s string

	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	e, err := Parse(s)

	if err != nil {
		return err
	}

	*t = e

	return nil
}
