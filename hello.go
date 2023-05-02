package hello

import "rsc.io/quote/v3"

const version = "2.0.0"

func Hello() string {
	return quote.HelloV3()
}

func Proverb() string {
	return quote.Concurrency()
}
