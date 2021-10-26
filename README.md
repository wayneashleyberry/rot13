[![Go](https://github.com/wayneashleyberry/rot13/actions/workflows/go.yml/badge.svg)](https://github.com/wayneashleyberry/rot13/actions/workflows/go.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/wayneashleyberry/rot13)](https://goreportcard.com/report/github.com/wayneashleyberry/rot13)
[![GoDoc](https://godoc.org/github.com/wayneashleyberry/rot13?status.svg)](https://godoc.org/github.com/wayneashleyberry/rot13)

> A stupid rot13 implementation, written in Go.

### Installation

```sh
go get github.com/wayneashleyberry/rot13/cmd/rot13/...
```

### Usage

```sh
rot13 'Hello, World!'
Uryyb, Jbeyq!
```

```sh
rot13 --decode 'Uryyb, Jbeyq!'
Hello, World!
```
