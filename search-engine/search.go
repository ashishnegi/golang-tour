// package needs to be main to generate an executable.
package main

// import other libs/packages.
import (
	"fmt"
	"math/rand"
	"time"
)

// define file scoped variables.
var (
	WebSearch   = fakeSearch("web-server")
	ImageSearch = fakeSearch("image-server")
	VideoSearch = fakeSearch("video-server")
)

// typedef alias
type Result string

// typedef func
type Search func(query string) Result

// function :
// func <func-name> ( argument-name argument-type ) return-type
func fakeSearch(server string) Search {
	// functions are first class citizens.
	// you can pass them around like variables.
	return func(query string) Result {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		return Result(fmt.Sprintf("%s result for %q\n", server, query))
	}
}

// Bing1.0
func Bing(query string) (results []Result) {
	results = append(results, WebSearch(query))
	results = append(results, ImageSearch(query))
	results = append(results, VideoSearch(query))
	return
}

// main function : looks like what we expected.
func main() {
	rand.Seed(time.Now().UnixNano())
	start := time.Now()
	results := Bing("golang")
	elapsed := time.Since(start)
	fmt.Println(results)
	fmt.Println(elapsed)
}
