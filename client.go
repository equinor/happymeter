package main

import (
	"log"
	"net/http"
	"net/url"
	"strings"
)

type Client struct {
	Endpoint string
	Tags     string
	Comment  string
}

func NewClient(endpoint, tags, comment string) (*Client, error) {
	happyClient := &Client{
		Endpoint: endpoint,
		Tags:     tags,
		Comment:  comment}

	return happyClient, nil
}

func (h *Client) Post(status string) {
	data := url.Values{}
	data.Add("happystatus", status)
	data.Add("tags", h.Tags)
	data.Add("comment", h.Comment)

	req, err := http.NewRequest("POST", h.Endpoint, strings.NewReader(data.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Unable to post to happymeter:%v\n", err)
		return
	}
	defer resp.Body.Close()
}
