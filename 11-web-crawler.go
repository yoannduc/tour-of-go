// Check official answer as it does not use channels
// https://github.com/golang/tour/blob/master/solutions/webcrawler.go

package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

type SafeFetchedUrl struct {
	url map[string]bool
	mux sync.Mutex
}

type CrawlResult struct {
	url  string
	body string
}

//c.mux.Lock()
//defer c.mux.Unlock()

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, ch chan *CrawlResult, quit chan int, mux *SafeFetchedUrl) {
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	// This implementation doesn't do either:
	crawlHelper(url, depth, fetcher, ch, quit, mux)
	close(ch)
}

func crawlHelper(url string, depth int, fetcher Fetcher, ch chan *CrawlResult, quit chan int, mux *SafeFetchedUrl) {
	if depth <= 0 {
		return
	}

	mux.mux.Lock()
	if _, ok := mux.url[url]; ok {
		mux.mux.Unlock()
		return
	}

	mux.url[url] = true

	mux.mux.Unlock()
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	select {
	// Send value to channel
	case ch <- &CrawlResult{url: url, body: body}:
		// Value successfully sent.
	// If quit has value (ex is closed), return, closing the ch channel in Walk
	case <-quit:
		return
	}

	for _, u := range urls {
		crawlHelper(u, depth-1, fetcher, ch, quit, mux)
	}

}

func main() {
	ch, quit := make(chan *CrawlResult), make(chan int)
	mux := &SafeFetchedUrl{url: make(map[string]bool)}
	go Crawl("https://golang.org/", 4, fetcher, ch, quit, mux)

	for {
		c, ok := <-ch

		if !ok {
			break
		}

		fmt.Printf("found: %s %q\n", c.url, c.body)
	}

	close(quit)
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
