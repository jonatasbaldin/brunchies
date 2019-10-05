package function

import (
	"log"
	"net/http"
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

func Handle(w http.ResponseWriter, r *http.Request) {
	config := NewSpotifyConfig()
	client := NewSpotifyClient(config)

	// port := os.Getenv("PORT")
	// if port == "" {
	// 	port = "8080"
	// }

	index(w, r, client)

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	index(w, r, client)
	// })

	// err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }
}
