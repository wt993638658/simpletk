package main

import (
	relation "github.com/wt993638658/simpletk/kitex_gen/relation/relationsrv"
	"log"
)

func main() {
	svr := relation.NewServer(new(RelationSrvImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
