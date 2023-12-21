package router

import (
  "github.com/ajoas/circleci-test/internal/foo"
  "github.com/ajoas/circleci-test/internal/numbers"
  "fmt"
)

func Handle(f string, n int) string {
  if f == "numbers" {
    nn := numbers.CheckLarger(n)
    fmt.Printf("numbers: %s\n",nn)
    return nn 
  }
  if f == "foo" {
    nf := foo.Foo(n)
    fmt.Printf("foo: %s\n",nf)
    return nf
  }
  return "Dang"
}
