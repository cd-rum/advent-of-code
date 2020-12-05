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
  file, err := os.Open("day1")
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
      if i + j == 2020 {
        fmt.Println(i * j)
        return
      }
    }
  }
}
