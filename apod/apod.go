// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package apod

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const apiURL = "https://api.data.gov/nasa/planetary/apod"

// Request defines a single APOD data request.
type Request struct {
	Date        time.Time // The date date of the APOD image to retrieve.
	Key         string    // The api.data.gov key for expanded usage.
	ConceptTags bool      // Return an ordered dictionary of concepts from the APOD explanation.
}

// Response defines a single APOD data response.
type Response struct {
	Url         string `json:"url"`
	Title       string `json:"title"`
	Explanation string `json:"explanation"`
}

// Fetch fetches APOD data as per the given request.
func Fetch(req *Request) (*Response, error) {
	url := fmt.Sprintf("%s?api_key=%s&concept_tags=%v&date=%s",
		apiURL, req.Key, req.ConceptTags, req.Date.Format("2006-01-02"))

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var data Response
	dec := json.NewDecoder(resp.Body)
	return &data, dec.Decode(&data)
}
