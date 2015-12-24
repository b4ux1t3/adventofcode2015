package main

import(
  "fmt"
  "io/ioutil"
  "log"
)

type Vertex struct{
  X, Y, Z int
}

func main() {
  // Read the string from a file
  input, err := ioutil.ReadFile("randout")

  handleError(err)

  // Convert the input file to a string
  inputString := string(input)

  // Variables to store current location and visted houses.
  x, y, z := 0, 0, 0
  var visited map[Vertex]int
  visited = make(map[Vertex]int)

  _, check := 0, false
  var coord Vertex

  // Loop through the characters, change x and y, then update the map
  for i := 0; i < len(inputString); i++{
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
    if string(inputString[i]) == "^"{
      y++
    } else if string(inputString[i]) == "v"{
      y--
    } else if string(inputString[i]) == ">"{
      x++
    } else if string(inputString[i]) == "<"{
      x--
    } else if string(inputString[i]) == "u"{
      z++
    } else if string(inputString[i]) == "u"{
      z--
    }
  }

  // Now we just need to print how many elements thre are in our map.
  fmt.Printf("%d houses visited.\n", len(visited))
}

// Poor man's error handling.
func handleError(err error){
  if err != nil{
    log.Fatal(err)
  }
}
