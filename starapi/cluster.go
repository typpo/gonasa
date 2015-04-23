// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package starapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// OpenCluster defines a single open cluster
type OpenCluster struct {
	Id        int64   `json:"id"`
	Label     string  `json:"label"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
	Logage    float64 `json:"logage"`
	Metal     float64 `json:"metal"`
	Distance  float64 `json:"distly"`
	Diameter  float64 `json:"diam"`
	X         float64 `json:"x"`
	Y         float64 `json:"y"`
	Z         float64 `json:"z"`
}

// OpenClusters returns a list of all local groups of galaxies.
func OpenClusters() ([]OpenCluster, error) {
	url := fmt.Sprintf("%s/open_cluster", apiURL)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var data []OpenCluster
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&data)
	return data, err
}
