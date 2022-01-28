# Library for timev2

[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/jjideenschmiede/gotimev2.svg)](https://golang.org/) [![Go](https://github.com/jjideenschmiede/gotimev2/actions/workflows/go.yml/badge.svg)](https://github.com/jjideenschmiede/gotimev2/actions/workflows/go.yml) [![Go Report Card](https://goreportcard.com/badge/github.com/jjideenschmiede/gotimev2)](https://goreportcard.com/report/github.com/jjideenschmiede/gotimev2) [![Go Doc](https://godoc.org/github.com/jjideenschmiede/gotimev2?status.svg)](https://pkg.go.dev/github.com/jjideenschmiede/gotimev2) ![Lines of code](https://img.shields.io/tokei/lines/github/jjideenschmiede/gotimev2) [![Developed with <3](https://img.shields.io/badge/Developed%20with-%3C3-19ABFF)](https://jj-dev.de/)

With this small library you can interact with the timev2 API.

## Install

```console
go get github.com/jjideenschmiede/gotimev2
```

## How to use?

### Get all times by month

With the following function you can read out all times for one month.

```go
// Define request
r := gotimev2.Request{
    SubDomain: "",
    Token:     "",
}

// Get all times by month
times, err := gotimev2.TimesByMonth("2006-01", r)
if err != nil {
    fmt.Println(err)
} else {
    fmt.Println(times)
}
```

### Get all clients

With this function you can read out all customers.

```go
// Define request
r := gotimev2.Request{
    SubDomain: "",
    Token:     "",
}

// Get all clients
clients, err := gotimev2.Clients(r)
if err != nil {
    fmt.Println(err)
} else {
    fmt.Println(clients)
}
```