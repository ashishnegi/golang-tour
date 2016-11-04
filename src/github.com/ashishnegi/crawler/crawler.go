package crawler

import (
	"fmt"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

type Result struct {
	urls  []string
	depth int
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, ch chan Result) {
	if depth > 0 {
		_, urls, err := fetcher.Fetch(url)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("success : ", url)
		}

		ch <- Result{urls, depth}
	} else {
		ch <- Result{make([]string, 0), 0}
	}
}

func CrawlTest() {
	ch := make(chan Result)
	mainUrl := "http://golang.org/"
	go Crawl(mainUrl, 4, fetcher, ch)
	store := make(map[string]bool)
	store[mainUrl] = true
	waiting := 1

	// Fetch URLs in parallel. (see go before Crawl)
	// Don't fetch the same URL twice. (store keeps track of that.)

	for waiting > 0 {
		result := <-ch
		waiting--
		fmt.Println(result)
		if result.depth > 0 {
			for _, u := range result.urls {
				if _, exists := store[u]; !exists {
					waiting++
					// fmt.Println("trying", u)
					go Crawl(u, result.depth-1, fetcher, ch)
				}
				store[u] = true
			}
		}
	}

}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"http://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
}
