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
type Search func(query string) Result

// function :
// func <func-name> ( argument-name argument-type ) return-type
func fakeSearch(server string) Search {
    // functions are first class citizens.
    // you can pass them around like variables.
    return func(query string) Result {
        var sleepFor = rand.Intn(100)
        time.Sleep(time.Duration(sleepFor) * time.Millisecond)
        return Result(fmt.Sprintf("%s result for %q : time : %d\n", server, query, sleepFor))
    }
}

func firstResult(query string, replicas ...Search) Result {
    resultChannel := make(chan Result)
    for i := range replicas {
        // gotcha // passing the i as parameter..
        go func(nested_i int) { resultChannel <- replicas[nested_i](query) }(i)
    }
    return <-resultChannel
}

// Bing3.0
// search results within a time frame across replicated clusters.
// with no locks, no condition variables, no callbacks.
func Bing(query string) (results []Result) {
    // channel is first class citizen. you can pass them around like values.
    // make a type safe channel.
    resultsChannel := make(chan Result)
    // go func() {... } () ==> starts function on a go routine.
    // to put value on a channel ==> channelName <- value.
    go func() {
        resultsChannel <- firstResult(query, WebSearch1, WebSearch2)
        fmt.Println("finished web search...");
    }()

    go func() {
        resultsChannel <- firstResult(query, ImageSearch1, ImageSearch2)
        fmt.Println("finished image search...");
    }()

    go func() {
        resultsChannel <- firstResult(query, VideoSearch1, VideoSearch2)
        fmt.Println("finished video search...");
    }()

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
            // but not go routines. // they will be lingering around untill web search is complete.
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

    fmt.Println("Enter text to exit");
    var input string
    fmt.Scanln(&input)
}
