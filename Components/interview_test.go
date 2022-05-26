package customerimporter

import (
  "testing"
  "fmt"
  "reflect"
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

func TestIncreaseCounter(t *testing.T) {
  Domain1 := Domain{ Domain: "gmail.com", Count: 1, Emails: []string{"jakub.pronobis"} }
  EmailDomains1 := []Domain{Domain1}
  EmailDomainToIncrease1 := "gmail.com"
  Username1 := "jakub.pronobis"
  Domain2 := Domain{ Domain: "gmail.com", Count: 2, Emails: []string{"jakub.pronobis", "jakub.pronobis"} }
  EmailDomains2 := []Domain{Domain2}

  EmailDomains3 := []Domain{}
  EmailDomainToIncrease2 := "gmail.com"
  Username2 := "jakub.pronobis"

  var tests = []struct {
    EmailDomains []Domain
    EmailDomainToIncrease string
    Username string
    want1 bool
    want2 []Domain
}{
    {EmailDomains1, EmailDomainToIncrease1, Username1, true, EmailDomains2},
    {EmailDomains3, EmailDomainToIncrease2, Username2, false, EmailDomains3},

}

  for _, tt := range tests {
      testname := fmt.Sprintf("%v,%v,%v,%v,%v", tt.EmailDomains, tt.EmailDomainToIncrease, tt.Username, tt.want1, tt.want2)
      t.Run(testname, func(t *testing.T) {
          ans1, ans2 := IncreaseCounter(tt.EmailDomains, tt.EmailDomainToIncrease, tt.Username)
          if ans1 != tt.want1 || !reflect.DeepEqual(ans2, tt.want2) {
              t.Errorf("got %v %v, want %v %v", ans1, ans2, tt.want1, tt.want2)
          }
      })
  }
}

func TestParseExcelFile(t *testing.T) {
  Domain1 := Domain{ Domain: "gmail.com", Count: 2, Emails: []string{"jakub.pronobis", "damian.pronobis"} }
  Domain2 := Domain{ Domain: "github.io", Count: 1, Emails: []string{"mhernandez0"} }
  Domain3 := Domain{ Domain: "cyberchimps.com", Count: 1, Emails: []string{"bortiz1"} }
  Domain4 := Domain{ Domain: "hubpages.com", Count: 1, Emails: []string{"dhenry2"} }
  Domain5 := Domain{ Domain: "outlook.com", Count: 1, Emails: []string{"jakub.pronobis"} }

  EmailDomains1 := []Domain{Domain1, Domain2, Domain3, Domain4, Domain5}

  var tests = []struct {
    File string
    want1 bool
    want2 []Domain
}{
    {"customers_test_1.csv", true, EmailDomains1},
}

  for _, tt := range tests {
      testname := fmt.Sprintf("%v,%v,%v", tt.File, tt.want1, tt.want2)
      t.Run(testname, func(t *testing.T) {
          ans1, ans2 := ParseExcelFile(tt.File)
          if ans1 != tt.want1 || !reflect.DeepEqual(ans2, tt.want2) {
              t.Errorf("got %v %v, want %v %v", ans1, ans2, tt.want1, tt.want2)
          }
      })
  }
}



func TestSortEmailDomains(t *testing.T) {
  Domain1 := Domain{ Domain: "EmailDomain", Count: 1, Emails: []string{"username"} }
  Domain2 := Domain{ Domain: "EmailDomain", Count: 51, Emails: []string{"username"} }
  EmailDomains1 := []Domain{Domain1, Domain2}

  Domain3 := Domain{ Domain: "EmailDomain", Count: 51, Emails: []string{"username"} }
  Domain4 := Domain{ Domain: "EmailDomain", Count: 1, Emails: []string{"username"} }
  EmailDomains2 := []Domain{Domain3, Domain4}

  Domain5 := Domain{ Domain: "EmailDomain", Count: 1549, Emails: []string{"username"} }
  Domain6 := Domain{ Domain: "EmailDomain", Count: 51, Emails: []string{"username"} }
  EmailDomains3 := []Domain{Domain5, Domain6}

  Domain7 := Domain{ Domain: "EmailDomain", Count: 1549, Emails: []string{"username"} }
  Domain8 := Domain{ Domain: "EmailDomain", Count: 51, Emails: []string{"username"} }
  EmailDomains4 := []Domain{Domain7, Domain8}


  var tests = []struct {
    EmailDomains []Domain
    want1 []Domain
}{
    {EmailDomains1, EmailDomains2},
    {EmailDomains3, EmailDomains4},
}

  for _, tt := range tests {
      testname := fmt.Sprintf("%v,%v", tt.EmailDomains, tt.want1)
      t.Run(testname, func(t *testing.T) {
          ans1 := SortEmailDomains(tt.EmailDomains)
          if !reflect.DeepEqual(ans1, tt.want1) {
              t.Errorf("got %v, want %v", ans1, tt.want1)
          }
      })
  }
}
