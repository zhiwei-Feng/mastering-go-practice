package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"sync"
	"time"
)

// 两个都不是缓存channel，只有sender和receiver都准备好了后它们的通讯才会发生(重点是都准备好)
var readValue = make(chan int)
var writeValue = make(chan int)

func set(newValue int) {
	writeValue <- newValue
}

func get() int {
	return <-readValue
}

func monitor() {
	var value int
	for {
		select {
		case newValue := <-writeValue:
			value = newValue
			fmt.Printf("%d ", value)
		case readValue <- value: // 只有先调用get()方法（receiver ready）才会不阻塞。
		}
	}
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please give an integer!")
		return
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Going to create %d random numbers.\n", n)
	rand.Seed(time.Now().Unix())
	go monitor()

	var w sync.WaitGroup
	for r := 0; r < n; r++ {
		w.Add(1)
		go func() {
			defer w.Done()
			set(rand.Intn(10 * n))
		}()
	}
	w.Wait()
	fmt.Printf("\nLast Value: %d\n", get())
}
