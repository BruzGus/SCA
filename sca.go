package main

import (
	"fmt"

	"github.com/BruzGus/SCA/configuration"
)

func main() {
	b:= configuration.GetConnection()

	defer b.Close()

	database:= b.DB()

	err:= database.Ping()

	if err != nil {
		panic(err.Error())
	}

	fmt.Println("Connection to PostgresSQL was successfull!")
	
}