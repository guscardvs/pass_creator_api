package main

import (
	"crypto/rand"
	"encoding/json"
	"io/ioutil"
	"math/big"
	mrand "math/rand"
	"net/http"
	"strings"
	"time"
)

var LETTERS = "abcdefghijklmnopqrstuvwxyz"
var PASSPHRASE_SIZE = 5

type Word struct {
	Value string `json:"word"`
}

type PassPhrase struct {
	Value string `json:"passphrase"`
}

func randomLetter() byte {
	return LETTERS[generateRandomInt(25)]
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

func WordRequest(newWords chan<- []Word) {
	var url = "http://api.datamuse.com/words?sp=" + string(randomLetter()) + "*"
	req, _ := http.NewRequest("GET", url, nil)
	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	var words []Word
	json.Unmarshal(body, &words)
	newWords <- words
}

func GetPassPhrase() (passPhrase PassPhrase) {
	var words []string
	var newWords chan []Word = make(chan []Word)
	for i := 0; i < PASSPHRASE_SIZE; i++ {
		go WordRequest(newWords)
	}
	for i := 0; i < PASSPHRASE_SIZE; i++ {
		words = append(words, chooseRandomWord(<-newWords).Value)
	}
	mrand.Seed(time.Now().UnixNano())
	mrand.Shuffle(len(words), func(i, j int) { words[i], words[j] = words[j], words[i] })
	passPhrase = PassPhrase{strings.Join(words, " ")}
	return
}
