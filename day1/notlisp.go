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
  floor, err := calculateFloor(inputString)

  handleError(err)

  fmt.Printf("Santa should go to Floor %d!\n", floor)
}

// Calculates the floor that Santa should go to by checking each character and incrementing and decrementing as such.
// Stops at the first newline character.
// Returns an error if there is an invalid character.
func calculateFloor(input string) (int, error){

  result := 0

  for i := 0; i < len(input); i++{
    //fmt.Println("Calculating \"" + string(input[i]) + "\". . .")
    if string(input[i]) == ")"{
      result--
    } else if string(input[i]) == "("{
      result++
    } else if string(input[i]) == "\n"{
      break
    } else{
      return 0, errors.New("Invalid type")
    }
  }

  return result, nil
}

// Poor man's error handling.
func handleError(err error){
  if err != nil{
    log.Fatal(err)
  }
}
