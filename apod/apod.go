// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package apod

import (
	"encoding/json"
	"fmt"
	"image"
	"net/http"
	"time"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
)

const apiURL = "https://api.data.gov/nasa/planetary/apod?api_key=%s&concept_tags=%v&date=%s"

// Response defines a single APOD data response.
type Response struct {
	Url         string `json:"url"`
	Title       string `json:"title"`
	Explanation string `json:"explanation"`
}

// Fetch fetches APOD data as per the given request properties.
//
//    @key:  The api.data.gov key for expanded usage.
//    @date: The date of the APOD image to retrieve.
//    @tags: Returns an ordered dictionary of concepts from the APOD explanation.
//           This is currently unused.
//
func Fetch(key string, date time.Time, tags bool) (Response, error) {
	url := fmt.Sprintf(apiURL, key, tags, date.Format("2006-01-02"))

	resp, err := http.Get(url)
	if err != nil {
		return Response{}, err
	}

	defer resp.Body.Close()

	var data Response
	dec := json.NewDecoder(resp.Body)
	return data, dec.Decode(&data)
}

// Image returns the APOD image as a decoded Go image, if possible.
func (r *Response) Image() (image.Image, error) {
	resp, err := http.Get(r.Url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	img, _, err := image.Decode(resp.Body)
	return img, err
}
