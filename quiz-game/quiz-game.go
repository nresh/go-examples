package main

import ("encoding/csv"
        "fmt"
        "os"
        // "strconv"
        "strings"
)

func main() {
  // args := os.Args[1:]
  // fmt.Printf("%v + %v = %v\n", args[0], args[1], add(strToI(args[0]), strToI(args[1])))

  f, _ := os.Open("quiz.csv")
  defer f.Close()

  csvReader := csv.NewReader(f)
  data, _   := csvReader.ReadAll()
  score     := 0

  for _, v := range data {
    values := strings.Split(v[0], "+")
    // fmt.Printf("%v %v %v\n", values[0], values[1], v[1])

    fmt.Printf("%v + %v = ", values[0], values[1])

    var input string

    fmt.Scanln(&input)

    if input == v[1] { score++ }
  }

  fmt.Printf("Your Score: %.2f", toPercentage(score, len(data)))

}

func toPercentage(num int, den int) float64 {
  return (float64(num) / float64(den)) * 100
}

// func add(x int, y int) int {
// 	return x + y
// }

// func strToI(str string) int {
//   val, err := strconv.Atoi(str)
//   if err != nil { fmt.Println(err) }
//
//   return val
// }
