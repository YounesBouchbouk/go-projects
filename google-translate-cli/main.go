package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/Younesbouchbouk/google-translate-cli/cli"
)

var sourceLang string
var targetLang string
var sourceText string

func init() {
	flag.StringVar(&sourceLang, "s", "en", "Source Language [en]")
	flag.StringVar(&targetLang, "t", "fr", "Target Language [fr]")
	flag.StringVar(&sourceText, "st", "", "Text To translate")

}

func main() {
	flag.Parse()

	if flag.NFlag() == 0 {
		//if zero flags have been set, we will show the usage options
		fmt.Println("Options:")
		// printing out the default values set above
		flag.PrintDefaults()
		os.Exit(1)
	}

	strChan := make(chan string)

	reqBody := &cli.RequestBody{
		SourceLang: sourceLang,
		TargetLang: targetLang,
		SourceText: sourceText,
	}

	var wg sync.WaitGroup

	wg.Add(1)

	go cli.RequestTranslate(reqBody, strChan, &wg)

	processedStr := strings.ReplaceAll(<-strChan, " + ", " ")

	fmt.Printf("%s\n", processedStr)

	close(strChan)
	wg.Wait()

}
