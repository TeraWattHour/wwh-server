package main

import (
	"fmt"
	"log"

	"github.com/terawatthour/we-were-here-server/pkg"
	"github.com/terawatthour/we-were-here-server/pkg/data"
	"github.com/terawatthour/we-were-here-server/schema"
)

func init() {
	if err := pkg.LoadEnvironment(".env"); err != nil {
		log.Fatal("Error loading environment variables... ", err)
	}

	if err := data.EstablishPostgresConnection(); err != nil {
		log.Fatal("Error establishing postgres connection... ", err)
	}
}

func main() {

	user := schema.User{}
	err := data.PostgresClient.QueryRow("SELECT * from user").Scan(&user)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(user)
}
