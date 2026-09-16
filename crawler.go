package main

import (
	"fmt"
	"sync"
)

type FetchedUrls struct {
	urlsMap map[string]bool
	mutex   sync.Mutex
}

var urlsMap = FetchedUrls{
	urlsMap: make(map[string]bool),
}

var wg sync.WaitGroup

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

func Crawl(url string, depth int, fetcher Fetcher) {
	defer wg.Done()
	if depth <= 0 {
		return
	}
	urlsMap.mutex.Lock()
	_, ok := urlsMap.urlsMap[url]
	if ok {
		urlsMap.mutex.Unlock()
		return
	}
	urlsMap.urlsMap[url] = true
	urlsMap.mutex.Unlock()
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		wg.Add(1)
		go Crawl(u, depth-1, fetcher)
	}
}

func main() {
	wg.Add(1)
	go Crawl("https://golang.org/", 4, fetcher)
	wg.Wait()

}

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
