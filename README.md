# Go Basic Reverse Proxy Example

[![Youtube video](https://img.shields.io/badge/youtube-video-brightgreen.svg)](https://youtu.be/tWSmUsYLiE4)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/traefik/traefik/blob/master/LICENSE.md)
[![Twitter](https://img.shields.io/twitter/follow/dadideo.svg?style=social)](https://twitter.com/intent/follow?screen_name=dadideo)


GBRPE (Go Basic Reverse Proxy Example) is a very small Golang application acting as a reverse proxy.
The code was presented during during a [talk at FOSDEM 2019](https://youtu.be/tWSmUsYLiE4) by [Julien Salleyron](https://mobile.twitter.com/juguul), Distinguished Engineer at [Traefik Labs (ex-Containous)](https://traefik.io/)

## Installation

Download the code with [git](https://git-scm.com/downloads) to use GBRPE.

```bash
git clone https://github.com/davidaparicio/go-reverse-proxy.git
cd 
```

## Requirements

Before the 1st run, generate self-signed SSL certificates for the HTTP2.

```bash
./generate_certs.sh
```

## Testing

```go
go run main.go
curl -ki https://localhost:8080
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)