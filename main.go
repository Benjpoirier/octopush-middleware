package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/lzientek/octopush-middleware/config"
	"github.com/lzientek/octopush-middleware/db"
	"github.com/lzientek/octopush-middleware/server"
)

func main() {
	enviroment := flag.String("e", "development", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	config.Init(*enviroment)
	db.Init()
	server.Init()
	defer db.CloseDB()
}
