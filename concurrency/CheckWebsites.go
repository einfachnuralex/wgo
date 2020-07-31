package concurrency

type WebsiteChecker func(string) bool

type result struct {
	url   string
	value bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)

	resultChannel := make(chan result)

	for _, url := range urls {
		go func(u string) {
			resultChannel <- result{
				url:   u,
				value: wc(u),
			}
		}(url)
	}

	for range urls {
		result := <-resultChannel
		results[result.url] = result.value
	}

	return results
}
