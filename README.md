[![Build Status](https://travis-ci.org/sonirico/go-fist.svg?branch=master)](https://travis-ci.org/sonirico/go-fist)
[![Coverage Status](https://coveralls.io/repos/github/sonirico/go-fist/badge.svg?branch=master)](https://coveralls.io/github/sonirico/go-fist?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/sonirico/go-fist)](https://goreportcard.com/report/github.com/sonirico/go-fist)
[![GoDoc](https://godoc.org/github.com/sonirico/go-fist?status.svg)](https://godoc.org/github.com/sonirico/go-fist)
[![contributions welcome](https://img.shields.io/badge/contributions-welcome-brightgreen.svg?style=flat)](https://github.com/sonirico/go-fist/issues)
[![dependencies: none](https://img.shields.io/badge/dependencies-none-brightgreen.svg)]()


Golang client to interact with [Fist](https://github.com/f-prime/fist), a minimalist full-text index search server with
a focus on keeping things simple

## Install

```
go get -u github.com/sonirico/go-fist
```

## Examples

```go
import fistClient "github.com/sonirico/go-fist"

// ...

client, err := fistClient.NewFistClient("localhost", "5575")
if err != nil {
    fmt.Println("Connection Error! Is Fist up and running?")
    return
}
// Obtain server version
version, _ := client.Version()
fmt.Println("Server version is " + version)
// Index some data
client.Index("articles", "a an the")
client.Index("TODO", "wash the car")
client.Index("TODO", "walk the dog")
client.Index("podcasts", "DSE - Daily software engineering")
// Search for "the" keyword
documents := client.Search("the")
fmt.Println(documents) // ["articles", "TODO"]
// Not needing articles?
client.Delete("the")
documents = client.Search("the")
fmt.Println(documents) // []
```

More detailed examples can be found under the `./examples` subpackage

## Release strategy

Every time a new version for the server shall be released, so will the
client so as to keep a direct and easy to follow client/server version
mirroring

## License

Released under the terms of the MIT license. Refer to [LICENSE](LICENSE)
for more details.
