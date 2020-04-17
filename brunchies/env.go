package main

import (
	"os"
)

type Env struct {
	SPOTIFY_ID     string
	SPOTIFY_SECRET string
}

func (e *Env) Load() {
	e.SPOTIFY_ID = os.Getenv("SPOTIFY_ID")
	e.SPOTIFY_SECRET = os.Getenv("SPOTIFY_SECRET")
}
