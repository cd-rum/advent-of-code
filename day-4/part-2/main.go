package main

import (
  "fmt"
  "log"
  "io/ioutil"
  "regexp"
  "strings"
)

func checkDelimiter(slicedContent []string, delimiter string) []string {
  if delimiter == "\n" {
    return slicedContent[:len(slicedContent) - 1]
  } else {
    lastElement := slicedContent[len(slicedContent) - 1]
    slicedContent[len(slicedContent) - 1] = lastElement[:len(lastElement) - 1]

    return slicedContent
  }
}

func extract(input string, field string) string {
  str := fmt.Sprintln("%v:(.*?)", field)
  re := regexp.MustCompile(str)
  match := re.FindStringSubmatch(input)
  return match
}

func main() {
  file, err := ioutil.ReadFile("input")
  if err != nil {
    log.Fatal(err)
  }

  fileContent := string(file)
  slicedContent := strings.Split(fileContent, "\n")
  input := checkDelimiter(slicedContent, "\n")
  valid := 0

  for _, passport := range input {
    if strings.Contains(passport, "byr:") &&
    strings.Contains(passport, "iyr:") &&
    strings.Contains(passport, "eyr:") &&
    strings.Contains(passport, "hgt:") &&
    strings.Contains(passport, "hcl:") &&
    strings.Contains(passport, "ecl:") &&
    strings.Contains(passport, "pid:") {
      str := extract(passport, "byr")
      valid++
      fmt.Println(str)
    }
  }

  fmt.Println(valid)

  return
}

