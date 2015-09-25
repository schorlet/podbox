package main

import "log"

type Track struct {
	Title      string `json:"title"`
	Artist     string `json:"artist"`
	Rank       int    `json:"rank"`
	SpotifyUri string `json:"spotifyUri,omitempty"`
}

func NewTrack(item Entry, url string) Track {
	return Track{
		Title:      item.Title,
		Artist:     item.Artist,
		Rank:       item.Rank,
		SpotifyUri: url,
	}
}

func getHot10() []Track {
	tracks := make([]Track, 0)

	feed, err := billboard()
	if err != nil {
		log.Println(err)
		return tracks
	}

	max := 10
	if len(feed.Items) < 10 {
		max = len(feed.Items)
	}

	for _, item := range feed.Items[:max] {
		url := previewUrl(item.Title, item.Artist)
		tracks = append(tracks, NewTrack(item, url))
	}
	return tracks
}
