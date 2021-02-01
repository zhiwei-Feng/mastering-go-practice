package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

type secret struct {
	RWM      sync.RWMutex
	M        sync.Mutex
	password string
}

var Password = secret{password: "myPassword"}

func Change(c *secret, pass string) {
	c.RWM.Lock()
	fmt.Println("LChange")
	time.Sleep(10 * time.Second)
	c.password = pass
	c.RWM.Unlock()
}

func show(c *secret) string {
	c.RWM.RLock()
	fmt.Print("show")
	time.Sleep(3 * time.Second)
	defer c.RWM.RUnlock()
	return c.password
}

func showWithLock(c *secret) string {
	c.M.Lock()
	fmt.Println("showWithLock")
	time.Sleep(3 * time.Second)
	defer c.M.Unlock()
	return c.password
}

func main() {
	var showFunction = func(c *secret) string { return "" }
	if len(os.Args) != 2 {
		fmt.Println("Using sync.RWMutex")
		showFunction = show
	} else {
		fmt.Println("Using sync.Mutex")
		showFunction = showWithLock
	}

	var w sync.WaitGroup
	fmt.Println("Pass:", showFunction(&Password))

	for i := 0; i < 15; i++ {
		w.Add(1)
		go func() {
			defer w.Done()
			fmt.Println("Go Pass:", showFunction(&Password))
		}()
	}

	go func() {
		w.Add(1)
		defer w.Done()
		Change(&Password, "123456")
	}()

	w.Wait()
	fmt.Println("Pass:", showFunction(&Password))
}
