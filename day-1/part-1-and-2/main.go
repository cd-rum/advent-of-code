package main

import (
  "os"
  "fmt"
  "log"
  "bufio"
  "strconv"
)

func main() {
  numbers := []int{}
  file, err := os.Open("input")
  if err != nil {
      log.Fatal(err)
  }

  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    n, _ := strconv.Atoi(scanner.Text())
    numbers = append(numbers, n)
  }

  for _, i := range numbers {
    for _, j := range numbers {
      for _, k := range numbers {
        if i + j + k == 2020 {
          fmt.Println(i * j * k)
          return
        }
      }
    }
  }
}
