package main

import (
	"fmt"

	"github.com/gorilla/mux.git"
)

func main() {
	r := mux.NewRouter()
	fmt.Printf("%+v", r)
}
