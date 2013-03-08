// Copyright 2013 Bartosz Szczesny <bszcz@bszcz.eu> {bszcz.eu/license/MIT}

package main

import (
	"fmt"
	"math/rand"
	mt64 "mt19937_64"
	"time"
)

func PrintUint64(i int, mt *mt64.MT, ch chan bool) {
	t := rand.Int() % 1000
	time.Sleep(time.Duration(t))

	if 13 == i {
		mt.Init(12345)
		fmt.Printf("-- \n")
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
