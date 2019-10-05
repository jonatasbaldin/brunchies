package function

import (
	"context"
	"testing"

	"github.com/zmb3/spotify"
	"golang.org/x/oauth2"
)

type Config struct {
	ClientID     string
	ClientSecret string
	TokenURL     string
}

func (c *Config) Token(ctx context.Context) (*oauth2.Token, error) {
	return &oauth2.Token{
		AccessToken: "ble",
	}, nil
}

type Client struct{}

func (c *Client) GetPlaylist(spotify.ID) (*spotify.FullPlaylist, error) {
	user := spotify.User{
		DisplayName: "Max",
	}

	followers := spotify.Followers{
		Count: 1,
	}

	simplePlaylist := spotify.SimplePlaylist{
		Owner:        user,
		ExternalURLs: map[string]string{"spotify": "url"},
		Images: []spotify.Image{
			{
				URL: "url",
			},
		},
	}

	return &spotify.FullPlaylist{
		SimplePlaylist: simplePlaylist,
		Followers:      followers,
	}, nil
}

func (c *Client) GetPlaylistTracks(spotify.ID) (*spotify.PlaylistTrackPage, error) {
	playlistTrack1 := spotify.PlaylistTrack{
		AddedAt: "2016-10-11T13:44:40Z",
		Track: spotify.FullTrack{
			SimpleTrack: spotify.SimpleTrack{
				Name: "track #1",
				Artists: []spotify.SimpleArtist{
					{
						Name: "Artist One",
					},
				},
			},
		},
	}

	playlistTrack2 := spotify.PlaylistTrack{
		AddedAt: "2017-12-11T13:44:40Z",
		Track: spotify.FullTrack{
			SimpleTrack: spotify.SimpleTrack{
				Name: "track #2",
				Artists: []spotify.SimpleArtist{
					{
						Name: "Artist One",
					},
					{
						Name: "Artist Two",
					},
				},
			},
		},
	}

	return &spotify.PlaylistTrackPage{
		Tracks: []spotify.PlaylistTrack{
			playlistTrack1,
			playlistTrack2,
		},
	}, nil
}

func Test_NewSpotifyConfig(t *testing.T) {
	config := NewSpotifyConfig()

	// Gets interface of c
	var i interface{} = config

	// Verify if i implements Client interface
	_, ok := i.(SpotifyConfig)
	if !ok {
		t.Errorf("%v does not implement SpotifyClient interface", config)
	}
}

func Test_NewSpotifyClient(t *testing.T) {
	config := Config{
		ClientID:     "SPOTIFY_ID",
		ClientSecret: "SPOTIFY_TOKEN",
		TokenURL:     spotify.TokenURL,
	}

	client := NewSpotifyClient(&config)

	// Gets interface of c
	var i interface{} = client

	// Verify if i implements SpotifyClient interface
	_, ok := i.(SpotifyClient)
	if !ok {
		t.Errorf("%v does not implement Client interface", client)
	}
}

func Test_GetPlaylistMetadata(t *testing.T) {
	client := Client{}
	playlistMetadata := getPlaylistMetadata(&client)

	if playlistMetadata.Author != "Max" {
		t.Errorf("expected %s, got %s", "Max", playlistMetadata.Author)
	}

	if playlistMetadata.Followers != 1 {
		t.Errorf("expected %d, got %d", 1, playlistMetadata.Followers)
	}

	if playlistMetadata.URL != "url" {
		t.Errorf("expected %s, got %s", "url", playlistMetadata.URL)
	}

	if playlistMetadata.ImageURL != "url" {
		t.Errorf("expected %s, got %s", "url", playlistMetadata.ImageURL)
	}
}

func Test_GetWeeksWithTracks(t *testing.T) {
	client := Client{}
	weeks := getWeeksWithTracks(&client)

	if weeks[0].Year != 2017 {
		t.Errorf("expected %d, got %d", 2018, weeks[0].Year)
	}

	if weeks[0].Week != 50 {
		t.Errorf("expected %d, got %d", 11, weeks[0].Week)
	}

	if weeks[0].Tracks[0].Name != "track #2" {
		t.Errorf("expected %s, got %s", "track #2", weeks[0].Tracks[0].Name)
	}

	if weeks[0].Tracks[0].Artists != "Artist One, Artist Two" {
		t.Errorf("expected %s, got %s", "Artist One, Artist Two2", weeks[0].Tracks[0].Artists)
	}

	if weeks[1].Year != 2016 {
		t.Errorf("expected %d, got %d", 2016, weeks[1].Year)
	}

	if weeks[1].Week != 41 {
		t.Errorf("expected %d, got %d", 41, weeks[1].Week)
	}

	if weeks[1].Tracks[0].Name != "track #1" {
		t.Errorf("expected %s, got %s", "track #1", weeks[1].Tracks[0].Name)
	}

	if weeks[1].Tracks[0].Artists != "Artist One" {
		t.Errorf("expected %s, got %s", "Artist One", weeks[1].Tracks[0].Artists)
	}
}

func Test_FormatArtists(t *testing.T) {
	track := spotify.FullTrack{
		SimpleTrack: spotify.SimpleTrack{
			Name: "track #1",
			Artists: []spotify.SimpleArtist{
				{
					Name: "Artist One",
				},
				{
					Name: "Artist Two",
				},
			},
		},
	}

	artists := formatTrackArtists(track)

	if artists != "Artist One, Artist Two" {
		t.Errorf("expected %s, got %s", "Artist One, Artist Two2", artists)
	}

}
