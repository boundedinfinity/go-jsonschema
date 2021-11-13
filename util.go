package jsonschema

import (
	"net/url"
	"path"
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
