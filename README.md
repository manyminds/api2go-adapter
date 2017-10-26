# api2go-adapter

[![Join the chat at https://gitter.im/manyminds/api2go-adapter](https://badges.gitter.im/Join%20Chat.svg)](https://gitter.im/manyminds/api2go-adapter?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)

[![Build Status](https://travis-ci.org/manyminds/api2go-adapter.svg?branch=master)](https://travis-ci.org/manyminds/api2go-adapter)
[![Coverage Status](https://coveralls.io/repos/manyminds/api2go-adapter/badge.svg?branch=master&service=github)](https://coveralls.io/github/manyminds/api2go-adapter?branch=master)
[![Go Report Card](http://goreportcard.com/badge/manyminds/api2go-adapter)](http://goreportcard.com/report/manyminds/api2go-adapter)

### Deprecated

All adapters got moved into the main repository and will only be updated there.

#### Upgrade
To Upgrade the gin-adapter from api2go-adapter simply install api2go with the build tag:

`go get -tags=gingonic github.com/manyminds/api2go` and replace `gingonic.New(r)` with `routing.Gin(r)`
