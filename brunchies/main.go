package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"text/template"
)

func index(w http.ResponseWriter, r *http.Request, client SpotifyClient) {
	t, err := template.ParseFiles("index.html")
	if err != nil {
		log.Println(err)
	}

	playlistMetadata := getPlaylistMetadata(client)
	weeks := getWeeksWithTracks(client)

	data := Data{
		PlaylistMetadata: playlistMetadata,
		Weeks:            weeks,
	}

	err = t.Execute(w, data)
	if err != nil {
		log.Println(err)
	}
}

func main() {
	config := NewSpotifyConfig()
	client := NewSpotifyClient(config)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		index(w, r, client)
	})

	err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	if err != nil {
		log.Fatal(err)
	}
}
