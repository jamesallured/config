package feeder

import (
	"io/ioutil"
	"path/filepath"
	"strings"
)

// TomlDirectory is a feeder that feeds using a directory of TOML files.
type TomlDirectory struct {
	Path string
}

// Feed returns all the content.
func (td TomlDirectory) Feed() (map[string]interface{}, error) {
	files, err := ioutil.ReadDir(td.Path)
	if err != nil {
		return nil, err
	}

	all := map[string]interface{}{}

	for _, f := range files {
		if f.IsDir() {
			continue
		}

		j := Toml{Path: filepath.Join(td.Path, f.Name())}

		items, err := j.Feed()
		if err != nil {
			return nil, err
		}

		k := strings.Split(f.Name(), ".")[0]
		all[k] = items
	}

	return all, nil
}
