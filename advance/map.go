package advance

import (
	"fmt"
	"strings"
	"unicode"
)

func Map() {
	fmt.Println("This is map")
	//basic()
	wordCount("hello world")
}

type MapVertex struct {
	Lat, Long float64
}

var m map[string]MapVertex

func basic() {
	m = make(map[string]MapVertex)
	m["Bell Labs"] = MapVertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])

	// Map literals
	vertexes := map[string]MapVertex{
		"first":  {0.1, 1.2},
		"second": {3.4, 3.6},
	}
	vertexes["first"] = MapVertex{2.0, 2.1}
	fmt.Println(vertexes, vertexes["second"].Long)
	delete(vertexes, "second")
	fmt.Println(vertexes, vertexes["second"].Long)
	first, ok := vertexes["second"]
	fmt.Println(first, ok)
}

func wordCount(s string) map[string]int {
	wordCount := make(map[string]int)
	for _, rune := range strings.ToLower((s)) {
		// Don't count empty spaces
		if unicode.IsSpace(rune) {
			continue
		}
		runeAsString := string(rune)
		_, ok := wordCount[runeAsString]
		if ok {
			wordCount[runeAsString]++
		} else {
			wordCount[runeAsString] = 1
		}
	}
	fmt.Println(wordCount)
	return wordCount
}
