package hello

import "rsc.io/quote/v3"

const version = "0.1.0"

func Hello() string {
	return quote.HelloV3()
}

func Proverb() string {
	return quote.Concurrency()
}
