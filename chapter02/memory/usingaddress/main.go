package main

import "fmt"

const metersToYards float64 = 1.09361

func main() {

	var meters float64
	fmt.Print("Enter Here Meters Walked: ")
	fmt.Scan(&meters)  // scans for user input sets to meters - requires & memory address - to store user input
	// in var meters - otherwise it immediately uses value stored within meters (0 currently as zero declaration)
	yards := meters * metersToYards
	fmt.Println(meters, "meters is", yards, "yards.")
}
