package main

import "fmt"
import . "github.com/karidoe/thaibahttext/numberutils"

func main() {
  s := NumberFormat("1234556.099", 3, ".", ",")
  fmt.Println(s)
}
