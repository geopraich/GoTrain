package main

import ("fmt"
	"encoding/json"
	"os"
	"strings"
)

type Mannie struct {
	// all exported upper case first letter / lower case will not export
	// can use tags to store in flied if another field is specified
	First string
	Last string
	Age int  // `json:"death race"` anything coming in as death race will be stored in age
	Status string
}

func main() {
	p1 := Mannie{"James","Bond", 20,"007"}
	bs, _ := json.Marshal(p1)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)
	fmt.Println(string(bs))

	var p2 Mannie  // initialised to zero values

	bys := []byte(`{"First":"Roger","Last":"Moore","Age":20,"Status":"009"}`)  // slice of bytes
	json.Unmarshal(bys, &p2)  // takes slice of bytes and address to p2 initialises fields

	fmt.Println(p2)

	json.NewEncoder(os.Stdout).Encode(p1)  // newEncoder returns *encoder Encode takes a *encoder writes p1 to scrn

	var p3 Mannie
	rdr := strings.NewReader(`{"First":"Roger","Last":"Moore","Age":20,"Status":"009"}`)
	json.NewDecoder(rdr).Decode(&p3)  // NewReader returns a *reader NewDecoder takes a *reader returns
	// a pointer to a decoder *Decoder which Decode takes  Decode initialises Mannie with rdr data
	fmt.Println(p3)
}
