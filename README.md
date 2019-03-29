# Pandora

[![Join the chat at https://gitter.im/yandex/pandora](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/yandex/pandora?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)
[![Build Status](https://travis-ci.org/yandex/pandora.svg)](https://travis-ci.org/yandex/pandora)
[![Coverage Status](https://coveralls.io/repos/yandex/pandora/badge.svg?branch=develop&service=github)](https://coveralls.io/github/yandex/pandora?branch=develop)
[![Read the Docs](https://readthedocs.org/projects/yandexpandora/badge/)](https://readthedocs.org/projects/yandexpandora/)

Pandora is a high-performance load generator in Go language. It has built-in HTTP(S) and HTTP/2 support and you can write your own load scenarios in Go, compiling them just before your test.

## How to start

### Binary releases
[Download](https://github.com/yandex/pandora/releases) available.

### Building from sources
We use go 1.11 modules.
If you build pandora inside $GOPATH, please make sure you have env variable `GO111MODULE` set to `on`.
```bash
git clone https://github.com/yandex/pandora.git
make deps
go install
```

You can also cross-compile for other arch/os:
```
GOOS=linux GOARCH=amd64 go build
```

### Running your tests
Run the binary with your config (see config examples at [examples](https://github.com/yandex/pandora/tree/develop/examples)):

```bash
# $GOBIN should be added to $PATH
pandora myconfig.yaml
```

Or use Pandora with [Yandex.Tank](http://yandextank.readthedocs.org/en/latest/configuration.html#pandora) and
[Overload](https://overload.yandex.net).
