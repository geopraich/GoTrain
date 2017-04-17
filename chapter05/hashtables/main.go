package main

import (
	"fmt"
	"net/http"
	"log"
	// "io/ioutil"
	"bufio"
	"os"
)

func main() {
	// wordPrintSearch()
	pageScan()
}

func wordPrintSearch() {
	// gets text file assigns response to res and error to err
	res, err := http.Get("http://www-01.sil.org/linguistics/wordlists/english/wordlist/wordsEN.txt")
	if err != nil { // if the error is not nil ie if there is an error placed in log file
		log.Fatalln(err) // prints out error stops program if fatal
	}

	words := make(map[string]string)

	sc := bufio.NewScanner(res.Body)  // bufio input-output buffer
	sc.Split(bufio.ScanWords)  // split the words within sc

	for sc.Scan() {
		words[sc.Text()] = ""  // words get assigned to words map[string]string
	}
	if err := sc.Err(); err != nil {  // error check
		fmt.Fprintln(os.Stderr, "reading input", err)
	}
	count := 0
	for key := range words {
		fmt.Println(key)
		if count == 200 {  // breaks after first 200 words in map have been printed out
			break
		}
		count++
	}
/*
	bs, _ := ioutil.ReadAll(res.Body)  // response has many fields one is body reads the body of text
	str := string(bs)  // casts slice of bytes to string assigns to str
	fmt.Println(str)  // prints out str all words
*/
}

func pageScan() {
	res, err := http.Get("http://www.gutenberg.org/files/2701/old/moby10b.txt")
	if err != nil {
		log.Fatalln(err)
	}

	// scan page
	scanner := bufio.NewScanner(res.Body)
	defer res.Body.Close()  // closes file defer until exit of function
	// set the split function fo the scanning operation
	scanner.Split(bufio.ScanWords)
	// create slice to hold counts
	buckets := make([]int, 12)
	// loop over words
	for scanner.Scan() {
		l := HashBucket(scanner.Text(), 6)
		buckets[l]++
	}
	fmt.Println(buckets)
}

func HashBucket(word string, buckets int) int {
	var sum int
	for _, value := range word {
		sum += int(value)
	}
	return sum % buckets
}