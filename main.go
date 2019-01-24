package main

import (
	"log"

	"github.com/JAG-UK/numerologySolver/app"
	"github.com/JAG-UK/numerologySolver/config"
)

func main() {
	config.InitConfig()
	conf := config.GetConfig()
	err := config.InitDB()
	if err != nil {
		log.Fatal("Fatal error encountered trying to open database connection: ", err.Error())
	}
	app.InitAndRun(conf)
}
