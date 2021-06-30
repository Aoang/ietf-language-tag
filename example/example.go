package main

import (
	"fmt"

	languageTag "github.com/Aoang/ietf-language-tag"
)

func main() {
	fmt.Println(languageTag.GetISO6391("zh-Hans"))
	fmt.Println(languageTag.GetISO31661("zh-Hans-CN"))
}
