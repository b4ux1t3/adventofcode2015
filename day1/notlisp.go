// In order to run this program, make sure there is a file in this directory named "input".
// Otherwise, it will not work!
package main

import (
  "fmt"
  "io/ioutil"
  "log"
  "errors"
  "os"
)

func main() {

  // Read the string from a file
  input, err := ioutil.ReadFile("input")

  handleError(err)

  // Convert the input file to a string
  inputString := string(input)

  fmt.Println(inputString)

  if len(os.Args) < 2 || os.Args[1] == "0"{
    floor, count, err := calculateFloor(inputString)
    handleError(err)
    fmt.Printf("Santa should go to Floor %d!\nThe instructions were %d characters long!\n", floor, count)
  } else if os.Args[1] == "1"{
    floor, count, err := calculateFloorNegOne(inputString)
    handleError(err)
    fmt.Printf("Santa hit Floor %d at instruction %d!\n", floor, count)
  } else {
    fmt.Println("Invalid argument. Use 1, 0, or nothing.")
  }



}

// Calculates the floor that Santa should go to by checking each character and incrementing and decrementing as such.
// Stops at the first newline character.
// If there is a character that is not a newline, ( or ) , this function will return the floor on which Santa ended up,
// the instruction (index) where that character was, and an error.
// NOTE: This function is called if there are no command line arguments, or if the 1th argument is 0
// NOTE: This is for part one of the problem
func calculateFloor(input string) (int, int, error){
  result, i := 0, 0

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

// Calculates at which instruction Santa will be at Floor -1
// Stops at the first newline character.
// If there is a character that is not a newline, ( or ) , this function will return the floor on which Santa ended up,
// the instruction (index) where that character was, and an error.
// If Santa never goes to Floor -1, this function will return the ending floor, the number of instructions, and an error
// NOTE: This function is called if the 1th argument is 1
// NOTE: This is for part two of the problem 
func calculateFloorNegOne(input string) (int, int, error){

  result, i := 0, 0

  for i = 0; i < len(input); i++{

    // Check if we've hit floor -1
    if result == -1{
      return result, i, nil
    }

    //fmt.Println("Calculating \"" + string(input[i]) + "\". . ."
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

  return result, i, errors.New("Did not hit Floor -1 at any point.")
}

// Poor man's error handling.
func handleError(err error){
  if err != nil{
    log.Fatal(err)
  }
}
