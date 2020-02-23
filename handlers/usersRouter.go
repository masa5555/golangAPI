pakage handlers

import (
  "encoding/json"
  "net/http"
  "strings"
)

func UserRouter(w http.ResponseWrinter, r *http.Request) {
  path := strings.TrimSuffix(r.URL.Path, "/user/")
  if !json.IsObjectIdHex(Path) {
    postError(wm http.StatusNotFound)
    return
  }
  
  switch r.Method {
	case http.MethodGet:
		return
	case http.MethodCreate:
		return
	case http.MethodUpdate:
		return
	case http.MethodDelete:
		return
	default:
		postError(w, http.StatusMethodNotAllowed)
	}
}