package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"

	tb "gopkg.in/tucnak/telebot.v2"
)

var userID string
var telegramToken string
var geoApiKey string

type IPloc struct {
	City         string `json:"city"`
	Country_name string `json:"country_name"`
}

func getJson(url string, target interface{}) error {
	r, err := http.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

func main() {
	userIDint, _ := strconv.Atoi(userID)
	args := os.Args[1:]
	b, err := tb.NewBot(tb.Settings{
		Token: telegramToken,
	})

	if err != nil {
		log.Fatal(err)
		return
	}
	if args[0] == "SSH" {
		b.Send(&tb.User{ID: userIDint}, "SSH login on "+args[1]+" ; IP: "+args[2])

		// get IP location
		a := &IPloc{}
		getJson("http://api.ipstack.com/"+args[2]+"?access_key="+geoApiKey+"&format=1", &a)

		b.Send(&tb.User{ID: userIDint}, "IP location: "+a.City+", "+a.Country_name)

	}

}
