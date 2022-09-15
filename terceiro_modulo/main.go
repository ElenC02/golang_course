package main

import (
	"net/http"

	"golang_course/terceiro_modulo/routes"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}