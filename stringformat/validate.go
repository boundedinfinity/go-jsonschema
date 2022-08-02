package stringformat

import (
	"fmt"
	"regexp"
)

var (
	regex_1_to_12      = `([0]?[0-9]|1[0-2])`
	regex_1_to_24      = `([0]?[0-9]|1[0-9]|2[0-4])`
	regex_1_to_31      = `([0]?[0-9]|[1-2][0-9]|3[0-1])`
	regex_0_to_59      = `([0]?[0-9]|[1-4][0-9]|5[0-9])`
	regex_0000_to_9999 = `\d{4}`

	regex_dur_sec    = `\d+[sS]`
	regex_dur_minute = `\d+[mM]`
	regex_dur_hour   = `\d+[hH]`
	regex_dur_time   = fmt.Sprintf(`T%v?%v?%v?`, regex_dur_hour, regex_dur_minute, regex_dur_sec)

	regex_dur_day   = `\d+[dD]`
	regex_dur_week  = `\d+[wW]`
	regex_dur_month = `\d+[mM]`
	regex_dur_year  = `\d+[yY]`
	regex_dur_date  = fmt.Sprintf(`%v?%v?`, regex_dur_year, regex_dur_month)

	STRINGFORMAT_REGEX_DATE     = regex_0000_to_9999 + "-" + regex_1_to_12 + "-" + regex_1_to_31
	STRINGFORMAT_REGEX_TIME     = regex_1_to_24 + ":" + regex_0_to_59 + ":" + regex_0_to_59
	STRINGFORMAT_REGEX_DATETIME = STRINGFORMAT_REGEX_DATE + "T" + STRINGFORMAT_REGEX_TIME
	STRINGFORMAT_REGEX_DURATION = ".*"
	// https://ihateregex.io/expr/email/
	STRINGFORMAT_REGEX_EMAIL       = `[^@ \t\r\n]+@[^@ \t\r\n]+\.[^@ \t\r\n]+`
	STRINGFORMAT_REGEX_IDNEMAIL    = ".*"
	STRINGFORMAT_REGEX_HOSTNAME    = ".*"
	STRINGFORMAT_REGEX_IDNHOSTNAME = ".*"
	// https://ihateregex.io/expr/ip/
	STRINGFORMAT_REGEX_IPV4 = `(\b25[0-5]|\b2[0-4][0-9]|\b[01]?[0-9][0-9]?)(\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){3}`
	// https://ihateregex.io/expr/ipv6/
	STRINGFORMAT_REGEX_IPV6 = `(([0-9a-fA-F]{1,4}:){7,7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:)|fe80:(:[0-9a-fA-F]{0,4}){0,4}%[0-9a-zA-Z]{1,}|::(ffff(:0{1,4}){0,1}:){0,1}((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])|([0-9a-fA-F]{1,4}:){1,4}:((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9]))`
	// https://ihateregex.io/expr/uuid/
	STRINGFORMAT_REGEX_UUID                = "^[0-9a-fA-F]{8}\b-[0-9a-fA-F]{4}\b-[0-9a-fA-F]{4}\b-[0-9a-fA-F]{4}\b-[0-9a-fA-F]{12}$"
	STRINGFORMAT_REGEX_URI                 = ".*"
	STRINGFORMAT_REGEX_URIREFERENCE        = ".*"
	STRINGFORMAT_REGEX_IRI                 = ".*"
	STRINGFORMAT_REGEX_IRIREFERENCE        = ".*"
	STRINGFORMAT_REGEX_URLTEMPLATE         = ".*"
	STRINGFORMAT_REGEX_JSONPOINTER         = ".*"
	STRINGFORMAT_REGEX_RELATIVEJSONPOINTER = ".*"

	STRINGFORMAT_REGEX = map[StringFormat]string{
		Date:                STRINGFORMAT_REGEX_DATE,
		DateTime:            STRINGFORMAT_REGEX_DATETIME,
		Time:                STRINGFORMAT_REGEX_TIME,
		Duration:            STRINGFORMAT_REGEX_DURATION,
		Email:               STRINGFORMAT_REGEX_DURATION,
		IdnEmail:            STRINGFORMAT_REGEX_DURATION,
		Hostname:            STRINGFORMAT_REGEX_HOSTNAME,
		IdnHostname:         STRINGFORMAT_REGEX_IDNHOSTNAME,
		Ipv4:                STRINGFORMAT_REGEX_IPV4,
		Uuid:                STRINGFORMAT_REGEX_UUID,
		Uri:                 STRINGFORMAT_REGEX_URI,
		UriReference:        STRINGFORMAT_REGEX_URIREFERENCE,
		Iri:                 STRINGFORMAT_REGEX_IRI,
		IriReference:        STRINGFORMAT_REGEX_IRIREFERENCE,
		UriTemplate:         STRINGFORMAT_REGEX_URLTEMPLATE,
		JsonPointer:         STRINGFORMAT_REGEX_JSONPOINTER,
		RelativeJsonPointer: STRINGFORMAT_REGEX_RELATIVEJSONPOINTER,
	}
)

func init() {
	for _, v := range STRINGFORMAT_REGEX {
		regexp.MustCompile(v)
	}
}
