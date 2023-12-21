package numbers

import (
  "fmt"
  "testing"
)

func TestNumbersTableDriven(t *testing.T) {
    var tests = []struct {
        a int
        want string
    }{
        {9, "XLarge"},
        {8, "Large"},
        {7, "Large"},
        {6, "Medium"},
        {5, "Medium"},
        {4, "Smedium"},
        {3, "Smedium"},
    }

    for _, tt := range tests {

        testname := fmt.Sprintf("%d", tt.a)
        t.Run(testname, func(t *testing.T) {
            ans := CheckLarger(tt.a)
            if ans != tt.want {
                t.Errorf("got %s, want %s", ans, tt.want)
            }
        })
    }
}