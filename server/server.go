package server

import (
	"github.com/lzientek/octopush-middleware/config"
)

func Init() {
	c := config.GetConfig()
	r := NewRouter()

	r.Run(c.GetString("app.port"))
}
