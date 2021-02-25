// Copyright (c) 2013 Bartosz Szczesny
// LICENSE: The MIT License (MIT)

// Check if "mt19937_64" random number generator is concurency safe.
// Uses the "Source interface" from "math/rand".
package main

import (
	"fmt"
	mt64 "github.com/ndau/mt19937_64"
	"math/rand"
	"time"
)

// Print random int63 number from "mt19937_64".
// Reset random number generator if "i" is 13.
func PrintInt63(i int, mt *rand.Rand, ch chan bool) {
	t := rand.Int() % 1000
	time.Sleep(time.Duration(t))

	if 13 == i {
		mt.Seed(12345)
		fmt.Printf("13 : (reset) \n")
	} else {
		fmt.Printf("%2d : %d \n", i, mt.Int63())
	}

	ch <- true
}

func main() {
	rand.Seed(time.Now().Unix())

	mt := rand.New(mt64.New())
	mt.Seed(12345)

	ch := make(chan bool)
	for i := 0; i < 21; i++ {
		go PrintInt63(i, mt, ch)
	}
	<-ch
}
