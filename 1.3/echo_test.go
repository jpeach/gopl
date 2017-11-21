package main

import (
	"fmt"
	"strings"
	"testing"
)

// Take the first 100 words for /usr/share/dict/words.
var words = []string{
	"A",
	"a",
	"aa",
	"aal",
	"aalii",
	"aam",
	"Aani",
	"aardvark",
	"aardwolf",
	"Aaron",
	"Aaronic",
	"Aaronical",
	"Aaronite",
	"Aaronitic",
	"Aaru",
	"Ab",
	"aba",
	"Ababdeh",
	"Ababua",
	"abac",
	"abaca",
	"abacate",
	"abacay",
	"abacinate",
	"abacination",
	"abaciscus",
	"abacist",
	"aback",
	"abactinal",
	"abactinally",
	"abaction",
	"abactor",
	"abaculus",
	"abacus",
	"Abadite",
	"abaff",
	"abaft",
	"abaisance",
	"abaiser",
	"abaissed",
	"abalienate",
	"abalienation",
	"abalone",
	"Abama",
	"abampere",
	"abandon",
	"abandonable",
	"abandoned",
	"abandonedly",
	"abandonee",
	"abandoner",
	"abandonment",
	"Abanic",
	"Abantes",
	"abaptiston",
	"Abarambo",
	"Abaris",
	"abarthrosis",
	"abarticular",
	"abarticulation",
	"abas",
	"abase",
	"abased",
	"abasedly",
	"abasedness",
	"abasement",
	"abaser",
	"Abasgi",
	"abash",
	"abashed",
	"abashedly",
	"abashedness",
	"abashless",
	"abashlessly",
	"abashment",
	"abasia",
	"abasic",
	"abask",
	"Abassin",
	"abastardize",
	"abatable",
	"abate",
	"abatement",
	"abater",
	"abatis",
	"abatised",
	"abaton",
	"abator",
	"abattoir",
	"Abatua",
	"abature",
	"abave",
	"abaxial",
	"abaxile",
	"abaze",
	"abb",
	"Abba",
	"abbacomes",
	"abbacy",
	"Abbadide",
}

// BenchmarkStringsJoin tests strings.Join().
func BenchmarkStringsJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strings.Join(words, " ")
	}
}

// BenchmarkImplicitJoin implicit join.
func BenchmarkImplicitJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Sprintln(words)
	}
}

// NOTE: strings.Join is faster than the reflection-based implicit join.
//
// $ go test -bench .
// goos: darwin
// goarch: amd64
// pkg: github.com/jpeach/gopl/1.3
// BenchmarkStringsJoin-8    	 1000000	      1066 ns/op
// BenchmarkImplicitJoin-8   	  100000	     11850 ns/op
// PASS
// ok  	github.com/jpeach/gopl/1.3	2.399s
