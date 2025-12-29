package main

import (
	"Gator/internal/config"
	"fmt"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println("Error on Read")
	}
	err = cfg.SetUser("Encelado")
	if err != nil {
		fmt.Println("Error on SetUser")
	}

	cfg, err = config.Read()
	if err != nil {
		fmt.Println("Error on Read, 2°")
	}
	fmt.Println(cfg)
}
