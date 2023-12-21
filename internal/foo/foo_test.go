package foo

import (
  "fmt"
  "testing"
)

func TestFooTableDriven(t *testing.T) {
    var tests = []struct {
        a int
        want string
    }{
        {0, "Sad"},
        {1, "Sad"},
        {3, "Sad"},
        {-1, "Sad"},
        {7, "Seven"},
        {8, "Sad"},
    }

    for _, tt := range tests {

        testname := fmt.Sprintf("%d", tt.a)
        t.Run(testname, func(t *testing.T) {
            ans := Foo(tt.a)
            if ans != tt.want {
                t.Errorf("got %s, want %s", ans, tt.want)
            }
        })
    }
}

func TestBarTableDriven(t *testing.T) {
    var tests = []struct {
        a int
        want string
    }{
        {0, "Sad"},
        {1, "Sad"},
        {3, "Three"},
        {-1, "Sad"},
        {8, "Sad"},
    }

    for _, tt := range tests {

        testname := fmt.Sprintf("%d", tt.a)
        t.Run(testname, func(t *testing.T) {
            ans := Bar(tt.a)
            if ans != tt.want {
                t.Errorf("got %s, want %s", ans, tt.want)
            }
        })
    }
}
