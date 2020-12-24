package main

import (
	"fmt"
	"time"
)

func UsingTime() {
	fmt.Println("Epoch time:", time.Now().Unix())
	t := time.Now()
	fmt.Println(t, t.Format(time.RFC3339))
	fmt.Println(t.Weekday(), t.Day(), t.Month(), t.Year())

	time.Sleep(time.Second)
	t1 := time.Now()
	fmt.Println("Time difference:", t1.Sub(t))

	formatT := t.Format("01 January 2006")
	fmt.Println(formatT)
	loc, _ := time.LoadLocation("Europe/Paris")
	londonTime := t.In(loc)
	fmt.Println("Paris:", londonTime)

	formatCustom := t.Format("2006/01/02 15:04:05")
	fmt.Println(formatCustom)

	waitParsTime := "2021/11/11 19:08:55"
	parsedTime, err := time.Parse("2006/01/02 15:04:05", waitParsTime)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Full:", parsedTime)
	}
}
