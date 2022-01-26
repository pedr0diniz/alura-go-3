package main

// the _ import means that package will only be used in runtime, not in compilation
import (
	"net/http"
	"store/routes"

	_ "github.com/lib/pq"
)

// default main function
func main() {
	routes.LoadRoutes()
	http.ListenAndServe(":8000", nil)
}
