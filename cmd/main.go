package cmd

import (
	"apiserver"
	"log"
)

func main() {
	srv := new(apiserver.Server)
	if err := srv.Run("8080"); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}

}