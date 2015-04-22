// This file is subject to a 1-clause BSD license.
// Its contents can be found in the enclosed LICENSE file.

/*
Package asterank provides Go bindings for the NASA/JPL Asterank API.
Ref: http://www.asterank.com/api

The Asterank database is a thin layer over the NASA/JPL Small Body Database,
merged with JPL delta-v data, published asteroid mass data, and some custom
calculations.


Usage:

Find the top 10 asteroids with a roughly circular orbit, low inclination
and semi-major axis less than 1.5 AU:

Find the first 10 asteroids with a roughly circular orbit, low inclination
and semi-major axis less than 1.5 AU:

	list, err := asterank.Fetch(
		10,
		asterank.LT("e", 0.1),
		asterank.LT("i", 4),
		asterank.LT("a", 1.5),
	)

	...

	for _, a := range list {
		fmt.Println(a["full_name"], a["price"])
	}


Find the first 10 asteroids with an estimated profit value > $10 billion.

	list, err := asterank.Fetch(
		10,
		asterank.GTE("profit", 10e9),
	)

	...

	for _, a := range list {
		fmt.Println(a["full_name"], a["price"])
	}
*/
package asterank
