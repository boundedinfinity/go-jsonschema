package idiomatic_test

import (
	"testing"

	"github.com/boundedinfinity/go-jsonschema/idiomatic"
	"github.com/boundedinfinity/go-mimetyper/mime_type"
	"github.com/stretchr/testify/assert"
)

var (
	yaml_string_1 = `
$id: https://www.boundedinfinity.com/test/string1
$schema: https://json-schema.org/draft/2020-12/schema
type: string
`
)

func Test_Unmarshal_1(t *testing.T) {
	m := idiomatic.NewMarshaler()

	actual, err := m.Unmarshal(mime_type.ApplicationYaml, []byte(yaml_string_1))

	assert.Nil(t, err)
	assert.Equal(t, nil, actual)
}
