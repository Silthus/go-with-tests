package concurrency

type WebsiteChecker func(url string) bool
type result struct {
	url    string
	result bool
}

func CheckWebsite(wc WebsiteChecker, url string) bool {
	return wc(url)
}

func CheckWebsites(wc WebsiteChecker, urls []string) (results map[string]bool) {
	results = make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		go func(u string) {
			resultChannel <- result{u, wc(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		r := <-resultChannel
		results[r.url] = r.result
	}

	return results
}
