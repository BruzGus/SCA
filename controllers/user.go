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
	user := models.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Fprintf(w, "Error: %s\n", err)
		return
	}
	db := configuration.GetConnection()
	defer db.Close()

	c: sha256.Sum256([]byte(user.Pass))
	pwd := base64.URLEncoding.EncodeToString(c[:32])

	db.Where("email =? and pin=?", user.Email , pwd).First(&user)

	

}
