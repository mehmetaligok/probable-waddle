package hello

import "rsc.io/quote/v3"

const version = "3.0.0-alpha.3"

func Hello() string {
	return quote.HelloV3()
}

func Proverb() string {
	return quote.Concurrency()
}
