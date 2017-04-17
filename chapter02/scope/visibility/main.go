package main

import "fmt"
import (
	"github.com/geopraich/GoTrain/chapter02/scope/packagescope/visible"
	"github.com/GoesToEleven/GolangTraining/27_code-in-process/39_packages/hello"
) // importing from visible package




func main() {
	fmt.Println(visible.MyName)
	visible.PrintName()
	hello.Hello()
}