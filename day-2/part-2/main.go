package main

import (
  "os"
  "fmt"
  "log"
  "regexp"
  "bufio"
  "strconv"
)

type password struct {
  first, second    int
  letter, password string
}

func main() {
  file, err := os.Open("input")
  if err != nil {
      log.Fatal(err)
  }

  passwords := []*password{}
  inputRegex := regexp.MustCompile(`^(\d+)-(\d+) (\w): (\w+)$`)

  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    text := scanner.Text()
    first, _ := strconv.Atoi(inputRegex.ReplaceAllString(text, "$1"))
    second, _ := strconv.Atoi(inputRegex.ReplaceAllString(text, "$2"))
    passwords = append(passwords, &password{
      first:    first,
      second:   second,
      letter:   inputRegex.ReplaceAllString(text, "$3"),
      password: inputRegex.ReplaceAllString(text, "$4"),
    })
  }

  correct := 0
  for _, pw := range passwords {
    if (string(pw.password[pw.first-1]) == pw.letter) != (string(pw.password[pw.second-1]) == pw.letter) {
      correct++
    }
  }

  fmt.Println(correct)
  return
}

