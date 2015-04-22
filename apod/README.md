## apod

Package apod provides Go bindings for the NASA [Astronomy Picture of the Day][1]
API.

[1]: https://data.nasa.gov/developer/external/planetary/#apod


### API Keys

In order to use this API, an API key is required. For this purpose,
you must request one [here][2]. Alternatively, you can supply "DEMO_KEY" as
the api key. It will work, but with reduced capacity and features and should
be considered for testing only.

[2]: https://data.nasa.gov/developer/external/planetary/#apply-for-an-api-key


### Usage

	date, err := time.Parse("2006-01-02", "2015-04-20")
	...

	resp, err := Fetch("DEMO_KEY", date, true)
	...

	img, err := resp.Image()
	...


### Install

    go get github.com/jteeuwen/gonasa/apod


### License

Unless otherwise stated, all of the work in this project is subject to a
1-clause BSD license. Its contents can be found in the enclosed LICENSE file.

