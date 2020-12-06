package main

import (
  "fmt"
  "log"
  "io/ioutil"
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

func main() {
  file, err := ioutil.ReadFile("input")
  if err != nil {
    log.Fatal(err)
  }

	fileContent := string(file)
	slicedContent := strings.Split(fileContent, "\n")
  input := checkDelimiter(slicedContent, "\n")

  deltax := 3
  deltay := 1
  trees := 0

  for y, x := 0, 0; y < len(input); y, x = y + deltay, x + deltax {
    if string(input[y][x % len(input[y])]) == "#" {
      trees++
    }
  }

  fmt.Println(trees)
  return
}

