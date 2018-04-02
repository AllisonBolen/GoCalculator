package main

import(
  "fmt"
  "bufio"
  "os"
  "reflect"
  "strings"
  "strconv"
  "math"
  )

  /*Continuously take user input until the user asks to quit.
    Keeps a running tally of the calculation.
    Supports +, -, *, /, and exponentiation of floating point and integer numbers
    (note, you should catch divide by zero exceptions in some way).
  */

func main(){
  /*this is my main calc*/
  scan := bufio.NewReader(os.Stdin)
  fmt.Println(reflect.TypeOf(scan))
  fmt.Println("Hello, this is your calculator! Format your requests like: <operator> <num> <num>")
  count := 1
  choice := readIn(scan)
  for choice != "Quit\n" {
    /* case switch statments */
    values := strings.Fields(choice)
    fmt.Print("Operation Count: ", count)
    fmt.Print("\n")
    switch operator := values[0]; operator{
    case "+" :
      fmt.Println("Addition Result: ", add(toNum(values[1], values[2])))
    case "-" :
      fmt.Println("Subtraction Result: ", sub(toNum(values[1], values[2])))
    case "*" :
      fmt.Println("Multiplication Result: ", mult(toNum(values[1], values[2])))
    case "/" :
      fmt.Println("Divison Result: ", div(toNum(values[1], values[2])))
    case "^" :
      fmt.Println("Exponentiation Result: ", expo(toNum(values[1], values[2])))
    default:
      fmt.Println("HELP MENU: \n \tThis supports the following operators: * ? _ + ^")
    }
    count = count + 1
    choice = readIn(scan)
  }
}

func toNum(a, b string) [2]float64{
  A, _ := strconv.ParseFloat(a, 64)
  B, _ := strconv.ParseFloat(b, 64)
  list := [2]float64{A, B}
  return list
}

func add(list [2]float64) float64{
  return  list[0] + list[1]
}

func sub(list [2]float64) float64{
  return list[0]-list[1]
}

func div(list [2]float64) float64{
  if list[1] == 0{
    fmt.Println("Cant divide by zero.")
    return 0
  }
  return list[0]/list[1]
}

func mult(list [2]float64) float64{
  return list[0]*list[1]
}

func expo(list [2]float64) float64{
  return math.Pow(list[0], list[1])
}

func readIn(scan *bufio.Reader) string{
  fmt.Println("What would you like to do: ")
  input, _ := scan.ReadString('\n')
  return input
}
