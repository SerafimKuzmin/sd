package main

import (
	"flag"
	"github.com/SerafimKuzmin/sd/src/cmd/mini_imdb"
	"log"

	"github.com/BurntSushi/toml"
)

var (
	configPath string
	sessionDB  string
)

func init() {
	flag.StringVar(&configPath, "config-path", "./config.toml", "path to config file")
	flag.StringVar(&sessionDB, "sesseion-db", "redis", "what db the application uses for session")
}

func main() {
	flag.Parse()

	timeTracker := mini_imdb.MiniIMDB{}
	_, err := toml.DecodeFile(configPath, &timeTracker)

	if err != nil {
		log.Fatal(err)
	}

	err = timeTracker.Run(sessionDB)

	if err != nil {
		log.Fatal(err)
		return
	}
}
