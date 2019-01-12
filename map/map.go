package main

import "fmt"

func main() {
	var nilTable map[string]int
	table := make(map[string]int)
	fmt.Println("var nilTable map[string]int => nilTable==nil?", nilTable == nil)
	fmt.Println("table := make(map[string]int) => table==nil?", table == nil)
	table = map[string]int{
		"lmj": 1,
		"sb":  2,
	}
	table["sb"] = 0
	table["ss"] = 3
	fmt.Println(table)

	for key, value := range table {
		fmt.Println(key, value)
	}

	value, ok := table["s"]
	fmt.Println(value, ok)

	value, ok = table["sb"]
	fmt.Println(value, ok)

}
