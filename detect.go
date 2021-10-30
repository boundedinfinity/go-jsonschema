package jsonschema

import (
	"fmt"
	"io/ioutil"
	"net/url"
	"path"
	"strings"

	"github.com/boundedinfinity/jsonschema/filetype"
	"github.com/boundedinfinity/jsonschema/uritype"
)

func (t *System) readContent(uri string, ut uritype.UriType, bs *[]byte) error {
	switch ut {
	case uritype.File:
		if err := t.readFile(uri, bs); err != nil {
			return err
		}
	case uritype.Http:
		if err := t.readHttp(uri, bs); err != nil {
			return err
		}
	default:
		return ErrUriTypeUnsupportedV(uri)
	}

	return nil
}

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

func (t *System) detectFileType(uri string, ft *filetype.FileType) error {
	ext := path.Ext(uri)

	switch ext {
	case ".yaml":
		*ft = filetype.Yaml
	case ".yml":
		*ft = filetype.Yaml
	case ".json":
		*ft = filetype.Json
	default:
		return ErrFileTypeUnsupportedV(uri)
	}

	return nil
}

func (t *System) detectUriType(uri string, ut *uritype.UriType) error {
	parsed, err := url.Parse(uri)

	if err != nil {
		return err
	}

	if parsed.Scheme == "http" || parsed.Scheme == "https" {
		*ut = uritype.Http
	}

	if parsed.Scheme == "file" {
		*ut = uritype.File
	}

	return nil
}
