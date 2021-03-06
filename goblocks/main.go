package main

import (
	"fmt"
	"github.com/ashishnegi/chans"
	"github.com/ashishnegi/crawler"
	"github.com/ashishnegi/interfaces"
	"github.com/ashishnegi/stringutils"
	"strings"
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

type Coder struct {
	name string
}

var coders map[string]Coder

func WordCount(s string) map[string]int {
	fields := strings.Fields(s)
	res := map[string]int{}
	for _, field := range fields {
		_, present := res[field]
		if present {
			res[field] += 1
		} else {
			res[field] = 1
		}
	}
	return res
}

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	a, b := -1, 1
	return func() int {
		c := a
		a = b
		b = c + a
		return b
	}
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

	// map :
	coders = make(map[string]Coder)
	coders["haskell"] = Coder{"Ashish Negi"}
	fmt.Println(coders)

	fmt.Println(WordCount("hi.. i love haskell , clojure, c++, liking rust and may be go.. :)"))

	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}

	interfaces.TestInterfaces()
	interfaces.DynamicTypingIsIt()
	interfaces.CheckSqrt()

	arr := make([]int, 0, 100000)
	for i := 0; i < cap(arr); i++ {
		arr = append(arr, i)
	}

	fmt.Println(chans.TwiceFastSum(arr))

	crawler.CrawlTest()
}
