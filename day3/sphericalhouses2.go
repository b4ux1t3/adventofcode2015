package main

import(
  "fmt"
  "io/ioutil"
  "log"
  "strings"
)

type Vertex3D struct{
  X, Y, Z int
}

type Vertex2D struct{
  X, Y int
}

func main() {
  input := getFile()

  var housesVisited int;
  if strings.Contains(input, "d") || strings.Contains(input, "u"){
    fmt.Println("Running in 3D mode.")
    housesVisited = run3D(input)
  } else {
    fmt.Println("Running in 2D mode.")
    housesVisited = run2d(input)
  }

  fmt.Printf("%d houses visited.\n", housesVisited)
}

// Gets file name from user, converts it to a string
func getFile() string{
  // Get the file name from the user
  var fileName string
  fmt.Println("Ho ho ho, where are the instructions?")
  fmt.Scanln(&fileName)

  // Read the string from a file
  input, err := ioutil.ReadFile(fileName)

  handleError(err)

  // Convert the input file to a string
  inputString := string(input)

  return inputString
}

// Runs the instructions in 3D mode
func run3D(instructions string) int{
  // Variables to store current location and visted houses.
  x, y, z := 0, 0, 0
  var visited map[Vertex3D]int
  visited = make(map[Vertex3D]int)

  _, check := 0, false
  var coord Vertex3D

  // Loop through the characters, change x and y, then update the map
  for i := 0; i < len(instructions); i++ {
    // Combine the coordinates into a string.
    coord.X = x
    coord.Y = y
    coord.Z = z
    //fmt.Printf("Checking coordinate: %d, %d\n" , coord.X, coord.Y)

    // Check if the given coordinate has been visited
    _, check = visited[coord]

    // If it hasn't been visited, add it to the list and set it to 1
    // Else, increment its present count
    if !check{
      visited[coord] = 1
    } else {
      visited[coord]++
    }

    // Finally, check which instruction is next, and increment or decrement x or y as needed
    if string(instructions[i]) == "^"{
      y++
    } else if string(instructions[i]) == "v"{
      y--
    } else if string(instructions[i]) == ">"{
      x++
    } else if string(instructions[i]) == "<"{
      x--
    } else if string(instructions[i]) == "u"{
      z++
    } else if string(instructions[i]) == "u"{
      z--
    }
  }

  return len(visited)
}

// Runs the instructions in 2D mode
func run2d(instructions string) int{
  // Variables to store current location and visted houses.
  x, y := 0, 0
  var visited map[Vertex2D]int
  visited = make(map[Vertex2D]int)

  _, check := 0, false
  var coord Vertex2D

  // Loop through the characters, change x and y, then update the map
  for i := 0; i < len(instructions); i++{
    // Combine the coordinates into a string.
    coord.X = x
    coord.Y = y
    //fmt.Printf("Checking coordinate: %d, %d\n" , coord.X, coord.Y)

    // Check if the given coordinate has been visited
    _, check = visited[coord]

    // If it hasn't been visited, add it to the list and set it to 1
    // Else, increment its present count
    if !check{
      visited[coord] = 1
    } else {
      visited[coord]++
    }

    // Finally, check which instruction is next, and increment or decrement x or y as needed
    if string(instructions[i]) == "^"{
      y++
    } else if string(instructions[i]) == "v"{
      y--
    } else if string(instructions[i]) == ">"{
      x++
    } else if string(instructions[i]) == "<"{
      x--
    }
  }

  return len(visited)
}

// Poor man's error handling.
func handleError(err error){
  if err != nil{
    log.Fatal(err)
  }
}
