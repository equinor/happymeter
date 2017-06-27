# Raspberry Pi "happymeter" written in Go

Uses "go-rpio" (github.com/stianeikeland/go-rpio) to listen for Pin changes and 
react to those based on which button was pressed.



## Buttons

To keep things very simple we have:
1. Happy
2. Mediocre
3. Sad


## Actions?

For each button action we want to do a HTTP post to a remote web service
