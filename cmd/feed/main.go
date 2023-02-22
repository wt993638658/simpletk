package main

import (
	feed "github.com/wt993638658/simpletk/kitex_gen/feed/feedsrv"
	"log"
)

func main() {
	svr := feed.NewServer(new(FeedSrvImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}