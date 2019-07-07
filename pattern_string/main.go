package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	if s.Scan() {
		pattern := strings.TrimSpace(s.Text())
		results := BuildFromPatternString(pattern)
		for _, r := range results {
			fmt.Println(r)
		}
	}

	if err := s.Err(); err != nil {
		log.Fatal(err)
	}
}

const (
	typePlain = iota
	typePattern
)

type element struct {
	Type       int
	ValueStr   string
	ValueSlice []string
}

func BuildFromPatternString(pattern string) []string {
	// break down into elements
	var elements []element
	patternCounts := 1

	slices := strings.Split(pattern, "(")
	for _, sa := range slices {
		if strings.Contains(sa, ")") {
			sb := strings.Split(sa, ")")
			sc := strings.Split(sb[0], ",")
			patternCounts *= len(sc)
			elements = append(elements, element{ValueSlice: sc, Type: typePattern})
			if len(sb) > 1 && sb[1] != "" {
				elements = append(elements, element{ValueStr: sb[1], Type: typePlain})
			}
		} else {
			elements = append(elements, element{ValueStr: sa, Type: typePlain})
		}
	}

	// build from elements
	results := make([]string, patternCounts)
	loopCount := 1
	for _, elm := range elements {
		if elm.Type == typePlain {
			for k := range results {
				results[k] += elm.ValueStr
			}
		} else if elm.Type == typePattern {
			key := 0
			for i := 1; i <= patternCounts/(loopCount*len(elm.ValueSlice)); i++ {
				for _, s := range elm.ValueSlice {
					for j := 1; j <= loopCount; j++ {
						results[key] += s
						key++
					}
				}
			}
			loopCount *= len(elm.ValueSlice)
		}
	}
	return results
}
