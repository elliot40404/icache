# ICACHE

A simple golang server for dynamic caching of images from external sources in-memory.

## Features

-   [x] Caching of images from external sources in-memory
-   [x] Dynamic resizing of images
-   [x] Support for multiple image formats
-   [x] Support for converting images to webp format
-   [x] On-the-fly resizing of images without ever storing them on disk

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

### GET /api/v1/image?url={url}&width={width}&height={height}&format={format}&resize={resize}&bypass={bypass}

Description: Returns a resized image from the cache or from the external source if not found in the cache.

```sh
curl --request GET \
  --url 'http://localhost:8080/api/v1/image?url=https%3A%2F%2Fpicsum.photos%2F3000%2F2000' \
  --header 'Accept-Encoding: gzip, compress, br'
```

#### Query Parameters

| Parameter | Type   | Description                                 | Default |
| :-------- | :----- | :------------------------------------------ | :------ |
| url       | string | The url of the image to be cached           | \*      |
| width     | int    | The width of the image to be resized to     | #       |
| height    | int    | The height of the image to be resized to    | #       |
| webp      | bool   | Whether to convert the image to webp or not | false   |
| resize    | bool   | Whether to resize the image or not          | true    |
| bypass    | bool   | Whether to bypass the cache or not          | false   |

**Note**
`*` - Required
`#` - Configurable via environment variables

Make sure to encode the url parameter.

### GET /api/v1/images

Description: Returns a list of all the images in the cache.

```sh
curl --request GET \
  --url 'http://localhost:8080/api/v1/images'
```

### DELETE /api/v1/image

Description: Deletes an image from the cache.

```sh
curl --request DELETE \
  --url 'http://localhost:8080/api/v1/image?url=https%3A%2F%2Fpicsum.photos%2F3000%2F2000'
```

#### Query Parameters

| Parameter | Type   | Description                                   | Default |
| :-------- | :----- | :-------------------------------------------- | :------ |
| url       | string | The url of the image to be deleted from cache | \*      |

**Note**
`*` - Required

Make sure to encode the url parameter.

### DELETE /api/v1/images

Description: Deletes all the images from the cache.

```sh
curl --request DELETE \
  --url 'http://localhost:8080/api/v1/images'
```

## Environment Variables

| Variable | Description                                                                  | Default |
| :------- | :--------------------------------------------------------------------------- | :------ |
| PORT     | The port on which the server will run                                        | 8080    |
| ADDR     | The address on which the server will run                                     | 0.0.0.0 |
| WIDTH    | The default width of the image to be resized to if not specified in request  | 400     |
| HEIGHT   | The default height of the image to be resized to if not specified in request | 400     |

**Note**
Do not set the `ADDR` if you are running the server in a docker container.

## SOME STATS

Using the following image: https://images8.alphacoders.com/133/1337846.png

My internet speed: 100mbps
CPU: Ryzen 9 5900x
RAM: 32gb 3600mhz
Disk: NVME SSD

Resize dimensions 300x300

| Type             | Size  | No Cache Times | Cached Times |
| :--------------- | :---- | :------------- | :----------- |
| no optimizations | 5.5mb | 2.84s          | 80ms         |
| resize           | 158kb | 825ms          | 3ms          |
| webp             | 148kb | 2.2s           | 2.5ms        |
| webp+resize      | 7.2kb | 997ms          | 0.80ms       |

## License

MIT
