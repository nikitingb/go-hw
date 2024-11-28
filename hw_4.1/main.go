package main

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

type Words struct {
	word string
	count int
}

func main() {
	var result []Words
	
	str := "ass ass daas daad ddff daad ffcc ass daad daad daad"
	str = strings.ToLower(str)
	re := regexp.MustCompile(`\w+`)
	words := re.FindAllString(str, -1)

	m := make(map[string]int)
	for i := 0; i < len(words); i++ {
		m[words[i]]++
	}

	for key, val := range m {
		result = append(result, Words{word: key, count: val})
	}
	sort.Slice(result, func(i, j int) bool {
		return result[i].count > result[j].count
	})

	for i := 0; i < len(result) && i < 11; i++ {
		fmt.Println(result[i].word)
	}
}