package main

import (
	"fmt"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	fmt.Printf("%+v", r)
}
