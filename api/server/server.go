package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/isneru/learning-go/api/router"
)

func ServeApi() {
	http.HandleFunc("/users", router.GetUsers)
	fmt.Println("Server is running on port 8080")
	
	log.Fatal(http.ListenAndServe(":8080", nil))
}