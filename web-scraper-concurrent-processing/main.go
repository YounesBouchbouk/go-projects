package main

import (
	"encoding/xml"
	"fmt"

	"github.com/YounesBouchbouk/web-scraper-processing/extract"
)

type RSS struct {
	XMLName xml.Name         `xml:"rss"`
	Channel *extract.Channel `xml:"channel"`
}

func main() {

	var r RSS

	channel := make(chan []byte)

	extract.ReadGoogleTrends(channel)

	for {
		select {
		case data := <-channel:
			go func() {
				err := xml.Unmarshal(data, &r)

				if err != nil {
					fmt.Println("error :", err)
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

			}()
		}
	}

}
