package responders

import (
	"encoding/json"
	"net/http"
)

func JSON(obj interface{}, statusCode int) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(statusCode)
		enc := json.NewEncoder(rw)
		enc.SetIndent("", "  ")
		enc.Encode(obj)
	})
}
