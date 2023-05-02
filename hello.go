package hello

import "rsc.io/quote/v3"

const version = "1.0.1"

func Hello() string {
	return quote.HelloV3()
}

func Proverb() string {
	return quote.Concurrency()
}
