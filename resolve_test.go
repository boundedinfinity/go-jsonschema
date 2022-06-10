package jsonschema_test

import (
	"testing"

	"github.com/boundedinfinity/jsonschema"
	"github.com/boundedinfinity/optioner"
	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	schema := &jsonschema.JsonSchmea{
		Id: optioner.Some("test-id"),
	}
	assert.Equal(t, schema.Id.Defined(), true)
}
