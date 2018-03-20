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

// Bing2.0
func Bing(query string) (results []Result) {
	// channel is first class citizen. you can pass them around like values.
	// make a type safe channel.
	resultsChannel := make(chan Result)

	// go func() {... } () ==> starts function on a go routine.
	// to put value on a channel ==> channelName <- value.
	go func() { resultsChannel <- WebSearch(query) }()
	go func() { resultsChannel <- ImageSearch(query) }()
	go func() { resultsChannel <- VideoSearch(query) }()

	// make a timeout of 70 milliseconds for complete query.
	timeout := time.After(70 * time.Millisecond)
	for i := 0; i < 3; i++ {
		// take value from a channel ==> <-channelName.
		select {
		case result := <-resultsChannel:
			results = append(results, result)
		case <-timeout:
			fmt.Println("Request timedout")
			return
			// all channels will be garbage collected.
			// but not go routines.
		}
	}
	// channel is garbage collected.
	// no need to close it.
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
