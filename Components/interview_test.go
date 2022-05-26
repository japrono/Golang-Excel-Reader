package customerimporter

import (
  "testing"
  "fmt"
)

func TestGetEmailDomain(t *testing.T) {
  var tests = []struct {
    EmailAddress string
    want1 bool
    want2 string
    want3 string
}{
    {"jakub.pronobis@gmail.com", true, "gmail.com", "jakub.pronobis"},
    {"jakub.pronobis@@gmail.com", true, "gmail.com", "jakub.pronobis@"},
    {"jakub.pronobis@@outlook.com", true, "outlook.com", "jakub.pronobis@"},
    {"jakub.pronobis", false, "", ""},
    {"", false, "", ""},
    {"@g.com", false, "", ""},
    {"user@", false, "", ""},
    {"@", false, "", ""},
    {"jakub@@@gmail.com", true, "gmail.com", "jakub@@"},
}

  for _, tt := range tests {
      testname := fmt.Sprintf("%v,%v,%v,%v", tt.EmailAddress, tt.want1, tt.want2, tt.want3)
      t.Run(testname, func(t *testing.T) {
          ans1, ans2, ans3 := GetEmailDomain(tt.EmailAddress)
          if ans1 != tt.want1 || ans2 != tt.want2 || ans3 != tt.want3 {
              t.Errorf("got %v %v %v, want %v %v %v", ans1, ans2, ans3, tt.want1, tt.want2, tt.want3)
          }
      })
  }
}








func TestAreHeadersValid(t *testing.T) {
  var tests = []struct {
    Header1 string
    Header2 string
    Header3 string
    Header4 string
    want1 bool
}{
    {"first_name", "last_name", "email", "gender", true},
    {"jakub.pronobis@@gmail.com", "true", "gmail.com", "jakub.pronobis@", false},
}

  for _, tt := range tests {
      testname := fmt.Sprintf("%v,%v,%v,%v,%v", tt.Header1, tt.Header2, tt.Header3, tt.Header4, tt.want1)
      t.Run(testname, func(t *testing.T) {
          ans1 := AreHeadersValid(tt.Header1, tt.Header2, tt.Header3, tt.Header4)
          if ans1 != tt.want1 {
              t.Errorf("got %v, want %v", ans1, tt.want1)
          }
      })
  }
}
