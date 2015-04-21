## apod

Package apod provides Go bindings for the NASA [Astronomy Picture of the Day][1]
API.

[1]: https://data.nasa.gov/developer/external/planetary/#apod


### Usage

	date, err := time.Parse("2006-01-02", "2015-04-20")
	...

	resp, err := Fetch(&Request{
		Date:        date,
		ConceptTags: true,
		Key:         "DEMO_KEY",
	})

	...

	fmt.Println(resp.Url)


### Install

    go get github.com/jteeuwen/gonasa/apod


### License

Unless otherwise stated, all of the work in this project is subject to a
1-clause BSD license. Its contents can be found in the enclosed LICENSE file.

