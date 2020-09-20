package main

import (
	"fmt"

	"github.com/mmcdole/gofeed"
)

func main() {
	ArticleNum := 1

	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL("http://feeds.twit.tv/twit.xml")
	fmt.Println("Feed title:" + feed.Title)
	fmt.Println("Description:" + feed.Description + "\n")
	fmt.Println("Last update:" + feed.Updated + "\n")

	for _, Item := range feed.Items {
		fmt.Printf("Article number: %d \n", ArticleNum)
		ArticleNum++
		fmt.Println("Article title: " + Item.Title)
		fmt.Println("Article Updated: " + Item.Updated)
		fmt.Println("Article Published: " + Item.Published)
		fmt.Println("Article Description: " + Item.Description)
		fmt.Println("Article Content: " + Item.Content)
		fmt.Println("Article Categories: ")
		fmt.Println(Item.Categories)
		fmt.Println("---\n")
	}

	fmt.Printf("\n%d Article(s) downloaded", ArticleNum)

}
