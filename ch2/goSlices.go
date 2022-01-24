package main

import "fmt"

func main() {
	aSlice := []float64{}
	fmt.Println(aSlice, len(aSlice), cap(aSlice))
	aSlice = append(aSlice, 1234.56)
	aSlice = append(aSlice, -34.0)
	fmt.Println(aSlice, "with length", len(aSlice), "and cap", cap(aSlice))

	t := make([]int, 4)
	t[0] = -1
	t[1] = -2
	t[2] = -3
	t[3] = -4
	t = append(t, -5)
	fmt.Println("Value of t: ", t)

	twoD := [][]int{{1, 2, 3}, {4, 5, 6}}
	for _, i := range twoD {
		for _, k := range i {
			fmt.Print(k, " ")
		}
		fmt.Println()
	}

	make2D := make([][]int, 2)
	fmt.Println(make2D)
	make2D[0] = []int{1, 2, 3, 4}
	make2D[1] = []int{-1, -2, -3, -4}
	fmt.Println(make2D)
}