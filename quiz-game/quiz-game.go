package main

import ("encoding/csv"
        "flag"
        "fmt"
        "os"
        // "strconv"
        "strings"
        "time"
)

func main() {
  // args := os.Args[1:]
  // fmt.Printf("%v + %v = %v\n", args[0], args[1], add(strToI(args[0]), strToI(args[1])))

  csvFilename := flag.String("csv", "quiz.csv", "a csv file in the format of 'question,answer'")
  timeLimit   := flag.Int("max_time", 30, "number of seconds allowed for quiz")
  flag.Parse()

  f, _ := os.Open(*csvFilename)
  defer f.Close()

  csvReader := csv.NewReader(f)
  data, _   := csvReader.ReadAll()
  score     := 0

  var input string
  fmt.Println("Press enter to begin.")
  fmt.Scanln(&input)

  // start timer once they press enter
  timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

quizLoop:
  for _, v := range data {
    values := strings.Split(v[0], "+")
    fmt.Printf("%v + %v = ", values[0], values[1])

    inputChan := make(chan string)

    go func() {
      fmt.Scanln(&input)
  		inputChan <- strings.TrimSpace(input)
    }()

    select {
      case <-inputChan:
        if input == v[1] { score++ }
      case <-timer.C:
        break quizLoop // need to break this way to exit both select and loop
    }
  }

  fmt.Printf("\nYour Score: %.2f", toPercentage(score, len(data)))
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
