// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

package asterank

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	list, err := Fetch(
		10,
		GTE("profit", 10e9),
	)

	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(len(list))

	for _, a := range list {
		fmt.Println(a["full_name"], a["price"])
	}
}
