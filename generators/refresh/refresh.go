package refresh

import (
	"github.com/gobuffalo/buffalo/generators"
	"github.com/markbates/gentronics"
)

// New generator for a .buffalo.dev.yml file
func New() (*gentronics.Generator, error) {
	g := gentronics.New()

	files, err := generators.Find("refresh")
	if err != nil {
		return nil, err
	}

	for _, f := range files {
		g.Add(gentronics.NewFile(f.WritePath, f.Body))
	}

	return g, nil
}