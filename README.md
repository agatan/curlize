## curlize

[![Build Status](https://travis-ci.org/agatan/curlize.svg?branch=master)](https://travis-ci.org/agatan/curlize)
[![GoDoc](https://godoc.org/github.com/agatan/curlize?status.svg)](https://godoc.org/github.com/agatan/curlize)

`curlize` converts `*http.Request` to curl shell commands.
`curlize` command is a mitm proxy that logs all requests as curl commands.

### Install

```
$ go get github.com/agatan/curlize # as a library
# or
$ go get github.com/agatan/curlize/cmd/curlize # as a command line tool
```

### CLI

`curlize` performs as an http proxy and reverse proxy, and dumps all requests as curl commands.

#### HTTP Proxy

Try `curlize -addr=:1234` and `curl -k -x localhost:1234 http://XXX/YYY/ZZZ`.

#### Reverse Proxy

Try `curlize -addr=:1234 -reverse=$UPSTREAM_HOST` and `curl localhost:1234/foo/bar`.

### Documentation

See https://godoc.org/github.com/agatan/curlize

### License

MIT
