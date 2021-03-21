package main

import "fmt"

func main() {
	a := [...]int{0, 1, 2, 3} // an array
	fmt.Println(a)
	s := a[1:3] // s == []int{1, 2}        cap(s) == 3
	fmt.Println(s, cap(s))
	s = a[:2] // s == []int{0, 1}        cap(s) == 4
	fmt.Println(s, cap(s))
	s = a[2:] // s == []int{2, 3}        cap(s) == 2
	fmt.Println(s, cap(s))
	s = a[:] // s == []int{0, 1, 2, 3}  cap(s) == 4
	fmt.Println(s, cap(s))

	sl := make([]int, 0, 3)
	for i := 0; i < 5; i++ {
		sl = append(sl, i)
		fmt.Printf("cap %v, len %v, %p\n", cap(sl), len(sl), sl)
	}
}
