package main

import (
	"flag"
	"fmt"
	"zero-chat/app/chat/cmd/mq/internal/config"
	"zero-chat/app/chat/cmd/mq/internal/listen"

	"github.com/joho/godotenv"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
)

var configFile = flag.String("f", "etc/chat.yaml", "Specify the config file")

func main() {
	flag.Parse()
	var c config.Config
	godotenv.Load()
	conf.MustLoad(*configFile, &c)

	// log、prometheus、trace、metricsUrl.
	if err := c.SetUp(); err != nil {
		panic(err)
	}

	serviceGroup := service.NewServiceGroup()
	defer serviceGroup.Stop()

	for _, mq := range listen.Mqs(c) {
		serviceGroup.Add(mq)
	}

	serviceGroup.Start()
	fmt.Printf("Starting mq server")

}
