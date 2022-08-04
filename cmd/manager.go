package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/JannikStr/dailynote/pkg/config"
)

type CommandManager struct {
	Reader        *bufio.Reader
	Configuration config.Config
	Commands      map[string]Command
	ShouldClose   bool
}

type Command struct {
	Name    string
	Help    string
	Handler CommandHandler
}

type CommandHandler interface {
	Execute()
}

var DailyNoteManager CommandManager

func (manager *CommandManager) Init(cfg *config.Config) {
	DailyNoteManager.Reader = bufio.NewReader(os.Stdin)
	DailyNoteManager.Configuration = *cfg
	DailyNoteManager.Commands = make(map[string]Command)
	DailyNoteManager.ShouldClose = false
}

func (manager *CommandManager) Run() {
	for !DailyNoteManager.ShouldClose {
		fmt.Print(" > ")
		input, err := DailyNoteManager.Reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}

		input = strings.TrimSpace(input)

		enteredCommand, ok := DailyNoteManager.Commands[input]

		if !ok {
			fmt.Println("Unknown command. Use 'help' to see a list of commands")
			continue
		}

		fmt.Println(enteredCommand.Help)
	}
}

func (manager *CommandManager) RegisterCommand(cmd Command) {
	DailyNoteManager.Commands[cmd.Name] = cmd
}
