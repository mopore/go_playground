package main

import (
	"fmt"
	"log"
	"net/http"
	"tut/apiwithdocker/middleware"
	"tut/apiwithdocker/user"
)

const (
	port = ":8080"
	apiV2Path = "/api/v2"
)

func defaultRoute(w http.ResponseWriter, r *http.Request) {
	testValue, ok := r.Context().Value("test-key").(string)
	if !ok {
		testValue = "not found"
	}

	response := fmt.Sprintf("Welcome to the example at \"%s\"\n" + 
		"Our example key value is \"%s\"", r.URL.Path, testValue)
	w.Write([]byte(response))
}

func main() {

	apiV2Router := http.NewServeMux()
	user.SetUserRoutes(apiV2Router)

	// Adds "/api/v2" as prefix
	main := http.NewServeMux()
	main.Handle(apiV2Path + "/", http.StripPrefix(apiV2Path, apiV2Router))

	// Add a default route for other than "/api/v2"
	main.HandleFunc("/", defaultRoute)

	stack := middleware.CreateStack(
		middleware.ValueInjector,
		middleware.Logging,
		// ...
	)

	server := http.Server{
		Addr:    port,
		Handler: stack(main),
	}

	log.Printf("Server started at %s\n", port)
	server.ListenAndServe()
}
