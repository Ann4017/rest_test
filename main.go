package main

import (
	"net/http"
	"rest_test/myapp"
)

func main() {
	http.ListenAndServe(":3000", myapp.New_http_handler())
}
