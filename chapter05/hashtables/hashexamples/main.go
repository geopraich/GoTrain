package main

import "fmt"

func main() {
	whichBucket()
}

func whichBucket() {
	// finds which bucket each letter will be placed in
	for num := 65; num <= 122; num++ {
		fmt.Println(num, " - ", string(num), " - ", num % 12)
	}
	fmt.Println(HashBucket("Go",12))
}

func HashBucket(word string, buckets int) int {
	// letter := rune(word[0]) - rune is int32 type
	letter := int(word[0])  // turns first letter into decimal
	bucket := letter % buckets  // finds remainder of letter divided by amount of buckets
	return bucket
}
