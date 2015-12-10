package main

import(
	"fmt"
	"strings"
	"io/ioutil"
	"log"
	"strconv"
)

func main(){
	// Load the file in to memory.
	input, err := ioutil.ReadFile("input")
	handleError(err)
	// Convert the file into a string.
	inputString := string(input)

	// Split the string into a bunch of strings at each newline
	newLineSplit := strings.Split(inputString, "\n")
	// Now we should have an array of strings that look like this: "LLxWWxHH"

	// We need an array of arrays that can store our strings once we split them below.
	// This is our actual data. Hence the name.
	// The minus one is to deal with the last line of input, which will almost always be empty
	data := make([][]int, len(newLineSplit) - 1)

	// Split each of the resulting strings into length, width, and height by using x as a delimiter
	// temp will hold the split strings until we go through and convert them into ints
	// side is an iterator that is declared outside of the loop to save a bit of time in large datasets
	var side int
	for line := 0; line < len(newLineSplit) - 1; line++{
		temp := strings.Split(newLineSplit[line], "x")
		// First we have to make the line inside of our data
		data[line] = make([]int, 3)
		// This loop turns each of the strings into an int
		for side = 0; side < len(temp); side++{
			// Now we just store the values of temp in each element of data[line]
			data[line][side], err = strconv.Atoi(temp[side])
			handleError(err)
		}
	}
	// Now we have an nx3 array, where n is the number of lines in our input

	// Determine the total area needed for each package
	// This should be a for loop that performs a function
	// on each of the elements in the array created above
	// Add up the areas and store them in total
	totalFeet := 0
	for line := 0; line < len(data); line++{
		totalFeet += calcArea(data[line][0], data[line][1], data[line][2])
	}

	fmt.Printf("The elves need a total of %d square feet of wrapping paper!\n\n", totalFeet)
}

// This takes three inputs, and finds the total square footage needed for a given length, width, and height returns an int
// First it will take the three measurments needed (lxw, lxh, wxh)
// To find the slack needed, it will search the three values (lxw, lxh, wxh) for the smallest such value, and add it to the total square footage
// Finally, multiply each of the measurements by 2 and add them to teh total, then return total
func calcArea(length int, width int, height int) int{
	var i int
	measurements := [3]int {length * width, length * height, width * height}
	smallest := -1
  for i = 0; i < len(measurements); i++{
		if measurements[i] < smallest || smallest < 0{
			smallest = measurements[i]
		}
	}
	total := smallest
	for i = 0; i < len(measurements); i++{
		total += 2 * measurements[i]
	}

	return total
}

// Poor man's error handling.
func handleError(err error){
  if err != nil{
    log.Fatal(err)
  }
}
