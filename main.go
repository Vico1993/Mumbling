package main

import (
	"bytes"
	"fmt"
	"strings"
)

func accum(s string) {
	result := make([]string, len(s))
	arrString := strings.SplitAfter(s, "")
	for i, l := range arrString {
		var buffer bytes.Buffer
		for j := 1; j <= i+1; j++ {
			buffer.WriteString(strings.ToLower(l))
		}
		result[i] = strings.Title(buffer.String())
	}
	fmt.Println(strings.Join(result, "-"))
}

func main() {
	accum("abcd")
	accum("RqaEzty")
	accum("cwAt")
}
