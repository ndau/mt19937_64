// Copyright (c) 2013 Bartosz Szczesny
// LICENSE: The MIT License (MIT)

// Check if "mt19937_64" random number generator is concurency safe.
package main

import (
	"fmt"
	mt64 "github.com/bszcz/mt19937_64"
	"math/rand"
	"time"
)

// Print random uint64 number from "mt19937_64".
// Reset random number generator if "i" is 13.
func PrintUint64(i int, mt *mt64.MT, ch chan bool) {
	t := rand.Int() % 1000
	time.Sleep(time.Duration(t))

	if 13 == i {
		mt.Init(12345)
		fmt.Printf("13 : (reset) \n")
	} else {
		fmt.Printf("%2d : %d \n", i, mt.Uint64())
	}

	ch <- true
}

func main() {
	rand.Seed(time.Now().Unix())

	mt := mt64.New()
	mt.Init(12345)

	ch := make(chan bool)
	for i := 0; i < 21; i++ {
		go PrintUint64(i, mt, ch)
	}
	<-ch
}
