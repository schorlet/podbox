package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type (
	Result struct {
		Tracks struct {
			Items []Item `json:"items"`
		} `json:"tracks"`
	}

	Item struct {
		PreviewURL string `json:"preview_url"`
	}
)

func previewUrl(track, artist string) string {
	res, err := search(track, artist)
	if err != nil {
		return ""
	}

	tracks := res.Tracks.Items
	if len(tracks) == 0 {
		return ""
	}
	return tracks[0].PreviewURL
}

func search(track, artist string) (*Result, error) {
	v := url.Values{}
	v.Set("q", fmt.Sprintf(`track:"%s" artist:"%s"`, track, artist))
	v.Set("type", "track")
	v.Set("limit", "1")

	url := "https://api.spotify.com/v1/search?" + v.Encode()
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("spotify: %s", resp.Status)
	}

	res := new(Result)
	dec := json.NewDecoder(resp.Body)

	if err := dec.Decode(res); err != nil {
		return nil, err
	}
	return res, nil
}
