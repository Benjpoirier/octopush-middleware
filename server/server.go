package server

import (
	"github.com/lzientek/octopush-middleware/config"
	"github.com/lzientek/octopush-middleware/db"
)

func Init() {
	db.Create()

	c := config.GetConfig()
	r := NewRouter()

	r.Run(c.GetString("app.port"))
}
