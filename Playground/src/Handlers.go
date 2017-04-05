package main

import (
	"net/http"
	"fmt"
)

func Index(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "Welcome")
}

func Login(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "Login")
}

