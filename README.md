# ipfs-go-logrus

[![Go Reference](https://pkg.go.dev/badge/github.com/paralin/ipfs-go-logrus.svg)](https://pkg.go.dev/github.com/paralin/ipfs-go-logrus)
[![Go Report Card Widget]][Go Report Card]

[Go Report Card Widget]: https://goreportcard.com/badge/github.com/paralin/ipfs-go-logrus
[Go Report Card]: https://goreportcard.com/report/github.com/paralin/ipfs-go-logrus

ipfs-go-logrus modifies [ipfs/go-log](https://github.com/ipfs/go-log) to use
logrus instead of zap to use over 4MB less binary size.

## Install

```sh
go get github.com/ipfs/go-log
```

## Usage

Once the package is imported under the name `logging`, an instance of `EventLogger` can be created like so:

```go
var log = logging.Logger("subsystem name")
```

## License

MIT
