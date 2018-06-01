package main

import (
	"flag"
)

var (
	fDiscord bool
	fIRC     bool
	dToken   string
	iToken   string
)

func init() {
	flag.BoolVar(&fDiscord, "de", false, "Turn on the Discord Bot")
	flag.StringVar(&dToken, "t", "", "Discord Token")
	flag.BoolVar(&fIRC, "ie", false, "Enable IRC bot")
	flag.StringVar(&iToken, "it", "", "Twitch OAuth token")
	flag.Parse()
}

func main() {
	if fDiscord {
		RunDiscord()
	}
	if fIRC {

	}
}
