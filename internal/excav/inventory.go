package excav

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type Inventory struct {
	// don't use this attribute directly. It's public only because HCL parsing
	XRepositories []*Repository `yaml:"repositories"`
}

//------------------------------------------------------------------------------
// Inventory public API
//------------------------------------------------------------------------------

// OpenInventory load the inventory YAML file. If file is empty string,
// the default value will be used which is '{currentdir}/inventory.yaml'.
//
// Example of the file is in testdata/inventory.yaml
func OpenInventory(file string) (*Inventory, error) {

	if file == "" {
		currenDir, _ := os.Getwd()
		file = filepath.Join(currenDir, "inventory.yaml")
	}

	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, fmt.Errorf("cannot read file %s", file)
	}

	repos := make([]*Repository, 0)
	err = yaml.Unmarshal(data, &repos)
	if err != nil {
		return nil, err
	}

	return &Inventory{XRepositories: repos}, nil
}

// GetByPath returns repo or null
func (inv *Inventory) Get(repoName string) *Repository {
	for _, repo := range inv.XRepositories {
		if repo.Name == repoName {
			return repo
		}
	}
	return nil
}

func (inv *Inventory) GetAll() []*Repository {
	return inv.XRepositories
}

// GetByTags returns you list of filtered repositories they're matching
// given tag or all given tags. If tags is empty, the function returns all
// tags
func (inv *Inventory) GetByTags(tags ...string) []*Repository {
	if len(tags) == 0 {
		return inv.XRepositories
	}

	filtered := []*Repository{}
	for _, repo := range inv.XRepositories {
		if repo.HasTags(tags...) {
			filtered = append(filtered, repo)
		}
	}
	return filtered
}
