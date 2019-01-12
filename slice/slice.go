package main

import "fmt"

/*
[0 0 0]
[1 2 0]
[1 2 3]
10 [0 0 0 0 0 0 0 0 0 1]

array:[5]int, slice:[]int, value:[1 2 3 4]
cap:4, len:4
array:[5]int, slice:[]int, value:[0 1 2 3]
cap:5, len:4
array:[5]int, slice:[]int, value:[0 1 2 3 4]
cap:5, len:5
[0 1 2 3 0 1 2 3], cap:10, len:8
[0 0 0], cap:5, len:3
*/

func main() {
	// array decalre
	var array1 [3]int
	var array2 [3]int = [3]int{1, 2}
	array3 := [...]int{1, 2, 3}
	array4 := [...]int{9: 1}
	fmt.Println(array1)
	fmt.Println(array2)
	fmt.Println(array3)
	fmt.Println(len(array4), array4)

	// slice
	fmt.Println()

	array := [5]int{0, 1, 2, 3, 4}
	s1 := array[1:]
	fmt.Printf("array:%T, slice:%T, value:%[2]v\n", array, s1)
	fmt.Printf("cap:%d, len:%d\n", cap(s1), len(s1))

	s1 = array[:4]
	fmt.Printf("array:%T, slice:%T, value:%[2]v\n", array, s1)
	fmt.Printf("cap:%d, len:%d\n", cap(s1), len(s1))

	s2 := s1[:len(array)]
	fmt.Printf("array:%T, slice:%T, value:%[2]v\n", array, s2)
	fmt.Printf("cap:%d, len:%d\n", cap(s2), len(s2))

	s3 := append(s1, s1...)

	fmt.Printf("%v, cap:%d, len:%d\n", s3, cap(s3), len(s3))

	// make init slice
	s4 := make([]int, 3, 5)
	fmt.Printf("%v, cap:%d, len:%d\n", s4, cap(s4), len(s4))
}
