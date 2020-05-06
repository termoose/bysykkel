# Bysykkel

## Test it locally
- `git clone git@github.com:termoose/bysykkel.git`
- `cd bysykkel`
- Run some tests if you'd like: `go test -v ./...`
- `go run main.go`
- Visit http://127.0.0.1:8080 and start biking!

## On Heroku
- https://sheltered-sea-82728.herokuapp.com/

## Notes
- Only seems to work in Chrome with this old version of [OpenLayers](https://openlayers.org)
- Rendering on the backend to avoid writing JavaScript
- Why is `station_id` a `string`?
