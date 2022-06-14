# Go Basic Reverse Proxy Example

[![Youtube video](https://img.shields.io/badge/youtube-video-brightgreen.svg)](https://youtu.be/tWSmUsYLiE4)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/traefik/traefik/blob/master/LICENSE.md)
[![Twitter](https://img.shields.io/twitter/follow/dadideo.svg?style=social)](https://twitter.com/intent/follow?screen_name=dadideo)
[![contributions welcome](https://img.shields.io/badge/contributions-welcome-brightgreen?logo=github)](CODE_OF_CONDUCT.md) 
[![Gitpod](https://img.shields.io/badge/Gitpod-ready--to--code-FFB45B?logo=gitpod)](https://gitpod.io/#https://github.com/davidaparicio/go-reverse-proxy) 


GBRPE (Go Basic Reverse Proxy Example) is a very small Golang application acting as a reverse proxy.

The code was presented during during a [talk at FOSDEM 2019](https://youtu.be/tWSmUsYLiE4):"How to write a reverse proxy with Go in 25 minutes" by [Julien Salleyron](https://mobile.twitter.com/juguul), Distinguished Engineer at [Traefik Labs](https://traefik.io/) (ex-Containous)

## Quick Start
The fastest way to start with GBRPE is the Gitpod-hosted version. When you're ready, you can install locally or host yourself.

One-click deploy sample URL Shortener application sample with GBRPE using Gitpod 

[![Open in Gitpod](https://gitpod.io/button/open-in-gitpod.svg)](https://gitpod.io/#https://github.com/davidaparicio/go-reverse-proxy)

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

<details>
  <summary>I used examples from the talks & materials</summary>

* FOSDEM2019 - How to write a reverse proxy with Go in 25 minutes by Julien Salleyron;
* GoLab2018 - Building a DIY proxy with the net package by Filippo Valsorda;

</details>