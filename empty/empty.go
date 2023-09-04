package main

import "fmt"

func main() {

	/*
	   s := i.(string)

	   fmt.Println("S:", s)

	   j, ok := i.(int)

	   	if ok {
	   		fmt.Println("j:", j)
	   	} else {

	   		fmt.Println("j Not an int")
	   	}

	*/
	var i any
	i = 7

	switch i.(type) {
	case int:

		fmt.Println("Is an int")

	case string:

		fmt.Println("Is a String")

	default:

		fmt.Printf("Unknown type%T\n", i)
	}

	fmt.Println(i)

	i = "Hi"

	fmt.Println(i)

	fmt.Println(max([]int{1, 4, 3, 7, 9}))
	fmt.Println(max([]float64{1, 4, 3, 7, 9}))

}

func maxInts(nums []int) int {

	if len(nums) == 0 {
		return 0
	}
	max := nums[0]
	for _, i := range nums[1:] {
		if i > max {
			max = i
		}
	}
	return max

}
func maxFloats(nums []float64) float64 {

	if len(nums) == 0 {
		return 0
	}
	max := nums[0]
	for _, i := range nums[1:] {
		if i > max {
			max = i
		}
	}
	return max

}

type Number interface {
	int | float64
}

func max[T Number](nums []T) T {

	if len(nums) == 0 {
		return 0
	}
	max := nums[0]
	for _, i := range nums[1:] {
		if i > max {
			max = i
		}
	}
	return max

}
