package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func Pack(s string) string {
	var result strings.Builder
	isRepeat := false
	mapRune := make(map[rune]int)
	for _, charCode := range s {
		i, ok := mapRune[charCode]
		if ok {
			isRepeat = true
			mapRune[charCode] = i + 1
		} else {
			mapRune[charCode] = 1
		}
	}
	if !isRepeat {
		return s
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
		result.WriteString(strconv.Itoa(value))
	}
	return result.String()
}

func main() {
	fmt.Println(Pack("") == "")
}