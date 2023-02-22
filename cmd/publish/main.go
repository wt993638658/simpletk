package main

import (
	publish "github.com/wt993638658/simpletk/kitex_gen/publish/publishsrv"
	"log"
)

func main() {
	svr := publish.NewServer(new(PublishSrvImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
