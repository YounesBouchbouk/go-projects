package cli

import (
	"log"
	"net/http"
	"sync"

	"github.com/Jeffail/gabs"
)

type RequestBody struct {
	SourceLang string
	TargetLang string
	SourceText string
}

const translateUrl = "https://translate.googleapis.com/translate_a/single"

func RequestTranslate(body *RequestBody, str chan string, wg *sync.WaitGroup) {

	//we have created a client in this step that will be used later on
	client := &http.Client{}

	req, err := http.NewRequest("Get", translateUrl, nil)

	if err != nil {

	}

	query := req.URL.Query()

	//you have received the body object in this function, it has been sent from main func
	// it will have sourceLang and targetLang already set, also a source text and that's
	//what we're accessing here and adding to the query
	query.Add("sl", body.SourceLang)
	query.Add("tl", body.TargetLang)
	query.Add("dt", "t")
	query.Add("q", body.SourceText)

	//once the query is created, we add it to the req.URL as RawQuery after encoding it in
	//JSON as JSON queries need to go
	req.URL.RawQuery = query.Encode()

	//this is the place where the request is actually made, till now we
	//were just creating the request
	res, err := client.Do(req)
	if err != nil {
		log.Fatalf("2 There was a problem: %s", err)
	}

	//you get response and you close the response at the end of this funcn using defer
	defer res.Body.Close()

	//you may get blocked if there are too many requests because golang can make
	//loads of requests in very less time
	if res.StatusCode == http.StatusTooManyRequests {
		str <- "You have been rate limited, Try again later."
		wg.Done()
		return
	}

	//you want to parse the json using gabs package
	parsedJson, err := gabs.ParseJSONBuffer(res.Body)
	//and then handle the error
	if err != nil {
		log.Fatalf("3 There was a problem - %s", err)
	}

	//get the nested elements at oth root of parsedJson variable
	nestOne, err := parsedJson.ArrayElement(0)
	if err != nil {
		log.Fatalf("4 There was a problem - %s", err)
	}

	//get one level deeper nested element
	nestTwo, err := nestOne.ArrayElement(0)
	if err != nil {
		log.Fatalf("5 There was a problem - %s", err)
	}

	//the translated string comes deep within so we have to extract it
	//we've put checks at each de-nesting stage with different numbers

	translatedStr, err := nestTwo.ArrayElement(0)
	if err != nil {
		log.Fatalf("6 There was a problem - %s", err)
	}

	str <- translatedStr.Data().(string)
	wg.Done()

}
