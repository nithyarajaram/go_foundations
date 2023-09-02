package main

import (
	"fmt"
	"sort"
)

func main() {

	/*
			var s []int
			fmt.Println(len(s))
			if s == nil {
				fmt.Println("Nil Slice")
			}

			s2 := []int{1, 2, 3, 4, 5, 6, 7}
			fmt.Printf("Slice: %v\n", s2)

			s3 := s2[0:4]
			fmt.Printf("Slice: %v\n", s3)

			fmt.Println(s2[0:3])

			s3 = append(s3, 100)
			fmt.Printf("S3 %v\n", s3)

			fmt.Printf("S2 %v\n", s2)

			fmt.Printf("S2 Length: %d, S2 Capacity: %d\n", len(s2), cap(s2))
			fmt.Printf("S3 Length: %d, S3 Capacity: %d\n", len(s3), cap(s3))


		var s []int
		for i := 0; i < 1_000; i++ {
			s = appendInt(s, i)
		}
	*/

	fmt.Println(concat([]string{"A", "B", "C"}, []string{"D", "E", "F"}))

	vs := []float64{1, 2, 3, 4}
	ev := []float64{}

	fmt.Println(median(vs))
	fmt.Println(median(nil))
	fmt.Println(median(ev))
}

func appendInt(s []int, v int) []int {

	idx := len(s)

	if len(s) < cap(s) {
		s = s[:len(s)+1]

	} else {
		s1 := make([]int, 2*len(s)+1)
		fmt.Printf("Reallocate from %d -> %d\n", len(s), 2*len(s)+1)
		copy(s1, s)
		s = s1[:len(s)+1]
	}

	s[idx] = v
	return s
}

func concat(s1, s2 []string) []string {

	s3 := make([]string, len(s1)+len(s2))
	copy(s3, s1)
	fmt.Println(s3)
	copy(s3[len(s1):], s2)
	return s3
}

func median(values []float64) (float64, error) {

	if len(values) == 0 {
		return 0, fmt.Errorf("median of an empty slice")
	}

	//Copy in order to not change values
	nums := make([]float64, len(values))
	copy(nums, values)
	sort.Float64s(nums)
	i := len(nums) / 2
	//if odd number of values, median is values[len/2]
	if len(nums)%2 == 0 {
		median := (nums[i] + nums[i-1]) / 2
		return median, nil
	}

	median := nums[i]
	return median, nil

}
