package main

import "fmt"

func main() {
	var arr1 = []int{1, 2, 3, 4}
	res := DeleteSlice(arr1, 1)
	fmt.Printf("res: %v \n", res)
}

/*
DeleteSlice:接受两个参数，index，数组本身，返回剪切index后的新数组
*/
//func DeleteSlice[T any](index int, arr []T) []T {
//	j := 0
//	var arr1 []T
//	for k, v := range arr {
//		println(k, v)
//		if k != index {
//			arr1 = append(arr1, v)
//			j++
//		}
//	}
//	return arr1[:j]
//}

func DeleteSlice[T any](arr []T, index int) []T {
	res := append(arr[:index], arr[index+1:]...)
	return res
}
