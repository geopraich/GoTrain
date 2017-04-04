package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
)
// if you declare a variable it must be used the _ tells the compiler you aren't using something
func main() {
	r, _ := http.Get("https://www.itjobswatch.co.uk/")
	page, _ := ioutil.ReadAll(r.Body)
	r.Body.Close()
	fmt.Printf("%s",page)
}