package router

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/isneru/learning-go/api/server/httputils"
	"github.com/isneru/learning-go/types"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	if httputils.AllowsMethod(w, r, "GET") {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		resp, err := http.Get("https://jsonplaceholder.typicode.com/users")
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	
		if resp.StatusCode != http.StatusOK {
			fmt.Println("Error. Status Code: ", resp.Status)
			return
		}
	
		body, err := io.ReadAll(resp.Body)
		
		var response []types.User

		if err != nil {
			fmt.Println("Error retrieving data. ",err.Error())
			return
		}
		
		json.Unmarshal(body, &response)

		 json.NewEncoder(w).Encode(response)
	}
}