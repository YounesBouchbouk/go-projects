package extract

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
)

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

type URLStruct struct {
	link    string
	country string
}

func ReadGoogleTrends(ch chan<- []byte) {

	var wg sync.WaitGroup

	urlList := []URLStruct{
		{
			link:    "https://trends.google.com/trends/trendingsearches/daily/rss?geo=GB",
			country: "GB",
		},
		{
			link:    "https://trends.google.com/trends/trendingsearches/daily/rss?geo=FR",
			country: "FR",
		},
		{
			link:    "https://trends.google.com/trends/trendingsearches/daily/rss?geo=FR",
			country: "FR",
		},
		{
			link:    "https://trends.google.com/trends/trendingsearches/daily/rss?geo=CA",
			country: "CA",
		},
	}

	for _, url := range urlList {
		wg.Add(1)

		go getGoogleTrends(url, ch, &wg)
	}

	//wait for all go routine to finish to close the channel
	// wg.Wait()
	close(ch)

	// //close the channel
}

func getGoogleTrends(url URLStruct, ch chan<- []byte, wg *sync.WaitGroup) {
	defer wg.Done()

	resp, err := http.Get(url.link)

	if err != nil {
		fmt.Println("error in Get request from getGoogleTrends")
	}

	defer resp.Body.Close()

	data, err := extractBytesFromresponse(resp)

	if err != nil {
		fmt.Println("erro extracting data Bytes")
	}

	ch <- data

}

func extractBytesFromresponse(resp *http.Response) ([]byte, error) {

	data, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return data, nil

}
