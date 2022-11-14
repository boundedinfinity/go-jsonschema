package testdata

import (
	"fmt"
	"os"
	"path/filepath"
)

// func GetTestDataPaths(paths ...string) []string {
// 	var res []string

// 	for _, path := range paths {
// 		res = append(res, GetTestDataPaths(path))
// 	}

// 	return res
// }

func GetTestDataPath(path string) string {
	url, err := os.Getwd()

	if err != nil {
		panic(err)
	}

	url = filepath.Join(url, "testdata", path)
	url = fmt.Sprintf("file://%v", url)

	return url
}

func GetTestDataKey(filename string, index int) string {
	return fmt.Sprintf("%v:%v", GetTestDataPath(filename), index)
}
