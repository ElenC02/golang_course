package main

import (
	"net/http"
	
	"github.com/Elen0207/golang_course/terceiro_modulo/src/routes"
)

func main() {
	routes.CarregaRotas()
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err.Error())
	}
}