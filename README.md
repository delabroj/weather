`weather` tells you the weather

# Running locally

* `make build`
* `export OPEN_WEATHER_MAP_API_KEY=[your open weather map api key]`
* `./weather`
* `curl "localhost:8000/weather?lat=45&lon=45"`

# Running unit tests

* `make check`

# Updating mocks

* Install genmocks if needed: `go get -u github.com/delabroj/genmocks`
* `make mocks`
