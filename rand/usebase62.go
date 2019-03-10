package main

import (
	"log"

	base62 "github.com/jbpratt78/apis/base62"
)

func main() {
	x := 100
	base62String := base62.ToBase62(x)
	log.Println(base62String)
	n := base62.ToBase10(base62String)
	log.Println(n)
}
