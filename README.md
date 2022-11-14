# JSON API

- Simple API to serve a JSON file as at GET endpoint

## Usage


### Build the app

- Make sure you have Go installed on your system
- Close this repo, and run `go mod tidy`
- Compile with `go build -o appname .`
- Rename `test-data.example.json` or supply your own valid JSON file

### Run the app

`./appname 8090 test-data.json`

- first argument is the port you want it to run on
- second argument is the JSON file
