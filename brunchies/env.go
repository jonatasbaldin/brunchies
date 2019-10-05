package main

import (
	"io/ioutil"
	"os"
)

type Env struct {
	SPOTIFY_ID     string
	SPOTIFY_SECRET string
}

func (e *Env) Load() {
	spotifyIdContent, err := ioutil.ReadFile("/var/openfaas/secrets/spotify_id")
	if err == nil {
		e.SPOTIFY_ID = string(spotifyIdContent)
	} else {
		e.SPOTIFY_ID = os.Getenv("SPOTIFY_ID")
	}

	spotifySecretContent, err := ioutil.ReadFile("/var/openfaas/secrets/spotify_secret")
	if err == nil {
		e.SPOTIFY_SECRET = string(spotifySecretContent)
	} else {
		e.SPOTIFY_SECRET = os.Getenv("SPOTIFY_SECRET")
	}
}
