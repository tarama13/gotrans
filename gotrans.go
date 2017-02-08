package main

import (
	"fmt"
	"os"
	"github.com/PuerkitoBio/goquery"
)

func PrintResult(word string) {
	url := "http://ejje.weblio.jp/content/" + word
	doc, _ := goquery.NewDocument(url)
	result := doc.Find(".content-explanation").Text()
	if len(result) == 0 {
		fmt.Printf("見つかりませんでした...\n")
		os.Exit(0)
	}
	fmt.Printf("%s\n", result)
}

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("usage: %s [English word]\n", os.Args[0]);
		os.Exit(0)
	}

	PrintResult(os.Args[1])
}
