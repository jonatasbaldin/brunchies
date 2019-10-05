package function

import (
	"context"
	"log"
	"sort"
	"time"

	"github.com/zmb3/spotify"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

const BRUNCHIES_ID spotify.ID = "4OyKDT6cLw96G7bd8nTfxD"

type SpotifyClient interface {
	GetPlaylist(spotify.ID) (*spotify.FullPlaylist, error)
	GetPlaylistTracks(spotify.ID) (*spotify.PlaylistTrackPage, error)
}

type SpotifyConfig interface {
	Token(ctx context.Context) (*oauth2.Token, error)
}

func NewSpotifyConfig() SpotifyConfig {
	env := Env{}
	env.Load()

	config := &clientcredentials.Config{
		ClientID:     env.SPOTIFY_ID,
		ClientSecret: env.SPOTIFY_SECRET,
		TokenURL:     spotify.TokenURL,
	}

	return config
}

func NewSpotifyClient(config SpotifyConfig) SpotifyClient {
	token, err := config.Token(context.Background())
	if err != nil {
		log.Fatalf("couldn't get token: %v", err)
	}

	client := spotify.Authenticator{}.NewClient(token)

	return &client
}

func getPlaylistMetadata(client SpotifyClient) PlaylistMetadata {
	playlist, err := client.GetPlaylist(BRUNCHIES_ID)
	if err != nil {
		log.Fatalf("couldn't get playlist: %s", err)
	}

	return PlaylistMetadata{
		Author:    playlist.Owner.DisplayName,
		Followers: playlist.Followers.Count,
		URL:       playlist.ExternalURLs["spotify"],
		ImageURL:  playlist.SimplePlaylist.Images[0].URL,
	}
}

func getWeeksWithTracks(client SpotifyClient) []Week {
	page, err := client.GetPlaylistTracks(BRUNCHIES_ID)
	if err != nil {
		log.Fatalf("couldn't get features playlists: %s", err)
	}

	var weeks []Week

	for _, track := range page.Tracks {
		artists := formatTrackArtists(track.Track)

		addedAt, _ := time.Parse(time.RFC3339, track.AddedAt)
		trackYear, trackWeek := addedAt.ISOWeek()

		track := Track{
			Name:    track.Track.Name,
			Artists: artists,
		}

		week := Week{
			Year: trackYear,
			Week: trackWeek,
		}

		if len(weeks) == 0 {
			weeks = append(weeks, week)
		}

		lastAddedWeek := weeks[len(weeks)-1]
		if !lastAddedWeek.Equal(week) {
			weeks = append(weeks, week)
		}

		if trackYear == week.Year && trackWeek == week.Week {
			specificWeek := getSpecificWeek(weeks, week)
			weeks[specificWeek].Tracks = append(weeks[specificWeek].Tracks, track)
		}
	}

	sort.Slice(weeks, func(i, j int) bool {
		return weeks[j].Week < weeks[i].Week
	})

	return weeks
}

func formatTrackArtists(track spotify.FullTrack) (artists string) {
	if len(track.Artists) > 1 {
		for i, artist := range track.Artists {
			if len(track.Artists)-1 > i {
				artists += artist.Name + ", "
			} else {
				artists += artist.Name
			}
		}

	} else {
		artists = track.Artists[0].Name
	}

	return artists
}
