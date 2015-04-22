// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package apod

import (
	"testing"
	"time"
)

func Test(t *testing.T) {
	date, err := time.Parse("2006-01-02", "2015-04-20")
	if err != nil {
		t.Fatal(err)
	}

	resp, err := Fetch("DEMO_KEY", date, true)
	if err != nil {
		t.Fatal(err)
	}

	_, err = resp.Image()
	if err != nil {
		t.Fatalf("expected valid image: %v", err)
	}
}
