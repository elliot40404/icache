# ICACHE
A simple golang server for dynamic caching of images from external sources in-memory.

## Features

- [x] Caching of images from external sources
- [x] Dynamic resizing of images
- [x] Support for multiple image formats

## Usage

### Method 1

1. Clone the repository and `cd` into it
1. Run `make build` or `go build -o bin/icache cmd/api/main.go`
1. Run `./bin/icache`

### Method 2 [Recommended]

1. Clone the repository and `cd` into it
1. Run `docker build -t icache .`
1. Run `docker run --rm -d --name icache -e PORT=8080 -p 8080:8080 icache`
1. or just run `docker compose up -d`

## API

### GET /api/v1/image?url={url}

```sh
curl --request GET \
  --url 'http://localhost:8080/api/v1/image?url=https%3A%2F%2Fpicsum.photos%2F3000%2F2000' \
  --header 'Accept-Encoding: gzip, compress, br'
```
Make sure to encode the url parameter.

## License
MIT