# Gorilla mux Router Adapter

This adapter must be used in combination with [api2go](https://github.com/manyminds/api2go) in order to be useful.
It allows you to use api2go within your [Gorilla mux](http://www.gorillatoolkit.org/pkg/mux) application.

## Example

```go
package main

import (
  "github.com/gorilla/mux"
  "github.com/manyminds/api2go"
  "github.com/manyminds/api2go-adapter/gorillamux"
  "net/http"
)

func main() {
  r := mux.NewRouter()
  api := api2go.NewAPIWithRouting(
    "api",
    api2go.NewStaticResolver("/"),
    api2go.DefaultContentMarshalers,
    gorillamux.New(r),
  )

  // Add your API resources here...
  // see https://github.com/manyminds/api2go for more information

  r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    w.Write("pong")
  }).Methods("GET")
  http.Handle("/", r)
}
```