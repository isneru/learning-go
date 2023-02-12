package httputils

import "net/http"

func AllowsMethod(w http.ResponseWriter, r *http.Request, method string) bool{
	if r.Method != method {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return false
	}
	return true
}