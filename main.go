package main

import (
	"net/http"
)

func main() {

	port := "5432"
	m := MakeHandler()
	err := http.ListenAndServe(":"+port, m)
	if err != nil {
		panic(err)
	}
}
