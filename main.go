package main

import (
	"bytes"
	"fmt"
	"strings"
)

func accum(s string) string {
	result := make([]string, len(s))
	arrString := strings.SplitAfter(s, "")
	for i, l := range arrString {
		var buffer bytes.Buffer
		for j := 1; j <= i+1; j++ {
			buffer.WriteString(strings.ToLower(l))
		}
		result[i] = strings.Title(buffer.String())
	}
	return strings.Join(result, "-")
}

func main() {
	fmt.Println(accum("abcd"))
	fmt.Println(accum("RqaEzty"))
	fmt.Println(accum("cwAt"))
}
