package main

import (
	"fmt"
	"github.com/jessevdk/go-flags"
	"github.com/stianeikeland/go-rpio"
	"log"
	"net"
	"os"
	"time"
)

type Options struct {
	Config string `short:"c" long:"config" description:"Provide a path to the input config file" required:"true"`
}

var options Options
var parser = flags.NewParser(&options, flags.Default)

func checkError(err error) {
	if err != nil {
		fmt.Println("Error")
		fmt.Println(err)
		panic(err)
	}
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

type Button struct {
	rpio.Pin
	Pressed bool
}

var (
	btn1   = Button{rpio.Pin(13), false}
	btn2   = Button{rpio.Pin(5), false}
	btn3   = Button{rpio.Pin(27), false}
	client *Client
)

func main() {
	log.Println("=== Happymeter ===")

	if _, err := parser.Parse(); err != nil {
		return
	}

	var config *Config
	config, err := ReadConfig(options.Config)
	if err != nil {
		panic(err)
	}

	if err := rpio.Open(); err != nil {
		log.Println(err)
		log.Println("Need to run as root, and device must be an Raspberry Pi!")
		os.Exit(1)
	}

	defer rpio.Close()

	comment := "Happymeter device, ip: " + getMyIP()
	client, err := NewClient(config.HappyMeter.Url, config.HappyMeter.Tags, comment)
	if err != nil {
		panic(err)
	}

	log.Println("My ip:", getMyIP())
	log.Println("Endpoint url:", config.HappyMeter.Url)

	btn1.Input()
	btn2.Input()
	btn3.Input()

	var count = 0
	for {
		if val := btn1.Read(); val == rpio.Low && !btn1.Pressed {
			btn1.Pressed = true
			btn2.Pressed = false
			btn3.Pressed = false
			count = 0
			log.Println("Button happy pressed!")
			go client.Post("above")
		}

		if val := btn2.Read(); val == rpio.Low && !btn2.Pressed {
			btn1.Pressed = false
			btn2.Pressed = true
			btn3.Pressed = false
			count = 0
			log.Println("Button medium pressed!")
			go client.Post("average")
		}

		if val := btn3.Read(); val == rpio.Low && !btn3.Pressed {
			btn1.Pressed = false
			btn2.Pressed = false
			btn3.Pressed = true
			count = 0
			log.Println("Button sad pressed!")
			go client.Post("below")
		}

		time.Sleep(time.Second / 5)

		if count == 3 {
			count = 0
			btn1.Pressed = false
			btn2.Pressed = false
			btn3.Pressed = false
		}

		count += 1
	}
}
