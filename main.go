package main

import (
	"fmt"
	"go-optional-type/optional"
	"io/ioutil"
	"log"
	"net/http"
)

func fetchURLContent(url string) optional.Optional[string] {
	resp, err := http.Get(url)
	if err != nil {
		return optional.Empty[string]()
	}
	defer resp.Body.Close()
	html, err := ioutil.ReadAll(resp.Body)
	return optional.Of[string](string(html))
}

func main() {
	_, err := optional.Map[string, int](fetchURLContent("https://baidu.com"), func(val string) int {
		return len(val)
	}).IfPresent(func(val int) { fmt.Printf("len of string, %v", val) }).Get()
	if err != nil {
		log.Println("error fetching content")
	}
}
