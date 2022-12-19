package jsonschema

import (
	"errors"

	"github.com/boundedinfinity/go-jsonschema/model"
)

var ErrWalkExit = errors.New("walk exit")

type Walker struct {
	baseFn     func(*model.JsonSchemaBase) error
	arrayFn    func(*model.JsonSchemaArray) error
	itemsFn    func(*model.JsonSchemaArray, model.JsonSchema) error
	stringFn   func(*model.JsonSchemaString) error
	objectFn   func(*model.JsonSchemaObject) error
	propertyFn func(*model.JsonSchemaObject, string, model.JsonSchema) error
	refFn      func(*model.JsonSchemaRef) error
	integerFn  func(*model.JsonSchemaInteger) error
	numberFn   func(*model.JsonSchemaNumber) error
	nullFn     func(*model.JsonSchemaNull) error
	booleanFn  func(*model.JsonSchemaBoolean) error
}

func Walk() *Walker {
	return &Walker{}
}

func (w *Walker) Base(fn func(*model.JsonSchemaBase) error) *Walker {
	w.baseFn = fn
	return w
}

func (w *Walker) Array(fn func(*model.JsonSchemaArray) error) *Walker {
	w.arrayFn = fn
	return w
}

func (w *Walker) Items(fn func(*model.JsonSchemaArray, model.JsonSchema) error) *Walker {
	w.itemsFn = fn
	return w
}

func (w *Walker) String(fn func(*model.JsonSchemaString) error) *Walker {
	w.stringFn = fn
	return w
}

func (w *Walker) Object(fn func(*model.JsonSchemaObject) error) *Walker {
	w.objectFn = fn
	return w
}

func (w *Walker) Property(fn func(*model.JsonSchemaObject, string, model.JsonSchema) error) *Walker {
	w.propertyFn = fn
	return w
}

func (w *Walker) Ref(fn func(*model.JsonSchemaRef) error) *Walker {
	w.refFn = fn
	return w
}

func (w *Walker) Integer(fn func(*model.JsonSchemaInteger) error) *Walker {
	w.integerFn = fn
	return w
}

func (w *Walker) Number(fn func(*model.JsonSchemaNumber) error) *Walker {
	w.numberFn = fn
	return w
}

func (w *Walker) Null(fn func(*model.JsonSchemaNull) error) *Walker {
	w.nullFn = fn
	return w
}

func (w *Walker) Boolean(fn func(*model.JsonSchemaBoolean) error) *Walker {
	w.booleanFn = fn
	return w
}

func (w *Walker) Run(schema model.JsonSchema) error {
	if schema == nil {
		return nil
	}

	if schema.Base() != nil && w.baseFn != nil {
		if err := w.baseFn(schema.Base()); err != nil {
			return handleErr(err)
		}
	}

	switch js := schema.(type) {
	case *model.JsonSchemaString:
		if w.stringFn != nil {
			if err := w.stringFn(js); err != nil {
				handleErr(err)
			}
		}
	case *model.JsonSchemaArray:
		if w.arrayFn != nil {
			if err := w.arrayFn(js); err != nil {
				handleErr(err)
			}
		}

		if js.Items.Defined() {
			if w.itemsFn != nil {
				if err := w.itemsFn(js, js.Items.Get()); err != nil {
					return err
				}
			}
		}
	case *model.JsonSchemaObject:
		if w.objectFn != nil {
			if err := w.objectFn(js); err != nil {
				return handleErr(err)
			}

			if len(js.Properties) > 0 {
				for name, property := range js.Properties {
					if w.propertyFn != nil {
						if err := w.propertyFn(js, name, property); err != nil {
							return handleErr(err)
						}
					}
				}
			}
		}
	case *model.JsonSchemaRef:
		if w.refFn != nil {
			if err := w.refFn(js); err != nil {
				return handleErr(err)
			}
		}
	case *model.JsonSchemaInteger:
		if w.integerFn != nil {
			if err := w.integerFn(js); err != nil {
				return handleErr(err)
			}
		}
	case *model.JsonSchemaNumber:
		if w.numberFn != nil {
			if err := w.numberFn(js); err != nil {
				return handleErr(err)
			}
		}
	case *model.JsonSchemaNull:
		if w.nullFn != nil {
			if err := w.nullFn(js); err != nil {
				return handleErr(err)
			}
		}
	case *model.JsonSchemaBoolean:
		if w.booleanFn != nil {
			if err := w.booleanFn(js); err != nil {
				return handleErr(err)
			}
		}
	}

	return nil
}

func handleErr(err error) error {
	if err == nil {
		return nil
	}

	if errors.Is(err, ErrWalkExit) {
		return nil
	}

	return err
}
