pakage handlers

import (
  "encoding/json"
  "net/http"
  "strings"
)

func UserRouter(w http.ResponseWrinter, r *http.Request) {
  //path := strings.TrimSuffix(r.URL.Path, "/user/")
  
  switch r.Method {
	case http.MethodGet:
		return
	case http.MethodPost:
		return
	case http.MethodPut: //update
		return
	case http.MethodDelete:
		return
	default:
		postError(w, http.StatusMethodNotAllowed)
	}
}