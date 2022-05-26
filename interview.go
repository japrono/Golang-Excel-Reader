// package customerimporter reads from the given customers.csv file and returns a
// sorted (data structure of your choice) of email domains along with the number
// of customers with e-mail addresses for each domain.  Any errors should be
// logged (or handled). Performance matters (this is only ~3k lines, but *could*
// be 1m lines or run on a small machine).
package main
import (
  "strconv"
  "fmt"
  "strings"
)

func GetEmailDomain(EmailAddress string) (bool, string) {
  at := strings.LastIndex(EmailAddress, "@")
  if at >= 0 {
      username, domain := EmailAddress[:at], EmailAddress[at+1:]

      var is_valid_email = true

      if username == "" || domain == "" {
        is_valid_email = false
        domain = ""
      }

      return is_valid_email, domain
  } else {
      var is_valid_email = false
      var domain = ""
      return is_valid_email, domain
  }
}

func main() {
    var is_valid_email, Domain = GetEmailDomain("jakub.pronobis@gmail.com")
    s := strconv.FormatBool(is_valid_email)
    fmt.Println("value of is_valid_email = " + s)
    fmt.Println("value of Domain = " + Domain)
}
