// In order to run this program, make sure there is a file in this directory named "input".
// Otherwise, it will not work!
package main

import (
  "fmt"
  "io/ioutil"
  "log"
  "errors"
)

func main() {

  input, err := ioutil.ReadFile("input")

  handleError(err)

  inputString := string(input)
  fmt.Println(inputString)
  floor, count, err := calculateFloor(inputString)

  handleError(err)

  fmt.Printf("Santa should go to Floor %d!\nThe instructiones were %d characters long!\n", floor, count)
}

// Calculates the floor that Santa should go to by checking each character and incrementing and decrementing as such.
// Stops at the first newline character.
// If there is a character that is not a newline, ( or ) , this function will return the floor on which Santa ended up,
// the instruction (index) where that character was, and an error.
func calculateFloor(input string) (int, int, error){

  result := 0
  i := 0
  for i = 0; i < len(input); i++{
    //fmt.Println("Calculating \"" + string(input[i]) + "\". . .")
    if string(input[i]) == ")"{
      result--
    } else if string(input[i]) == "("{
      result++
    } else if string(input[i]) == "\n"{
      break
    } else{
      fmt.Printf("Got to floor %d before the input broke at instruction %d.\n", result, i)
      return result, i, errors.New("Invalid type; remember, you can only use \"(\" and \")\"!")
    }
  }

  return result, i, nil
}

// Poor man's error handling.
func handleError(err error){
  if err != nil{
    log.Fatal(err)
  }
}
