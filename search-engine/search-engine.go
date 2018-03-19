package search

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	WebSearch   = fakeSearch("web-server")
	ImageSearch = fakeSearch("image-server")
	VideoSearch = fakeSearch("video-server")
)

type Result string

type Search func(query string) Result

func fakeSearch(server string) Search {
	return func(query string) Result {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		return Result(fmt.Sprintf("%s result for %q\n", server, query))
	}
}

func Bing(query string) (results []Result) {
	results = append(results, WebSearch(query))
	results = append(results, ImageSearch(query))
	results = append(results, VideoSearch(query))
	return
}

func main() {
	rand.Seed(time.Now().UnixNano())
	start := time.Now()
	results := Bing("golang")
	elapsed := time.Since(start)
	fmt.Println(results)
	fmt.Println(elapsed)
}
