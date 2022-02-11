package main

import ("fmt"; "os"; "strconv")

func add(x int, y int) int {
	return x + y
}

func strToI(str string) int{
  val, err := strconv.Atoi(str)
  if err !=nil { fmt.Println(err) }

  return val
}

func main() {
  args := os.Args[1:]
  fmt.Printf("%v + %v = %v\n", args[0], args[1], add(strToI(args[0]), strToI(args[1])))
}
