package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func Pack(s string) string {
	var result strings.Builder
	mapRune := make(map[rune]int)
	for _, charCode := range s {
		i, ok := mapRune[charCode]
		if ok {
			mapRune[charCode] = i + 1
		} else {
			mapRune[charCode] = 1
		}
	}
	keys := make([]rune, 0, len(mapRune))
	for k := range mapRune {
		keys = append(keys, k)
	}
	sort.SliceStable(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})
	for _, k := range keys {
		result.WriteRune(k)
		value := mapRune[k]
		if value > 1 {
			result.WriteString(strconv.Itoa(value))
		}
	}
	return result.String()
}

func main() {
	fmt.Println(Pack("") == "")
}