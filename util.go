package jsonschema

import (
	"io/ioutil"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strings"
)

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

func (t *System) uris2files(uris ...string) ([]string, error) {
	var files []string

	for _, uri := range uris {
		var uriType JsonSchemaUriType

		if err := t.detectUriType(uri, &uriType); err != nil {
			return files, err
		}

		switch uriType {
		case JsonSchemaUriType_File:
			p := strings.TrimPrefix(uri, "file://")
			stat, err := os.Stat(p)

			if err != nil {
				return files, err
			}

			if stat.IsDir() {
				infos, err := ioutil.ReadDir(p)

				if err != nil {
					return files, err
				}

				for _, info := range infos {
					file := filepath.Join(p, info.Name())
					file = "file://" + file
					files = append(files, file)
				}
			} else {
				files = append(files, uri)
			}
		case JsonSchemaUriType_Http:
			files = append(files, uri)
		}

	}

	return files, nil
}
