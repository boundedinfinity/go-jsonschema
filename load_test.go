package jsonschema_test

import (
	"testing"

	"github.com/boundedinfinity/go-jsonschema"
	"github.com/boundedinfinity/go-jsonschema/model"
	"github.com/boundedinfinity/go-jsonschema/testdata"
	"github.com/stretchr/testify/assert"
)

func Test_Load_string_single(t *testing.T) {
	path := testdata.GetTestDataPath("strings/single.schema.yaml")
	sys := jsonschema.New()
	err := sys.LoadPath(path)

	assert.Nil(t, err)
}

func Test_Load_string_single_duplicate(t *testing.T) {
	path := testdata.GetTestDataPath("strings/single.schema.yaml")
	sys := jsonschema.New()
	err := sys.LoadPath(path, path)

	assert.NotNil(t, err)
	assert.ErrorIs(t, err, model.ErrPathDuplicate)
}

func Test_Load_string_ref(t *testing.T) {
	sys := jsonschema.New()
	err := sys.LoadPath(
		testdata.GetTestDataPath("strings/single.schema.yaml"),
		testdata.GetTestDataPath("strings/ref.schema.yaml"),
		testdata.GetTestDataPath("strings/array.schema.yaml"),
	)

	assert.Nil(t, err)

	err = sys.Check()
	assert.Nil(t, err)
}

func Test_Load_object_ref(t *testing.T) {
	sys := jsonschema.New()
	err := sys.LoadPath(
		testdata.GetTestDataPath("objects/obj-1.schema.yaml"),
		testdata.GetTestDataPath("objects/obj-2.schema.yaml"),
		testdata.GetTestDataPath("objects/obj-3.schema.yaml"),
		testdata.GetTestDataPath("strings/single.schema.yaml"),
	)

	assert.Nil(t, err)

	err = sys.Check()
	assert.Nil(t, err)
}

func Test_Load_object_bad_ref(t *testing.T) {
	sys := jsonschema.New()
	err := sys.LoadPath(
		testdata.GetTestDataPath("objects/obj-3.schema.yaml"),
	)

	assert.Nil(t, err)

	err = sys.Check()
	assert.ErrorIs(t, err, model.ErrRefNotFound)
}

func Test_Load_resolve(t *testing.T) {
	sys := jsonschema.New()
	err := sys.LoadPath(
		testdata.GetTestDataPath("strings/ref.schema.yaml"),
	)

	assert.Nil(t, err)

	err = sys.Check()
	assert.ErrorIs(t, err, model.ErrRefNotFound)
}
