# yuri

[![GoDoc](https://godoc.org/github.com/eskriett/yuri?status.svg)](https://godoc.org/github.com/eskriett/yuri)
[![Build
Status](https://travis-ci.org/eskriett/yuri.svg?branch=master)](https://travis-ci.org/eskriett/yuri)

The blazing fast way to **Y**ank **URI**s from text.

### Usage

    go get -u github.com/eskriett/yuri

```go
import "github.com/eskriett/yuri"

func main() {
    yuri.YankURIs([]byte("yuri lives at https://github.com/eskriett/yuri"))
    // []string{"https://github.com/eskriett/yuri"}
}
```

#### cmd/yuri

    go get -u github.com/eskriett/yuri/cmd/yuri

```shell
$ echo "I want to extract: http://example.com/" | yuri
http://example.com/
```

### Implementation details

yuri tries to extracting URIs of numerous schemes from text as fast as possible.
Compared to most similar tools which use regular expressions, yuri uses a
[DFA](https://en.wikipedia.org/wiki/Deterministic_finite_automaton) built using
[ragel](http://www.colm.net/open-source/ragel/) for performance.

The schemes yuri is currently able to extract are:

```
ftp
http
https
hxxp
hxxps
mailto
```

While the tool works well in many cases, it may sometimes return URIs which are
not fully compliant with their corresponding RFC (yuri is loosely based on the
ABNF provided by [RFC3987](https://www.ietf.org/rfc/rfc3987.txt)). If full RFC
complicance is a requirement for a given scheme, a post-pass URI validation
phase is advisable.

### Contributing

Please note that I developed yuri to solve a specific problem and am aware there
may be problems. If you notice yuri fails to extract a URI or want to submit
improvements (e.g. support for an additional scheme), please submit a pull
request which includes tests for your changes.
