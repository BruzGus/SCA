package configuration

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	//
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

//Configuration estructura para mapear la DATA
//del archivo de configuracion.
type Configuration struct {
	Server   string
	Port     string
	User     string
	Password string
	DataBase string
}

//GetConfiguration obtenemos la configuracion
//de la base de datos.
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

//GetConnection obtenemos la coneccion a la base de datos.
func GetConnection() *gorm.DB {
	c := GetConfiguration()
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", c.Server, c.Port, c.User,c.DataBase, c.Password)
	db, err := gorm.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
