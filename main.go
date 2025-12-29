package main

import (
	"Gator/internal/config"
	"fmt"
	"os"
)

func main() {
	cmd_list := os.Args

	if len(cmd_list) < 2 {
		fmt.Println("error. missing arguments")
		os.Exit(1)
	}

	cfg, err := config.Read()
	if err != nil {
		fmt.Println("Error on Read")
	}

	newState := state{}
	newState.cfg = &cfg
	cmds := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}

	cmds.register("login", handlerLogin)
	cmd_inst := command{
		name:      cmd_list[1],
		arguments: cmd_list[2:],
	}
	err = cmds.run(&newState, cmd_inst)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
