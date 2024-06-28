package json_schema

// func NewRef(ref string) *JsonSchemaRef {
// 	schema := &JsonSchemaRef{
// 		JsonSchemaCommon: JsonSchemaCommon{
// 			Schema: model.SCHEMA_VERSION_2020_12,
// 			Ref:    ref,
// 		},
// 	}

// 	return schema
// }

// type JsonSchemaRef struct {
// 	JsonSchemaCommon
// 	Schema JsonSchema
// }

// var _ JsonSchema = &JsonSchemaRef{}

// func (t JsonSchemaRef) TypeName() string {
// 	return ""
// }

// func (t *JsonSchemaRef) Common() *JsonSchemaCommon {
// 	return &t.JsonSchemaCommon
// }

// func (t JsonSchemaRef) Copy() JsonSchema {
// 	return &JsonSchemaRef{
// 		JsonSchemaCommon: t.Common().Copy(),
// 		Schema:           t.Copy(),
// 	}
// }

// func (t JsonSchemaRef) Validate() error {
// 	if err := t.Common().Validate(); err != nil {
// 		return nil
// 	}

// 	return nil
// }
