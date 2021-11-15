package main

import (
	thrifttest "github.com/cloudwego/kitex/output/kitex_gen/thrifttest/combineservice"
	"log"
)

func main() {
	svr := thrifttest.NewServer(new(CombineServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
