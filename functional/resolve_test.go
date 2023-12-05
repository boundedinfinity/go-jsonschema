package functional_test

// import (
// 	"testing"

// 	"github.com/boundedinfinity/go-jsonschema"
// 	"github.com/stretchr/testify/assert"
// )

// func Test_Load_object_1(t *testing.T) {
// 	sys := jsonschema.New()

// 	err := sys.Load(getTestDataPath("strings"), getTestDataPath("objects/obj-1.schema.yaml"))
// 	assert.Nil(t, err)

// 	err = sys.Resolve()
// 	assert.Nil(t, err)

// 	opt := sys.ById("https://www.boundedinfinity.com/schema/test/object-1")
// 	assert.True(t, opt.Defined())
// 	assert.True(t, opt.Get().Type.Defined())
// 	assert.Equal(t, opt.Get().Type.Get(), objecttype.Object)
// }

// func Test_Load_object_2(t *testing.T) {
// 	sys := jsonschema.New()

// 	err := sys.Load(getTestDataPath("strings"), getTestDataPath("objects/obj-2.schema.yaml"))
// 	assert.Nil(t, err)

// 	err = sys.Resolve()
// 	assert.Nil(t, err)
// }

// func Test_Load_object_3(t *testing.T) {
// 	sys := jsonschema.New()

// 	err := sys.Load(getTestDataPath("strings"), getTestDataPath("objects/obj-3.schema.yaml"))
// 	assert.Nil(t, err)

// 	err = sys.Resolve()
// 	assert.Nil(t, err)
// }
