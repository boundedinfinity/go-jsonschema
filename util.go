package jsonschema

import (
	"net/url"
	"path"

	"github.com/boundedinfinity/mimetyper/file_extention"
)

func mkExt(ft file_extention.FileExtention) string {
	return ".schema" + ft.String()
}

func urlJoin(root string, paths ...string) (string, error) {
	u, err := url.Parse(root)

	if err != nil {
		return "", err
	}

	components := []string{u.Path}
	components = append(components, paths...)

	u.Path = path.Join(components...)

	return u.String(), nil
}
