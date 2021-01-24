package urls

import (
	"net/http"
	"github.com/pavanvemana/go-rest/views"
	"log"
)

func init()  {
	log.Print("Registering urls")
	http.HandleFunc("/users", (*views.UserView)(nil).Get)
	http.HandleFunc("/user/create", (*views.UserView)(nil).Post)
	http.HandleFunc("/user/delete", (*views.UserView)(nil).Delete)
}