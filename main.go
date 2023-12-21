package main

import (
  "os"
  "fmt"
  "strconv"
  "github.com/ajoas/circleci-test/internal/router"
)

func main() {
  var f, n string
  var m int
  var e error
  f, n = os.Args[1], os.Args[2]
  m, e = strconv.Atoi(n)
  if e != nil {
    fmt.Fprintf(os.Stderr, "%v", e)
    os.Exit(1)
  }
  router.Handle(f,m)
}
