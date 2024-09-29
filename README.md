# GoFaux

A versatile mock server that captures and stores JSON requests, allowing you to replay them as needed. Simply select from your stored responses to simulate real API interactions with ease.

## Get started
`go build -o gofaux cmd/gofaux/main.go`

## Usage

`./gofaux --url "URL_HERE" --name "JSON_NAME"`

`./gofaux --server`

### Flags 
`--url`: URL to fetch JSON from 

`--name`: Name of the stored JSON

`--server`: Starts a Gin server to return stored JSON 


