package router

import (
  "fmt"
  "testing"
)

func TestRouterTableDriven(t *testing.T) {
    var tests = []struct {
        a string
        b int
        want string
    }{
        {"numbers", 0, "Zero"},
        {"numbers", 15, "Large"},
    }

    for _, tt := range tests {

        testname := fmt.Sprintf("%s %d", tt.a, tt.b)
        t.Run(testname, func(t *testing.T) {
            ans := Handle(tt.a, tt.b)
            if ans != tt.want {
                t.Errorf("got %s, want %s", ans, tt.want)
            }
        })
    }
  }
