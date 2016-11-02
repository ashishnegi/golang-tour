package main

import (
	"fmt"
	"github.com/ashishnegi/stringutils"
)

func returnTwoValues(num int) (int, int) {
	return num / 2, num / 3
}

func hereBePointers(q int) int {
	p := &q
	return (*p + 100)
}

type WillCrash struct {
	shouldCrash bool
}

func tryCrashing(crash WillCrash) bool {
	p := &crash
	return p.shouldCrash // crash.shouldCrash will also work.
}

func arraysAreHere(size int) [10]int {
	var a [10]int // do no know how to make variable size arrays yet..
	for i := 0; i < 10; i++ {
		a[i] = i
	}

	return a
}

func slicesOfArrays() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	s[1] = 33
	fmt.Println("changing slice : ", s, " changed primes :( : ", primes)
}

func dynamicSlices(size int) []int {
	a := make([]int, size)
	for i := 0; i < size; i++ {
		a[i] = i
	}
	return a
}

func main() {
	const clojure, haskell = "clojure: yo!!\n", "haskell: hell yaa!!!\n"
	fst, _ := returnTwoValues(100)

	defer fmt.Println("you know what i love... but such is life..")

	fmt.Println(stringutils.Reverse("Hello, go !!!."),
		"\n",
		haskell,
		clojure,
		fst,
		hereBePointers(101))

	crash := WillCrash{false}

	fmt.Println("should i crash ? ", tryCrashing(crash)) // do not know how to pass pointers yet.

	fmt.Println("arrays are here and there size can not be modified: ", arraysAreHere(100))

	slicesOfArrays()
	fmt.Println(dynamicSlices(20))
}
