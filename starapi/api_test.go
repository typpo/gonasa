// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package starapi

import "testing"

func TestStars(t *testing.T) {
	_, err := Stars()
	if err != nil {
		t.Fatal(err)
	}
}

func TestStar(t *testing.T) {
	_, err := StarByName("Sun")
	if err != nil {
		t.Fatal(err)
	}
}

func TestExoplanets(t *testing.T) {
	_, err := Exoplanets()
	if err != nil {
		t.Fatal(err)
	}
}

func TestLocalGroups(t *testing.T) {
	_, err := LocalGroups()
	if err != nil {
		t.Fatal(err)
	}
}

func TestOpenClusters(t *testing.T) {
	_, err := OpenClusters()
	if err != nil {
		t.Fatal(err)
	}
}
