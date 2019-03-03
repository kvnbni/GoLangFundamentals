package main

import (
	"fmt"
	"sort"
)

func main() {
	xi := []int{5, 8, 2, 43, 17, 987, 14, 12, 21, 1, 4, 2, 3, 93, 13}
	xs := []string{"random", "rainbow", "delights", "in", "torpedo", "summers", "under", "gallantry", "fragmented", "moons", "across", "magenta"}

	fmt.Println(xi) //print xi before sorting
	sort.Ints(xi)   //sorting
	fmt.Println(xi) //print xi after sorting

	fmt.Println(xs)  //print xs before sorting
	sort.Strings(xs) //sorting
	fmt.Println(xs)  //print xs after sorting
}
