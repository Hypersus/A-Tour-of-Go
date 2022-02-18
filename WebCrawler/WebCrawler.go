package main

import (
	"fmt"
	"sync"
	"time"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}
	
// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	// This implementation doesn't do either:
	if depth <= 0 {
		return
	}
	//set a flag default to fetch the url
	toFetch := true
	resCache := make(chan string, cap(cache))
	//fmt.Printf("depth is %v \n", depth)
	mu.Lock()
	//fmt.Printf("%v enter critical section with depth %v \n",url,depth)
	// no url in cache so simply send the url 
	close(cache)
	// sends and receives happen in the same goroutine so the cache need to be closed to iterate ??
	for item := range cache {
		//fmt.Printf("item is %v ,url is %v \n", item, url)
		if item != url {
			resCache <- item
		} else {
			toFetch=false
		}
	}
	resCache <- url
	cache = resCache
	if !toFetch {
		//fmt.Printf("will not fetch %v \n",url)
		defer mu.Unlock()
		return
	}
	//to fetch url
	mu.Unlock()
	//fmt.Printf("%v successfully unlock with depth %v \n",url,depth)
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		//fmt.Printf("go fetch url %v \n",u)
		go Crawl(u, depth-1, fetcher)
	}
	time.Sleep(time.Second)
	//fmt.Println("line 57")
	return
}

func main() {
	Crawl("https://golang.org/", 4, fetcher)
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
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
var cache = make(chan string, 10)
var mu sync.Mutex
