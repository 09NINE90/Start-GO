package main

import (
	"github.com/go-chi/jwtauth/v5"
	"log"
)

var tokenAuth *jwtauth.JWTAuth

func init() {
	tokenAuth = jwtauth.New("HS256", []byte("secret"), nil)

	// For debugging/example purposes, we generate and print
	// a sample jwt token with claims `user_id:123` here:
	_, tokenString, _ := tokenAuth.Encode(map[string]interface{}{"user_id": 123})
	log.Printf("DEBUG: a sample jwt is %s\n\n", tokenString)
}
