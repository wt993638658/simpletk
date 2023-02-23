package main

import (
	comment "github.com/wt993638658/simpletk/kitex_gen/comment/commentsrv"
	"log"
)

func main() {
	svr := comment.NewServer(new(CommentSrvImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
