package main

import (
	"github.com/ChimeraCoder/anaconda"
)

const (
	consumerKey  = ""
	consumerSecret = ""
	accessToken = ""
	accessTokenSecret = ""
)

func main() {
	anaconda.SetConsumerKey(consumerKey)
	anaconda.SetConsumerSecret(consumerSecret)
	anaconda.NewTwitterApi(accessToken, accessTokenSecret)

	
}