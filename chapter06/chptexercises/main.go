package main

import (
	"fmt"
	"sort"

)

type people []string
// methods attached to people - implements sort interface for type people
func (p people) Len() int		{ return len(p) }
func (p people) Swap(i, j int)		{ p[i], p[j] = p[j], p[i] }
func (p people) Less(i, j int) bool	{ return p[i] < p[j] }

func main() {
	studyGroup := people{"Zeno","Beno","Pleno","Feno","Aeno"}
	sort.Strings(studyGroup)  // sorts alphabetically []string in place no assignment permanent
	fmt.Println(studyGroup)
	sort.Sort(studyGroup)  // it now is implements Interface interface
	sort.Sort(sort.Reverse(studyGroup))  // syntax for reverse sort
	fmt.Println(studyGroup)

	// can't attach methods to a non-type therefore cast to StringSlice() to allow sort.Sort()
	s := []string{"zeno","beno"}
	sort.Sort(sort.StringSlice(s))
	fmt.Println(s)


	n := []int{2,23,55,1,33,100}
	sort.Sort(sort.IntSlice(n))  // cast to IntSlice therefore implements interface type then can be sorted
	fmt.Println(n)

	sort.Sort(sort.Reverse(sort.IntSlice(n)))  // reverse sort
	fmt.Println(n)

	sort.Ints(n)  // sorts a slice of []int
	fmt.Println(n)
}