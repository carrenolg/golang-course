package hello

import "rsc.io/quote/v3"

// Hello return "Hello, world"
func Hello() string {
	return quote.HelloV3()
}

// Proverb retunr a Go concurrency proverb
func Proverb() string {
	return quote.Concurrency()
}
