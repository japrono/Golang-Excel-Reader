package customerimporter

import (
  "fmt"
  "strings"
  "encoding/csv"
  "io"
  "log"
  "os"
  "sort"
)

type Domain struct {
    Domain    string
    Count     int
    Emails    []string
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

func AreHeadersValid(Header1 string, Header2 string, Header3 string, Header4 string) (bool) {
  if(Header1 == "first_name" && Header2 == "last_name" && Header3 == "email" && Header4 == "gender") {
    return true
  } else {
    return false
  }
}

func IncreaseCounter(EmailDomains []Domain, EmailDomainToIncrease string, Username string) (bool, []Domain) {
  IncreasedCounter := false

  for index := range EmailDomains {
    if EmailDomains[index].Domain == EmailDomainToIncrease {
      EmailDomains[index].Count = EmailDomains[index].Count + 1
      EmailDomains[index].Emails = append(EmailDomains[index].Emails, Username)
      IncreasedCounter = true
      break
    }
  }

  return IncreasedCounter, EmailDomains
}

func ParseExcelFile(File string) (bool, []Domain) {
  EmailDomains := []Domain{}
  HeaderValidated := false

  f, err := os.Open(File)

  if err != nil {
      log.Fatal(err)
      return false, []Domain{}
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
          return false, []Domain{}
      }

      if HeaderValidated == false {
        if AreHeadersValid(rec[0], rec[1], rec[2], rec[3]) {
          HeaderValidated = true
          continue
        } else {
          log.Fatal("Excel headers are not correct")
          return false, []Domain{}
        }
      }

      var is_valid_email, EmailDomain, username = GetEmailDomain(rec[2])

      if is_valid_email == true {
        IncreasedCounter, _ := IncreaseCounter(EmailDomains, EmailDomain, username)

        if IncreasedCounter == false {
          EmailDomainObj := Domain{Domain: EmailDomain, Count: 1, Emails: []string{username} }
          EmailDomains = append(EmailDomains, EmailDomainObj)
        }

      } else {
        fmt.Printf("Warning: This email address is not valid: %v", rec[2])
        fmt.Printf("\n\tFull details [first_name, last_name, email, gender, ip_address] %v %v %v %v %v", rec[0], rec[1], rec[2], rec[3], rec[4])
        fmt.Printf("\n")
      }
  }

  EmailDomains = SortEmailDomains(EmailDomains)
  return true, EmailDomains
}

func PrintResults(EmailsDomains []Domain) {
  f, err := os.Create("results.txt")

  if err != nil {
      log.Fatal(err)
  }

  defer f.Close()

  for index := range EmailsDomains {
      string1 := fmt.Sprintf("%v %v \n", EmailsDomains[index].Count, EmailsDomains[index].Domain)
      _, err2 := f.WriteString(string1)

      if err2 != nil {
        log.Fatal(err2)
      }

      for index2 := range EmailsDomains[index].Emails {
        string2 := fmt.Sprintf("\t %v", EmailsDomains[index].Emails[index2])

        _, err3 := f.WriteString(string2)

        if err3 != nil {
          log.Fatal(err3)
        }
      }

      _, err4 := f.WriteString("\n\n")
      if err4 != nil {
        log.Fatal(err4)
      }
  }

  fmt.Println("\n\nOutput saved to file results.txt")
}

func SortEmailDomains(EmailDomains []Domain) ([]Domain) {
  sort.Slice(EmailDomains, func(i, j int) bool {
    return EmailDomains[i].Count > EmailDomains[j].Count
  })

  return EmailDomains
}
