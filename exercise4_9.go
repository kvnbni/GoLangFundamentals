package main

import "fmt"

func main() {
	m := map[string][]string{
		"bond_james":      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		"moneypenny_miss": []string{`James Bond`, `Literature`, `Computer Science`},
		"no_dr":           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}
	//print values along with index position in slice
	for k, v := range m {
		fmt.Println("For the key ", k, "the values are ")
		for i, val := range v {
			fmt.Println(i, val)
		}
	}
	// add new entry to map m
	m["benni_kevin"] = []string{`Momma`, `Pappa`, `Sis`}

	//delete entry from map m
	delete(m, "no_dr")

	for k, v := range m {
		fmt.Println(k, v)
	}

}