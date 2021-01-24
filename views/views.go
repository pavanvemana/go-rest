package views

import (
	"net/http"
	"github.com/pavanvemana/go-rest/db"
	"encoding/json"
	"io"
	"log"
)

type HttpRequest interface{
	Get(w http.ResponseWriter, r *http.Request)
	Post(w http.ResponseWriter, r *http.Request)
}

type UserView struct{
	
}

func (c *UserView) Get(w http.ResponseWriter, r *http.Request){
	db_conn := db.Connect("gotest")
	rows := db.Query(db_conn, "SELECT user_id, username, created_on from users")
	var op []db.User
	for rows.Next() {
		var user db.User
		err := rows.Scan(&user.Username, &user.Email, &user.Created_on)
		if (err != nil){
			panic(err)
		}
		op = append(op, user)
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	marshalled, _ := json.Marshal(op)
	io.WriteString(w, string(marshalled))
}


func (c *UserView) Post(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	username := r.PostFormValue("username")
	email := r.PostFormValue("email")
	password := r.PostFormValue("password")
	created_on := r.PostFormValue("created_on")

	user := &db.User{
		Username: username, 
		Password: password,
		Email: email,
		Created_on: created_on,
	}
	user.Save()
}

func (c *UserView) Delete(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	user_id := r.PostFormValue("user_id")
	db_conn := db.Connect("gotest")
	deleted := db.Exec(db_conn,
		"DELETE FROM users where user_id=$1", user_id)
	if(deleted){
		log.Printf("User - %s deleted successfully", user_id)
	}
	w.Header().Set("Status-Code", "200")
	io.WriteString(w, "Success")
}
