package stringformat

//go:generate enumer -path=./main.go

type StringFormat string

const (
	Date                StringFormat = "date"
	DateTime            StringFormat = "date-time"
	Duration            StringFormat = "duration"
	Email               StringFormat = "email"
	Hostname            StringFormat = "hostname"
	IdnEmail            StringFormat = "idn-email"
	IdnHostname         StringFormat = "idn-hostname"
	Ipv4                StringFormat = "ipv4"
	Ipv6                StringFormat = "ipv6"
	Iri                 StringFormat = "iri"
	IriReference        StringFormat = "iri-reference"
	JsonPointer         StringFormat = "json-pointer"
	Regex               StringFormat = "regex"
	RelativeJsonPointer StringFormat = "relative-json-pointer"
	Time                StringFormat = "time"
	Uri                 StringFormat = "uri"
	UriReference        StringFormat = "uri-reference"
	UriTemplate         StringFormat = "uri-template"
	Uuid                StringFormat = "uuid"
)
