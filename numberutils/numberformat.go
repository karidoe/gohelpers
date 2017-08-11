package numberutils

import "fmt"
import "strings"

type NumberClass struct {
  number        string
  decimals      int     "0"
  decPoint      string  "."
  thousandsSep  string  ","
}

func (c * NumberClass) reverseString(str string, withThousandsSep bool) string {
  var result string
  var count int

  arr := strings.Split(str, "")

  for i :=(len(arr) - 1); i >= 0; i-- {
    result += arr[i]
    count++
    if (count == 3 && withThousandsSep && i > 0) {
      result += c.thousandsSep
      count = 0
    }
  }

  return result
}

func NumberFormat(number string, decimals int, decPoint, thousandsSep string) string {
  var intString, decString string
  var obj = &NumberClass{}

  if decPoint != "" {
    obj.decPoint = "."
  }

  if thousandsSep != "" {
    obj.thousandsSep = ","
  }

  num := strings.SplitN(number, ".", 2)

  if len(num) < 2 {
    if strings.Index(number, ".") == -1 {
      intString = num[0]
    } else {
      intString = "0"
      decPoint = num[0]
    }
  } else {
    intString = num[0]
    decString = num[1]
  }

  fmt.Sprintf("0")

  str := obj.reverseString(intString, true)
  str = obj.reverseString(str, false)

  if decimals > 0 {
    decString += strings.Repeat("0", (decimals - len(decString)))
  }

  result := fmt.Sprintf("%s%s%s", str,  decPoint, decString)
  return result
}
