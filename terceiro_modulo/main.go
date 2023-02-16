package main

import (
	"net/http"

	"github.com/Elen0207/golang_course/terceiro_modulo/src/routes"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}