package main

import (
	"net/http"
	"time"
)

func main() {

	go DoEvery(10*time.Second)

	http.ListenAndServe(":5000", ChiRouter().InitRouter())
}