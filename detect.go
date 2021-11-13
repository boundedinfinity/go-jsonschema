package jsonschema

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func (t *System) readFile(uri string, bs *[]byte) error {
	p := uri
	p = strings.TrimPrefix(p, "file://")

	bs2, err := ioutil.ReadFile(p)

	if err != nil {
		return err
	}

	*bs = bs2

	return nil
}

func (t *System) readHttp(uri string, bs *[]byte) error {
	return fmt.Errorf("TODO")
}
