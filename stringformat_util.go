package jsonschema

import (
	"regexp"

	"github.com/boundedinfinity/jsonschema/stringformat"
)

var (
	STRINGFORMAT_REGEX_DATE                = ".*"
	STRINGFORMAT_REGEX_TIME                = ".*"
	STRINGFORMAT_REGEX_DATETIME            = ".*"
	STRINGFORMAT_REGEX_DURATION            = ".*"
	STRINGFORMAT_REGEX_EMAIL               = ".*"
	STRINGFORMAT_REGEX_IDNEMAIL            = ".*"
	STRINGFORMAT_REGEX_HOSTNAME            = ".*"
	STRINGFORMAT_REGEX_IDNHOSTNAME         = ".*"
	STRINGFORMAT_REGEX_IPV4                = ".*"
	STRINGFORMAT_REGEX_IPV6                = ".*"
	STRINGFORMAT_REGEX_UUID                = ".*"
	STRINGFORMAT_REGEX_URI                 = ".*"
	STRINGFORMAT_REGEX_URIREFERENCE        = ".*"
	STRINGFORMAT_REGEX_IRI                 = ".*"
	STRINGFORMAT_REGEX_IRIREFERENCE        = ".*"
	STRINGFORMAT_REGEX_URLTEMPLATE         = ".*"
	STRINGFORMAT_REGEX_JSONPOINTER         = ".*"
	STRINGFORMAT_REGEX_RELATIVEJSONPOINTER = ".*"

	STRINGFORMAT_REGEX = map[stringformat.StringFormat]string{
		stringformat.Date:                STRINGFORMAT_REGEX_DATE,
		stringformat.DateTime:            STRINGFORMAT_REGEX_DATETIME,
		stringformat.Time:                STRINGFORMAT_REGEX_TIME,
		stringformat.Duration:            STRINGFORMAT_REGEX_DURATION,
		stringformat.Email:               STRINGFORMAT_REGEX_DURATION,
		stringformat.IdnEmail:            STRINGFORMAT_REGEX_DURATION,
		stringformat.Hostname:            STRINGFORMAT_REGEX_HOSTNAME,
		stringformat.IdnHostname:         STRINGFORMAT_REGEX_IDNHOSTNAME,
		stringformat.Ipv4:                STRINGFORMAT_REGEX_IPV4,
		stringformat.Uuid:                STRINGFORMAT_REGEX_UUID,
		stringformat.Uri:                 STRINGFORMAT_REGEX_URI,
		stringformat.UriReference:        STRINGFORMAT_REGEX_URIREFERENCE,
		stringformat.Iri:                 STRINGFORMAT_REGEX_IRI,
		stringformat.IriReference:        STRINGFORMAT_REGEX_IRIREFERENCE,
		stringformat.UriTemplate:         STRINGFORMAT_REGEX_URLTEMPLATE,
		stringformat.JsonPointer:         STRINGFORMAT_REGEX_JSONPOINTER,
		stringformat.RelativeJsonPointer: STRINGFORMAT_REGEX_RELATIVEJSONPOINTER,
	}
)

func init() {
	// These should never fail
	for _, v := range STRINGFORMAT_REGEX {
		regexp.MustCompile(v)
	}
}
