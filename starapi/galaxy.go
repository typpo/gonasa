// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package starapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// GalaxyGroup defines a single galaxy.
type GalaxyGroup struct {
	Id        int64   `json:"id"`
	Label     string  `json:"label"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
	Distance  string  `json:"distly"`
	X         float64 `json:"x"`
	Y         float64 `json:"y"`
	Z         float64 `json:"z"`
}

// LocalGroups returns a list of all local groups of galaxies.
func LocalGroups() ([]GalaxyGroup, error) {
	url := fmt.Sprintf("%s/local_groups", apiURL)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var data []GalaxyGroup
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&data)
	return data, err
}
