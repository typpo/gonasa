// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package asterank

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

const apiURL = "http://asterank.com/api/asterank?query={%s}&limit=%d"

// Fetch performs an data query and returns at most limit results
// for the given set of Parameter filters.
func Fetch(limit int, query ...Parameter) ([]Asteroid, error) {
	if len(query) == 0 {
		return nil, nil
	}

	qstr := make([]string, len(query))
	for i, f := range query {
		qstr[i] = f.String()
	}

	url := fmt.Sprintf(apiURL, strings.Join(qstr, ","), limit)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var data []Asteroid
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&data)
	return data, err
}
