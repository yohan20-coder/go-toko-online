package controllers

import (
	"fmt"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to Aplikasi Golang Andi")
}

func Test(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}
