package main

import (
  "testing"
  "fmt"
)

func TestGetEmailDomain(t *testing.T) {
  var tests = []struct {
    EmailAddress string
    want1 bool
    want2 string
}{
    {"jakub.pronobis@gmail.com", true, "gmail.com"},
    {"jakub.pronobis@@gmail.com", true, "gmail.com"},
    {"jakub.pronobis@@outlook.com", true, "outlook.com"},
    {"jakub.pronobis", false, ""},
    {"", false, ""},
    {"@g.com", false, ""},
    {"user@", false, ""},
}

  for _, tt := range tests {
      testname := fmt.Sprintf("%v,%v,%v", tt.EmailAddress, tt.want1, tt.want2)
      t.Run(testname, func(t *testing.T) {
          ans1, ans2 := GetEmailDomain(tt.EmailAddress)
          if ans1 != tt.want1 || ans2 != tt.want2 {
              t.Errorf("got %v %v, want %v %v", ans1, ans2, tt.want1, tt.want2)
          }
      })
  }
}
