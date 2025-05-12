package main

import (
	"fmt"
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	//arr := strings.FieldsFunc(s, func(r rune) bool {
	//	return r == ','
	//})
	//fmt.Printf("%q\n", arr) // ["a" "b" "c" "d" "e"]
	//return map[string]int{"x": 1}
	m := make(map[string]int)
	wordsList := strings.Fields(s)
	for key, word := range wordsList {
		fmt.Println("key: ", key, "value: ", word)
		m[word]++
	}
	return m
}

func main() {
	dummy := WordCount("I am learning Go!")
	fmt.Println(dummy)
	wc.Test(WordCount)
}
