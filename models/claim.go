package models

import jwt "github.com/dgrijalva/jwt-go"

//Claim , estructura para el manejo 
//de las peticiones de los Token.
type Claim struct {
	User `json:"user"`
	jwt.StandardClaims
}
