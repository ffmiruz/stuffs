package main

import (
	"fmt"
	"math/rand"
)

func main() {
	num := 256
	success := 0
	for i := 0; i < 1000; i++ {
		a := randint(num)
		if a*a <= num {
			success++
		}
	}
	fmt.Println(success * num / 1000)

}

func randint(max int) int {
	min := 0
	return (rand.Intn(max-min) + min)
}
