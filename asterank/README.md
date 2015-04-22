## Asterank

Package asterank provides Go bindings for the [NASA/JPL Asterank API][1].

[1]: http://www.asterank.com

The Asterank database is a thin layer over the NASA/JPL Small Body Database,
merged with JPL delta-v data, published asteroid mass data, and some custom
calculations.

> Asterank is a scientific and economic database of over 600,000 asteroids.
>
> We've collected, computed, or inferred important data such as asteroid mass
> and composition from multiple scientific sources. With this information, we
> estimate the costs and rewards of mining asteroids.
>
> Details on orbits and basic physical parameters are sourced from the Minor
> Planet Center and NASA JPL. Composition data is based on spectral
> classification and size. Our calculations incorporate conclusions from
> multiple scientific publications in addition to cross-referencing known
> meteorite data.


### Usage

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


Example result of the above:

	65679 (1989 UQ) 7.187709104774245e+10
	165 Loreley 1.6839856767077997e+11
	2357 Phereclos (1981 AC) 1.7754735017872946e+11
	1583 Antilochus (1950 SA) 2.1799831132867468e+11
	2674 Pandarus (1982 BC3) 1.9612031398729446e+11
	52 Europa 2.2093608573683997e+11
	(1999 CW8) 2.1706534476712106e+11
	85774 (1998 UT18) 9.501678062457524e+10
	1269 Rollandia (1930 SH) 2.4179032725926727e+11
	7474 (1992 TC) 8.400800732160027e+10


### Install

    go get github.com/jteeuwen/gonasa/asterank


### License

Unless otherwise stated, all of the work in this project is subject to a
1-clause BSD license. Its contents can be found in the enclosed LICENSE file.

