package main

import (
	"fmt"
	fdn "github.com/fefive-sxh/base-algorithm/sfOffer/find_dup_num"

)


func main() {
	fmt.Println("verify code is correct")
	nums := fdn.GenerateRandomArray(10)
	i := fdn.FindDuplication(nums)
	fmt.Println(nums)
	fmt.Println(i)
	fmt.Println("Pass")

	fmt.Println("test find duplication nums by using a bool map")
	nums1 := fdn.GenerateRandomArray(10)
	i1 := fdn.FindDuplication(nums1)
	fmt.Println(nums1)
	fmt.Println(i1)
	fmt.Println("Pass")

	fmt.Println("test find duplication nums by using a bool map")
	nums2 := fdn.GenerateRandomArray(10)
	i2 := fdn.FindDuplication(nums2)
	fmt.Println(nums2)
	fmt.Println(i2)
	fmt.Println("Pass")
}

