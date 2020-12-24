package exercises

/*
Write a constant generator iota for the powers of number 4.
*/
const (
	P4_0 int = 1 << (iota * 2)
	P4_1
	P4_2
	P4_3
)
