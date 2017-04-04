package main

import "fmt"
import 	"github.com/geopraich/GoTrain/chapter02/scope/packagescope/visible"  // importing from visible package




func main() {
	fmt.Println(visible.MyName)
	visible.PrintName()
}