# JSON API

- Simple API to serve a JSON file as at GET endpoint
- Works with any valid JSON file

## Usage

### Build the app

- Make sure you have Go installed on your system
- Clone this repo, and run `go mod tidy`
- Compile with `go build -o appname .`
- Rename `test-data.json.example` or supply your own valid JSON file

### Run the app

`./appname 8090 test-data.json`

- first argument is the port you want it to run on
- second argument is the JSON file to serve

### Access the data via URL

`somehost:8090/data`
