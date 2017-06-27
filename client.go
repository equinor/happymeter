package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
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
	var jsonStr = []byte(`{"happystatus":"` + status + `","tags":"` + h.Tags + `","comment":"` + h.Comment + `"}`)
	req, err := http.NewRequest("POST", h.Endpoint, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("response status:", resp.Status)
	fmt.Println("response headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response body:", string(body))
}
