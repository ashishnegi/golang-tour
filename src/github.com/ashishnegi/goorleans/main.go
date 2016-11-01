package main

import (
	"fmt"
	"github.com/ashishnegi/stringutils"
)

func returnTwoValues(num int) (int, int) {
	return num / 2, num / 3
}

func main() {
	const clojure, haskell = "clojure: yo!!\n", "haskell: hell yaa!!!\n"
	fst, _ := returnTwoValues(100)

	fmt.Println(stringutils.Reverse("Hello, go !!!."),
		"\n",
		haskell,
		clojure,
		fst)
}
