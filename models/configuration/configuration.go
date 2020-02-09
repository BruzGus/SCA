package configuration

import (
	"encoding/json"
	"log"
	"os"
	 	 
)
//Configuration ...., estructura para mapear la DATA del archivo de configuracion.
type Configuration struct {
	Server   string
	Port     string
	User     string
	Password string
	DataBase string
}
//GetConfiguration ..., obtenemos la configuracion de la base de datos.
func GetConfiguration() Configuration {
	var c Configuration

	file, err := os.Open("./config.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&c)
	if err != nil {
		log.Fatal(err)
	}
	return c
}
