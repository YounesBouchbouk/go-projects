package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type RSS struct {
	XMLName xml.Name `xml:"rss"`
	Channel *Channel `xml:"channel"`
}

type Channel struct {
	Title     string `xml:"title"`
	ItemsList []Item `xml:"item"`
}

type Item struct {
	Title     string `xml:"title"`
	Link      string `xml:"link"`
	Traffic   string `xml:"approx_traffic"`
	NewsItems []News `xml:"news_item"`
}

type News struct {
	HeadLine     string `xml:"news_item_title"`
	HeadLineLink string `xml:"news_item_url"`
}

func main() {
	var r RSS

	data := readGoogleTreands()

	err := xml.Unmarshal(data, &r)

	if err != nil {
		fmt.Println("error :", err)
		os.Exit(1)
	}

	fmt.Println("\n bellow are all the google Search Treands For Today")
	fmt.Println("................................................")

	for i := range r.Channel.ItemsList {
		rank := (i + 1)
		fmt.Println("#", rank)
		fmt.Println("Search Term:", r.Channel.ItemsList[i].Title)
		fmt.Println("Link to trend:", r.Channel.ItemsList[i].Link)
		fmt.Println("Headline:", r.Channel.ItemsList[i].NewsItems[0].HeadLine)
		fmt.Println("Link to article:", r.Channel.ItemsList[i].NewsItems[0].HeadLineLink)
		fmt.Println("--------------------------")
	}

}

func getGoogleTrends() *http.Response {
	resp, err := http.Get("https://trends.google.com/trends/trendingsearches/daily/rss?geo=US")

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return resp
}

func readGoogleTreands() []byte {
	resp := getGoogleTrends()

	data, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return data

}
