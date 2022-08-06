package jsonschema_test

// func getTestDataPath(filename string) string {
// 	dir, err := os.Getwd()

// 	if err != nil {
// 		panic(err)
// 	}

// 	path := filepath.Join(dir, "testdata")

// 	if filename != "" {
// 		path = filepath.Join(path, filename)
// 	}

// 	return path
// }

// func getTestDataKey(filename string, index int) string {
// 	return fmt.Sprintf("%v:%v", getTestDataPath(filename), index)
// }

// func Test_Load_string_dir(t *testing.T) {
// 	path := getTestDataPath("strings")
// 	sys := jsonschema.New()
// 	err := sys.Load(path)

// 	assert.Nil(t, err)
// 	assert.Equal(t, 4, len(sys.SourceMap))
// }

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
