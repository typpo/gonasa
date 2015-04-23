// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package starapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Planet defines a single (exo)planet.
type Planet struct {
	Id         int64   `json:"id"`
	NumPlanets int64   `json:"numplanets"`
	Texture    int64   `json:"texture"`
	Label      string  `json:"label"`
	CreatedAt  string  `json:"created_at"`
	UpdatedAt  string  `json:"updated_at"`
	Distance   float64 `json:"distance"`
	X          float64 `json:"x"`
	Y          float64 `json:"y"`
	Z          float64 `json:"z"`
}

// Exoplanets returns a list of all exoplanets.
func Exoplanets() ([]Planet, error) {
	url := fmt.Sprintf("%s/exo_planets", apiURL)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var data []Planet
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&data)
	return data, err
}
