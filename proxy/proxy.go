package main

type Proxy struct {
	realServer *Application
	maxAllowedRequests int
	rateLimiter map[string]int
}

func NewProxy() *Proxy {
	return &Proxy{
		realServer: &Application{},
		maxAllowedRequests: 2,
		rateLimiter: make(map[string]int),
	}
}

func (p *Proxy) handleRequest(url, method string) (int, string) {
	if !p.checkRateLimit(url) {
		return 429, "Too Many Requests"
	}
	
	return p.realServer.handleRequest(url, method)
}

func (p *Proxy) checkRateLimit(url string) bool {
	if p.rateLimiter[url] >= p.maxAllowedRequests {
		return false
	}
	
	p.rateLimiter[url]++
	return true
}