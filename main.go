package main

import (
	"github.com/JannikStr/dailynote/cmd"
	"github.com/JannikStr/dailynote/pkg/config"
)

func main() {
	cfg, exists := config.LoadConfig()

	if exists {
		config.CreateConfigFolder(cfg)
	}
	exitCmd := cmd.Command{
		Name:    "exit",
		Help:    "exit -> will exit DailyNoteManager and return to shell",
		Handler: nil,
	}

	cmd.DailyNoteManager.Init(&cfg)

	cmd.DailyNoteManager.RegisterCommand(exitCmd)

	cmd.DailyNoteManager.Run()

}
