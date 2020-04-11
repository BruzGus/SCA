package controllers

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/BruzGus/SCA/configuration"
	"github.com/BruzGus/SCA/models"
)

// Login controla toda la logica para el usuario
func Login(w http.ResponseWriter, r *http.Request) {
	/*user := models.User{}*/
	tecnico := models.Tecnico{}
	err := json.NewDecoder(r.Body).Decode(&tecnico)
	if err != nil {
		fmt.Fprintf(w, "Error: %s\n", err)
		return
	}

	db := configuration.GetConnection()
	defer db.Close()

	c := sha256.Sum256([]byte(tecnico.Pin))
	pwd := base64.URLEncoding.EncodeToString(c[:32])

	db.Where("email =? and pin=?", tecnico.Email, pwd).First(&tecnico)

	if tecnico.ID != ""{
		
	}

}
