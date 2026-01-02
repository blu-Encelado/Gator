package main

import (
	"Gator/internal/config"
	"Gator/internal/database"
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
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
	dbURL := cfg.DbURL

	db, err := sql.Open("postgres", dbURL)
	dbQueries := database.New(db)
	newState.db = dbQueries

	cmds := commands{
		registeredCommands: make(map[string]func(*state, command) error),
	}

	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)

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
