package main

import (
	"encoding/xml"
	"fmt"
	"net/http"
)

type (
	Feed struct {
		// XMLName xml.Name `xml:"rss"`
		Items []Entry `xml:"channel>item"`
	}

	Entry struct {
		Title  string `xml:"chart_item_title"`
		Artist string `xml:"artist"`
		Rank   int    `xml:"rank_this_week"`
	}
)

func billboard() (*Feed, error) {
	url := "http://www.billboard.com/rss/charts/hot-100"
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("billboard: %s", resp.Status)
	}

	feed := new(Feed)
	dec := xml.NewDecoder(resp.Body)

	if err := dec.Decode(feed); err != nil {
		return nil, err
	}
	return feed, nil
}
