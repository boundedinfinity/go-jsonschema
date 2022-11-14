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
	err := sys.Load(path)

	assert.Nil(t, err)
}

func Test_Load_string_single_duplicate(t *testing.T) {
	path := testdata.GetTestDataPath("strings/single.schema.yaml")
	sys := jsonschema.New()
	err := sys.Load(path, path)

	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), model.ErrSchemaIdDuplicate.Error())
}

func Test_Load_string_ref(t *testing.T) {
	path1 := testdata.GetTestDataPath("strings/single.schema.yaml")
	path2 := testdata.GetTestDataPath("strings/ref.schema.yaml")
	sys := jsonschema.New()
	err := sys.Load(path1, path2)

	assert.NotNil(t, err)

	err = sys.Resolve()
	assert.NotNil(t, err)
}

// func Test_Load_string_plain(t *testing.T) {
// 	path := getTestDataPath("strings/plain.schema.yaml")
// 	sys := jsonschema.New()
// 	err := sys.Load(path)

// 	assert.Nil(t, err)
// 	assert.Equal(t, 1, len(sys.SourceMap))

// 	obj := sys.ById("https://www.boundedinfinity.com/schema/test/string-plain")

// 	assert.True(t, obj.Defined())

// 	assert.Equal(t, "https://www.boundedinfinity.com/schema/test/string-plain", obj.Get().Id.Get())
// 	assert.Equal(t, objecttype.String, obj.Get().Type.Get())
// }

// func Test_Load_string(t *testing.T) {
// 	sys := jsonschema.New()
// 	err := sys.Load(getTestDataPath("strings/single.schema.yaml"))

// 	assert.Nil(t, err)
// 	assert.Equal(t, 1, len(sys.SourceMap))

// 	sys.Clear()
// 	err = sys.Load(getTestDataPath("strings/multi.schema.yaml"))

// 	assert.Nil(t, err)
// 	assert.Equal(t, 2, len(sys.SourceMap))
// }

// func Test_Load_bad_extention(t *testing.T) {
// 	sys := jsonschema.New()
// 	err := sys.Load(getTestDataPath("strings/normal.yaml"))

// 	assert.NotNil(t, err)
// 	assert.True(t, errors.Is(err, jsonschema.ErrUnsupportedFileType))
// }
