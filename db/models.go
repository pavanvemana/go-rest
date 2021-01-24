package db

import (
	"log"
	"crypto/sha256"
	"fmt"
)


type User struct{
	Username string `json:"username"`
	Password string `json:"password"`
	Email string `json:"email"`
	Created_on string `json:"created_at"`
}


func (u *User)  Save(){
	db_conn := Connect("gotest")
	password := fmt.Sprintf("%x", sha256.Sum256([]byte(u.Password)))
	inserted := Exec(db_conn, 
		"INSERT into users (username, password, email, created_on) values ($1, $2, $3, $4)", 
		u.Username, password, u.Email, u.Created_on)
	if (inserted){
		log.Print("Recorded inserted successfully")
	}
}