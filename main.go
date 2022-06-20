package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"

	gabs "github.com/Jeffail/gabs/v2"
)

var (
	leftHandWordsURL  = "https://fly.wordfinderapi.com/api/search?letters=qwertasdfgzxcvb&page_token=1&page_size=2000&word_sorting=points&group_by_length=false&dictionary=wwf2"
	rightHandWordsURL = "https://fly.wordfinderapi.com/api/search?letters=yuiophjklnm&page_token=1&page_size=2000&word_sorting=points&group_by_length=false&dictionary=wwf2"
	flagLeft          = flag.Bool("l", false, "Left hand words")
	flagRight         = flag.Bool("r", false, "right hand words")
	flagWordCount     = flag.Int("n", 3, "Number of words to generate in pw")
)

func getKeyboardWords(random bool, list bool) (returnedWord string, wordList []string) {
	flag.Parse()
	var wordsURL string
	if *flagLeft == true {
		wordsURL = leftHandWordsURL

	}
	if *flagRight == true {
		wordsURL = rightHandWordsURL
	}

	if !*flagLeft && !*flagRight {
		wordsURL = leftHandWordsURL
	}
	
	WordsResponse, err := http.Get(wordsURL)
	if err != nil {
		log.Fatal(err)
	}
	bodyBytes, err := ioutil.ReadAll(WordsResponse.Body)
	parsedResponses, err := gabs.ParseJSON([]byte(bodyBytes))
	if err != nil {
		panic("no words")
	}
	Words := parsedResponses.Path("word_pages.0.word_list").Children()
	for _, i := range Words {
		wordList = append(wordList, i.Path("word").Data().(string))
	}
	if random {
		rand.Seed(time.Now().Unix())
		newArray := []string{}
		for i := 0; i < *flagWordCount; i++ {
			newArray = append(newArray, wordList[rand.Intn(len(wordList))])
		}
		return wordList[rand.Intn(len(wordList))], newArray
	} else {
		return "", wordList
	}
}

func main() {
	_, wordList := getKeyboardWords(true, false)
	fmt.Println(strings.Join(wordList[:], "-"))
}
