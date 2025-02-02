package main

import (
	"fmt"
	"log"

	"github.com/wolv89/bootgator/internal/config"
)

func main() {

	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	cfg.SetUser("pete")

	cfg2, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(cfg2)

}
