package main

import (
	"os"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"time"
	"strings"
)

type Word struct {
	Id          string `json:"int"`
	Slug        string `json:"slug"`
	Description []Description `json:"back"`
}

type Description struct {
	Type    string `json:"type"`
	Content string `json:"content"`
}

func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max - min) + min
}

func main() {
	configFile, err := os.Open("/Users/sumankum/Projects/src/goutils/vocabulary/words.json")
	if err != nil {
		log.Println("opening words file", err.Error())
	}

	jsonParser := json.NewDecoder(configFile)
	var vocabulary = []Word{}
	if err = jsonParser.Decode(&vocabulary); err != nil {
		log.Println("parsing config file", err.Error())
	}
	index := random(0, 1000)

	fmt.Println("======================================================================================================")
	fmt.Printf("           word: %s\n", vocabulary[index].Slug)
	fmt.Printf("           meaning: %s\n", strings.Replace(strings.Replace(vocabulary[index].Description[1].Content, "<strong>", "", -1), "</strong>", "", -1))
	fmt.Println("======================================================================================================")
}
