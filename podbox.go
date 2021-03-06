package main

import (
	"log"
	"sort"
)

type (
	Track struct {
		Title      string `json:"title"`
		Artist     string `json:"artist"`
		Rank       int    `json:"rank"`
		SpotifyUri string `json:"spotifyUri,omitempty"`
	}
	Tracks []Track
)

func NewTrack(item Entry, url string) Track {
	return Track{
		Title:      item.Title,
		Artist:     item.Artist,
		Rank:       item.Rank,
		SpotifyUri: url,
	}
}

func (ts Tracks) Len() int {
	return len(ts)
}
func (ts Tracks) Swap(i, j int) {
	ts[i], ts[j] = ts[j], ts[i]
}
func (ts Tracks) Less(i, j int) bool {
	return ts[i].Rank < ts[j].Rank
}

func getHot10() Tracks {
	tracks := make(Tracks, 0)

	feed, err := billboard()
	if err != nil {
		log.Println(err)
		return tracks
	}

	size := 10
	if len(feed.Items) < size {
		size = len(feed.Items)
	}

	input := make(chan Entry, size)
	output := make(chan Track, size)

	for i := 0; i < 2; i++ {
		go worker(input, output)
	}

	for _, item := range feed.Items[:size] {
		input <- item
	}
	close(input)

	for i := 0; i < size; i++ {
		tracks = append(tracks, <-output)
	}
	close(output)

	sort.Sort(tracks)
	return tracks
}

func worker(input <-chan Entry, output chan<- Track) {
	for item := range input {
		url := previewUrl(item.Title, item.Artist)
		output <- NewTrack(item, url)
	}
}
