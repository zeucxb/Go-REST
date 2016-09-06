package main

import (
	"net/http"
	"server/routes"
)

func main() {
	routes.R.ServeFiles("/avatar/*filepath", http.Dir("./avatar"))

	// Fire up the server
	http.ListenAndServe("localhost:3000", routes.R)
}
