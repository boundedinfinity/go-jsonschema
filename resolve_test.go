package jsonschema_test

import (
	"testing"

	"github.com/boundedinfinity/jsonschema"
	"github.com/stretchr/testify/assert"
)

func Test_Load_object_1(t *testing.T) {
	sys := jsonschema.New()

	err := sys.Load(getTestDataPath("strings"), getTestDataPath("objects"))
	assert.Nil(t, err)

	err = sys.Resolve()
	assert.Nil(t, err)
}
