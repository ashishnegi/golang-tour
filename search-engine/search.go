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
	WebSearch1   = fakeSearch("web-server1")
	WebSearch2   = fakeSearch("web-server2")
	ImageSearch1 = fakeSearch("image-server1")
	ImageSearch2 = fakeSearch("image-server2")
	VideoSearch1 = fakeSearch("video-server1")
	VideoSearch2 = fakeSearch("video-server2")
)

// typedef alias
type Result string

// typedef func
type Search func(query string, cancel <-chan struct{}) Result

// function :
// func <func-name> ( argument-name argument-type ) return-type
func fakeSearch(server string) Search {
	// functions are first class citizens.
	// you can pass them around like variables.
	return func(query string, cancel <-chan struct{}) Result {
		server_query_simulation_timeout := time.After(time.Duration(rand.Intn(100)) * time.Millisecond)
		select {
		case <-server_query_simulation_timeout:
			fmt.Println("success: %s", server)
			return Result(fmt.Sprintf("%s : got result for %q\n", server, query))
		case <-cancel:
			fmt.Println("cancel: %s", server)
			return Result(fmt.Sprintf("%s : no result for %q\n", server, query))
		}
	}
}

func firstResult(query string, cancel <-chan struct{}, replicas ...Search) Result {
	resultChannel := make(chan Result, len(replicas))
	for i := range replicas {
		// gotcha // passing the i as parameter..
		go func(nested_i int) { resultChannel <- replicas[nested_i](query, cancel) }(i)
	}
	return <-resultChannel
}

// Bing3.0
// search results within a time frame across replicated clusters.
// with no locks, no condition variables, no callbacks.
func Bing(query string) (results []Result) {
	// channel is first class citizen. you can pass them around like values.
	// make a type safe channel.
	resultsChannel := make(chan Result, 3)
	cancelChannel := make(chan struct{})

	// go func() {... } () ==> starts function on a go routine.
	// to put value on a channel ==> channelName <- value.
	go func() { resultsChannel <- firstResult(query, cancelChannel, WebSearch1, WebSearch2) }()
	go func() { resultsChannel <- firstResult(query, cancelChannel, ImageSearch1, ImageSearch2) }()
	go func() { resultsChannel <- firstResult(query, cancelChannel, VideoSearch1, VideoSearch2) }()

	// make a timeout of 70 milliseconds for complete query.
	timeout := time.After(70 * time.Millisecond)
	for i := 0; i < 3; i++ {
		// take value from a channel ==> <-channelName.
		select {
		case result := <-resultsChannel:
			results = append(results, result)
		case <-timeout:
			close(cancelChannel)
			fmt.Println("Request timedout")
			return
			// all channels will be garbage collected.
			// but not go routines. // they will be lingering around untill web search is complete.
		}
	}

	// signal other goroutines that we are done.
	close(cancelChannel)

	// resultsChannel channel is garbage collected.
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

	// for clean shutdown we can add waitgroups.
	fmt.Println("Enter text to exit")
	var input string
	fmt.Scanln(&input)
}
