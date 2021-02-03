package main

import (
	"container/list"
	"fmt"
	"strconv"
)

func printList(l *list.List) {
	for t := l.Back(); t != nil; t = t.Prev() {
		fmt.Print(t.Value, " ")
	}
	fmt.Println()

	for t := l.Front(); t != nil; t = t.Next() {
		fmt.Print(t.Value, " ")
	}
	fmt.Println()
}

func main() {
	values := list.New()

	e1 := values.PushBack("One")
	e2 := values.PushBack("Two")
	// One Two

	values.PushFront("Three")          // Three One Two
	values.InsertBefore("Four", e1)    // Three Four One Two
	values.InsertAfter("Five", e2)     // Three Four One Two Five
	values.Remove(e2)                  // Three Four One Five
	values.Remove(e2)                  // Three Four One Five, not work BECAUSE e2 is not exist.
	values.InsertAfter("FiveFive", e2) // Three Four One Five, not work BECAUSE e2 is not exist.
	values.PushBackList(values)        // Three Four One Five Three Four One Five

	printList(values)

	values.Init()
	fmt.Printf("After Init(): %v\n", values)

	for i := 0; i < 20; i++ {
		values.PushFront(strconv.Itoa(i))
	}

	printList(values)
}
