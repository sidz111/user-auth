package main

import (
	"fmt"
	"log"

	"github.com/sidz111/user-auth/config"
)

func main() {
	if err := config.ConnectDB(); err != nil {
		log.Fatal("Failed")
	} else {
		fmt.Println("Connected ")
	}
}
