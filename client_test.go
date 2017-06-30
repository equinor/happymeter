package main

import "testing"

func TestNewClient(t *testing.T) {
	url := "http://testhappymeter.cloudapp.net/api/storehappydocument"
	comment := "Happymeter device, ip: " + getMyIP()
	tags := "testing"
	_, err := NewClient(url, tags, comment)
	if err != nil {
		panic(err)
	}
}
