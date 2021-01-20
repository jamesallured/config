package feeder

import (
	"os"
	"path/filepath"

	"github.com/pelletier/go-toml"
)

// Toml is a feeder that feeds using a single TOML file.
type Toml struct {
	Path string
}

// Feed returns all of the content contained by the TOML file.
func (t *Toml) Feed() (map[string]interface{}, error) {
	fl, err := os.Open(filepath.Clean(t.Path))
	if err != nil {
		return nil, err
	}
	defer fl.Close()

	items := make(map[string]interface{})

	if err := toml.NewDecoder(fl).Decode(&items); err != nil {
		return nil, err
	}
	return items, nil
}
