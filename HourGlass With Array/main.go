package main

import "fmt"

func main() {
	fmt.Println(hourGlass())
}

func hourGlass() int32 {
	arr := [][]int32{
		{3, 7, -3, 0, 1, 8},
		{1, -4, -7, -8, -6, 5},
		{-8, 1, 3, 3, 5, 7},
		{-2, 4, 3, 1, 2, 7},
		{2, 4, -5, 1, 8, 4},
		{5, -7, 6, 5, 2, 8},
	}
	sum := [16]int32{}
	temp := 0
	for key1 := 0; key1 < 4; key1++ {
		if key1 == 0 {
			temp = 0
		} else if key1 == 1 {
			temp = 4
		} else if key1 == 2 {
			temp = 8
		} else if key1 == 3 {
			temp = 12
		}
		for key2 := 0; key2 < 4; key2++ {
			sum[key2+temp] += arr[key1][key2] + arr[key1][key2+1] + arr[key1][key2+2]
			sum[key2+temp] += arr[key1+1][key2+1]
			sum[key2+temp] += arr[key1+2][key2] + arr[key1+2][key2+1] + arr[key1+2][key2+2]
		}
	}
	biggest := sum[0]
	for i := range sum {
		if sum[i] > biggest {
			biggest = sum[i]
		}
	}
	return biggest
}
