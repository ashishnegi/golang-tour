package bench

import (
	"encoding/json"
	// "strconv"
)

func fib(nth int) int {
	a, b, c, count := 0, 1, 0, 1
	for count < nth {
		c = a + b
		a = b
		b = c
		count++
	}
	return a
}

func hello(nth int) int {
	world("hello world...")
	n := fib(nth)
	world("nth world") // replace with strconv.Itoa(nth)
	return n
}

func world(s string) bool {
	return s == "not world"
}

type response2 struct {
	Pages  []int    `json:"page"`
	Fruits []string `json:"fruits"`
}

func tryJson(s string, i int) (string, int) {
	res := &response2{
		Pages:  []int{i, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		Fruits: []string{s, "peach", "pear", "apple", "peach", "pear", "apple", "peach", "pear", "apple", "peach", "pear", "apple", "peach", "pear", "apple", "peach", "pear"}}
	resJson, _ := json.Marshal(res)
	response := response2{}
	json.Unmarshal([]byte(resJson), &response)
	return response.Fruits[0], response.Pages[0]
}
