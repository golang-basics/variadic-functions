package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(format(" steve hook ", strings.TrimSpace, strings.ToUpper))
	fmt.Println(format(format(" steve hook ", strings.TrimSpace, strings.ToUpper), strconv.Quote))

	str := format(" steve hook ", strings.TrimSpace, strings.ToUpper)
	fmt.Println(format(str, strconv.Quote))
	fmt.Println(format(" steve hook ", strings.TrimSpace, strings.ToUpper, strconv.Quote))
}

type formatter func(s string) string

func format(s string, formatters ...formatter) string {
	for _, formatter := range formatters {
		s = formatter(s)
	}
	return s
}
