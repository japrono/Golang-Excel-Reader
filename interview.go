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
  "encoding/csv"
  "io"
  "log"
  "os"
)

type EmailDomain struct {
    Domain string
    Count     int
    Emails    []int
}

func GetEmailDomain(EmailAddress string) (bool, string, string) {
  at := strings.LastIndex(EmailAddress, "@")
  if at >= 0 {
      username, domain := EmailAddress[:at], EmailAddress[at+1:]

      var is_valid_email = true

      if username == "" || domain == "" {
        is_valid_email = false
        domain = ""
        username = ""
      }

      return is_valid_email, domain, username
  } else {
      var is_valid_email = false
      var domain = ""
      var username = ""
      return is_valid_email, domain, username
  }
}

func ParseExcelFile() (bool, string) {
  f, err := os.Open("customers.csv")

  if err != nil {
      log.Fatal(err)
      return false, ""
  }

  defer f.Close()

  csvReader := csv.NewReader(f)
  for {
      rec, err := csvReader.Read()
      if err == io.EOF {
          break
      }
      if err != nil {
          log.Fatal(err)
      }

      fmt.Printf("%+v\n", rec)
      fmt.Printf("%v %v %v %v", rec[0], rec[1], rec[2], rec[3])
      var is_valid_email, EmailDomain, username = GetEmailDomain(rec[2])
      fmt.Printf("\n")
      fmt.Printf("%v %v %v", is_valid_email, EmailDomain, username)

      fmt.Printf("\n")

  }


  return true, ""
}

func main() {
    var is_valid_email, Domain, username = GetEmailDomain("jakub.pronobis@gmail.com")
    s := strconv.FormatBool(is_valid_email)
    fmt.Println("value of is_valid_email = " + s)
    fmt.Println("value of Domain = " + Domain)
    fmt.Println("value of username = " + username)

    ParseExcelFile();

}
