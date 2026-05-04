package main 

import (
	"net/http"
	"encoding/xml"
	"io"
	"html"
	"context"
)

func fetchFeed (ctx context.Context, feedURL string) (*RSSFeed, error){
	// create a request to call the webserver
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return nil, err
	}
	// adjust the header to identify as gator 
	req.Header.Set("User-Agent", "gator")
	// initialize the client to post the request 
	client := &http.Client{}
	// makle the client call to get response 
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	// defer to close response body until return
	defer res.Body.Close()
	// get data from the response body 
	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	// unmarshal data into RSSFeed struct 
	RF := &RSSFeed{}
	if err = xml.Unmarshal(data, RF); err !=nil {
		return nil, err
	}
	// unescape the html strings to ensure readability
	RF.Channel.Title = html.UnescapeString(RF.Channel.Title)
	RF.Channel.Description = html.UnescapeString(RF.Channel.Description)
	for i := range RF.Channel.Item {
		RF.Channel.Item[i].Title = html.UnescapeString(RF.Channel.Item[i].Title)
		RF.Channel.Item[i].Description = html.UnescapeString(RF.Channel.Item[i].Description)
	}
	// return the unmarshalled and unescaped struct
	return RF, nil
}


type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}
