package main

import (
	"github.com/JannikStr/dailynote/pkg/config"
)

func main() {
	cfg, exists := config.LoadConfig()

	if exists {
		config.CreateConfigFolder(cfg)
	}

}
