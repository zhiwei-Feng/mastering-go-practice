package exercises

/*
Write a Go program that converts an existing array into a map.
*/
func ArrToMap(arr []int) map[int]int {
	m := map[int]int{}
	for i, v := range arr {
		m[i] = v
	}
	return m
}
