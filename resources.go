package main

import (
	"crypto/rand"
	"encoding/json"
	"io/ioutil"
	"math/big"
	"net/http"
	"strings"
)

var LETTERS = "abcdefghijklmnopqrstuvwxyz"
var PASSPHRASE_SIZE = 5

type Word struct {
	Value string `json:"word"`
}

type PassPhrase struct {
	Value string `json:"passphrase"`
}

func newWord(value string) Word {
	return Word{value}
}

func randomLetter() byte {
	return LETTERS[generateRandomInt(25)]
}

func WordRequest() []Word {
	var url = "http://api.datamuse.com/words?sp=" + string(randomLetter()) + "*"
	req, _ := http.NewRequest("GET", url, nil)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	var words []Word
	json.Unmarshal(body, &words)

	return words
}

func generateRandomInt(max int) int {
	bmax := big.NewInt(int64(max))
	response, _ := rand.Int(rand.Reader, bmax)
	return int(response.Int64())
}

func chooseRandomWord(words []Word) Word {
	var randomInt = generateRandomInt(len(words) - 1)

	return words[randomInt]
}

func GetPassPhrase() PassPhrase {
	var words []string

	for i := 0; i < PASSPHRASE_SIZE; i++ {
		newWords := WordRequest()
		words = append(words, chooseRandomWord(newWords).Value)
	}
	return PassPhrase{strings.Join(words, " ")}
}
