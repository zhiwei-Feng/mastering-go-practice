package main

import "fmt"
import ex "TimesAndDates/exercises"

func main() {
	UsingTime()
	fmt.Println("======= 分割线 =========")

	fmt.Println(ex.P4_0, ex.P4_1, ex.P4_2, ex.P4_3)
	fmt.Println(ex.Sunday, ex.Monday, ex.Tuesday, ex.Wednesday, ex.Thursday, ex.Friday, ex.Saturday)
	arr := []int{1, 2, 3, 4, 5}
	arrmap := ex.ArrToMap(arr)
	for key, val := range arrmap {
		fmt.Printf("%v->%v\n", key, val)
	}
}
