package main

import "fmt"

func main() {
	// Call the maxHeightOfTriangle function with the given base and height
	fmt.Println(maxHeightOfTriangle(2, 4))
}

// This function returns the maximum height of a triangle with two different colored sides
func maxHeightOfTriangle(red int, blue int) int {
	// Call the helper function maxHeightOfTriangle2 to get the result
	result := maxHeightOfTriangle2(red, blue)
	// Call the helper function maxHeightOfTriangle2 again with the colors reversed
	result2 := maxHeightOfTriangle2(blue, red)
	// If the result2 is greater than or equal to the result, update the result
	if result2 >= result {
		result = result2
	}
	// Return the result
	return result
}

// Function to find the maximum height of a triangle with two different colors
func maxHeightOfTriangle2(color1 int, color2 int) int {
	// Initialize value to 1
	value := 1
	// Loop until value is odd and color2 is greater than or equal to value
	for {
		if value%2 == 0 && color1 >= value {
			// If value is even and color1 is greater than or equal to value, subtract value from color1
			color1 -= value
		} else if value%2 == 1 && color2 >= value {
			// If value is odd and color2 is greater than or equal to value, subtract value from color2
			color2 -= value
		} else {
			// If value is odd and color2 is less than value, break the loop
			break
		}
		// Increment value by 1
		value++
	}
	// Return the maximum height of the triangle
	return value - 1
}
