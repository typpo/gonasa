// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package starapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// Star defines a single star.
type Star struct {
	Id                int64   `json:"id"`
	Label             string  `json:"label"`
	CreatedAt         string  `json:"created_at"`
	UpdatedAt         string  `json:"updated_at"`
	X                 float64 `json:"x"`
	Y                 float64 `json:"y"`
	Z                 float64 `json:"z"`
	Luminosity        float64 `json:"lum"`
	ColorBv           float64 `json:"colorb_v"`
	MagnitudeAbsolute float64 `json:"absmag"`
	MagnitudeApparent float64 `json:"appmag"`
	TexNum            float64 `json:"texnum"`
	Distance          float64 `json:"distly"` // Lightyears
	DCalc             float64 `json:"dcalc"`
	Plx               float64 `json:"plx"`
	PlxErr            float64 `json:"plxerr"`
	VX                float64 `json:"vx"`
	VY                float64 `json:"vy"`
	VZ                float64 `json:"vz"`
	Speed             float64 `json:"speed"`
	HipNum            float64 `json:"hipnum"`
}

// Stars returns a list of all stars in the database.
func Stars() ([]Star, error) {
	url := fmt.Sprintf("%s/stars", apiURL)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var data []Star
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&data)
	return data, err
}

// StarByName returns a single star for the given name.
// Returns nil if the star could not be found.
func StarByName(name string) (*Star, error) {
	url := fmt.Sprintf("%s/stars/%s", apiURL, url.QueryEscape(name))
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var data Star
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&data)
	return &data, err
}
