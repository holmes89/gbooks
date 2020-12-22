package client

import (
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
)

const baseURL = "https://www.googleapis.com/books/v1/volumes"

type Results struct {
	Total int `json:"totalItems"`
	Items []Item
}

type Item struct {
	ID         string `json:"id"`
	VolumeInfo struct {
		Title         string   `json:"title"`
		Authors       []string `json:"authors"`
		PublishedDate string   `json:"publishedDate"`
	} `json:"volumeInfo"`
}

func SearchByTitle(title string) (int, []Item, error) {

	searchTerm := fmt.Sprintf("intitle:%s", title)
	return search(searchTerm)
}

func SearchByAuthor(name string) (int, []Item, error) {

	searchTerm := fmt.Sprintf("inauthor:%s", name)
	return search(searchTerm)
}

func SearchByISBN(isbn string) (int, []Item, error) {

	searchTerm := fmt.Sprintf("isbn:%s", isbn)
	return search(searchTerm)
}

func search(searchTerm string) (int, []Item, error) {
	resp, err := resty.New().
		SetQueryParam("q", searchTerm).
		SetQueryParam("maxResults", "40").
		R().Get(baseURL)
	if err != nil {
		return 0, nil, err
	}

	if err := resp.Error(); err != nil {
		return 0, nil, fmt.Errorf("%+v", err)
	}

	var results Results
	if err := json.Unmarshal(resp.Body(), &results); err != nil {
		return 0, nil, err
	}

	return results.Total, results.Items, nil
}
