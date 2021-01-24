package main

import (
	"net/http"
	"log"
	_ "github.com/pavanvemana/go-rest/urls"
)

func main()  {
	
	log.Fatal(http.ListenAndServe(":8080", nil))
}