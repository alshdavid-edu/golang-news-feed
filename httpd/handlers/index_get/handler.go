package handlersIndexGet

import (
	"encoding/json"
	"net/http"
)

func Handler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		bytes, _ := json.Marshal(&Response{Message: "Hello!"})
		w.WriteHeader(200)
		w.Write(bytes)
	}
}
