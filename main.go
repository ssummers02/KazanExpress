package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

type Response struct {
	IconURL string `json:"icon_url"`
	ID      string `json:"id"`
	URL     string `json:"url"`
	Value   string `json:"value"`
}

func getjoke(string string) (string, error) {
	resp, err := http.Get(string)
	if err != nil {
		log.Println(err)
		return "", err
	}
	defer resp.Body.Close()
	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}
	var chucknorris Response
	err = json.Unmarshal(body, &chucknorris)
	if err != nil {
		log.Println(err)
		return "", err
	}
	return chucknorris.Value, nil
}

func random() {
	s, _ := getjoke("https://api.chucknorris.io/jokes/random")

	fmt.Println(s)
}

func dump(count int) {
	rand.Seed(time.Now().UnixNano())
	Categorirs := []string{"animal", "career", "celebrity", "dev", "explicit", "fashion", "food", "history", "money", "movie", "music", "political", "religion", "science", "sport", "travel"}

	for i := 0; i < count; i++ {
		r := rand.Intn(len(Categorirs))
		fileName := Categorirs[r] + ".txt"

		str, err := getjoke("https://api.chucknorris.io/jokes/random?category=" + Categorirs[r])
		if err != nil {
			log.Println("get joke error", err)
			return
		}
		f, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Println(err)
		}
		defer f.Close()
		if _, err := f.WriteString(str + "\n"); err != nil {
			log.Println(err)
		}
	}
}
func main() {
	argsWithProg := os.Args

	if len(argsWithProg) < 2 || len(argsWithProg) > 4 {
		log.Println("error arg")
		return
	}
	if argsWithProg[1] == "random" {
		random()
	}
	if argsWithProg[1] == "dump" {
		count := 5
		if argsWithProg[2] != "-n" {
			log.Println("error arg n")
			return
		}
		count, err := strconv.Atoi(argsWithProg[3])
		if err != nil {
			log.Println("error arg n")
			return
		}
		dump(count)
	}
}
