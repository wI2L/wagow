# :phone: wagow

[![GoDoc](https://godoc.org/github.com/wI2l/wagow?status.svg)](https://godoc.org/github.com/wI2l/wagow) [![Go Report Card](https://goreportcard.com/badge/github.com/wI2L/wagow)](https://goreportcard.com/report/github.com/wI2L/wagow) [![License](https://img.shields.io/badge/license-MIT-blue.svg)](/LICENSE)

A simple HTTP service that sends Wake-on-Wan magic packets.

### Installation

First, get the source code using `go get` or using `git clone` and ensure all dependencies are installed.
```sh
# Clone
$ git clone https://github.com/wI2L/wagow.git
# Install dependencies.
$ go get -u github.com/golang/dep/...
$ dep ensure
```

### Building

To build a docker image and launch it, you can use those commands that will produce a lightweight image and run it in background.

```sh
$ CGO_ENABLED=0 GOOS=linux go build -a -tags netgo -ldflags '-w' cmd/wagow/wagow.go
$ docker build -t wI2L/wagow .
$ docker run -d --publish 8080:8080 --name wagow wi2l/wagow
```

### Usage

The service expose a single route at `POST /` that sends a wake-on-wan magic packet to the destination machine upon receiving a request.

The parameters can be sent using either `application/json` or `application/x-form-www-urlencoded` content types.

   - `address`: IP address or FQDN of the machine in the format _host:port_
   - `target`: MAC-48 hardware address of the machine
   - `password`: the SecureOn password (_optional_)

If the destination address does not include a port, a default one (**9**) will be added automatically.

## License

Copyright (c) 2017, William Poussier (MIT Licence)
