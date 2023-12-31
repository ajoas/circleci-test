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
        {29, "XXXLarge"},
        {28, "XXXLarge"},
        {27, "XXXLarge"},
        {26, "XXXLarge"},
        {25, "XXXLarge"},
        {24, "XXXLarge"},
        {23, "XXXLarge"},
        {22, "XXXLarge"},
        {21, "XXXLarge"},
        {20, "XLarge"},
        {19, "XLarge"},
        {18, "XLarge"},
        {17, "XLarge"},
        {16, "XLarge"},
        {15, "XLarge"},
        {14, "XLarge"},
        {13, "XLarge"},
        {12, "XLarge"},
        {11, "XLarge"},
        {10, "XLarge"},
        {9, "XLarge"},
        {8, "Large"},
        {7, "Large"},
        {6, "Medium"},
        {5, "Medium"},
        {4, "Smedium"},
        {3, "Smedium"},
        {2, "Small"},
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
