package main

import (
	"fmt"
	"github.com/stianeikeland/go-rpio"
	yaml "gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"net"
	"os"
	"time"
)

var (
	btn1   = rpio.Pin(13)
	btn2   = rpio.Pin(5)
	btn3   = rpio.Pin(15)
	client *Client
)

type Config struct {
	HappyMeter struct {
		Url  string `yaml:"url"`
		Tags string `yaml:"tags"`
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Error")
		fmt.Println(err)
		panic(err)
	}
}

func getConfig(file string) (*Config, error) {
	config := &Config{}
	configData, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	err = yaml.Unmarshal(configData, &config)
	if err != nil {
		log.Fatalf("error: %v", err)
		return nil, err
	}

	return config, nil
}

func getMyIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Fatal(err)
		return ""
	}

	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			return ipnet.IP.String()
		}
	}
	return "n/a"
}

func main() {
	var config *Config
	fmt.Println("=== Happymeter ===")

	config, err := getConfig("config.yml")
	if err != nil {
		panic(err)
	}

	if err := rpio.Open(); err != nil {
		fmt.Println(err)
		fmt.Println("Need to run as root, and device must be an Raspberry Pi!")
		os.Exit(1)
	}

	defer rpio.Close()

	comment := "Happymeter device, ip: " + getMyIP()
	client, err := NewClient(config.HappyMeter.Url, config.HappyMeter.Tags, comment)
	if err != nil {
		panic(err)
	}

	btn1.Input()
	btn2.Input()
	btn3.Input()

	for {
		if val := btn1.Read(); val == rpio.Low {
			fmt.Println("Button happy pressed!")
			go client.Post("happy")
		} else if val := btn2.Read(); val == rpio.Low {
			fmt.Println("Button medium pressed!")
			go client.Post("medium")
		} else if val := btn3.Read(); val == rpio.Low {
			fmt.Println("Button sad pressed!")
			go client.Post("sad")
		}

		time.Sleep(time.Second / 5)
	}
}
