package main

import (
	"net/http"
)

func main() {
	http.ListenAndServe(":5000", ChiRouter().InitRouter())
}