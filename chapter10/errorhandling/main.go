package main

import (
	"os"
	"fmt"
	"log"
)

func init() {
	nf, err := os.Create("log.txt")
	if err != nil {
		fmt.Println(err)
	}
	log.SetOutput(nf)  // sets to write err log to the newfile log.txt - nf
}

// idiomatic way of dealing with errors - 4 different ways below
func main() {
	_, err := os.Open("no-file.txt.")
	if err != nil {
		// fmt.Println("error happened:", err)
		// log.Println("err happened", err) log. will write to log.txt on err due to init
		// log.Fatalln(err)
		// panic(err)
	}
}

// Println formats using the default formats for its operands and writes to standard output
// Fatal functions call os.Exit(1) after writing the log message
// Panic functions call the panic after writing the log message
// Exit exits the code with status code non-zero which indicates error (0 success)