// Copyright (c) 2013 Bartosz Szczesny
// LICENSE: The MIT License (MIT)

// Reproduce the test vectors from the original "mt19937_64" random number generator.
//
// $ diff <(go run mt19937_64_test_vectors.go) ../data/mt19937-64.out.txt # check if no diff
//
// data source: http://www.math.sci.hiroshima-u.ac.jp/~m-mat/mt19937_64/mt19937-64.out.txt
package main

import (
	"fmt"
	mt64 "github.com/bszcz/mt19937_64"
)

func main() {
	initKey := []uint64{0x12345, 0x23456, 0x34567, 0x45678}
	var i uint64

	mt := mt64.New()
	mt.InitByArray(initKey)

	fmt.Printf("1000 outputs of genrand64_int64()\n") // line from mt19937-64.out.txt
	for i = 0; i < 1000; i++ {
		fmt.Printf("%20d ", mt.Uint64())
		if i%5 == 4 {
			fmt.Printf("\n")
		}
	}

	fmt.Printf("\n1000 outputs of genrand64_real2()\n") // line from mt19937-64.out.txt
	for i = 0; i < 1000; i++ {
		fmt.Printf("%10.8f ", mt.Real2())
		if i%5 == 4 {
			fmt.Printf("\n")
		}
	}
}
