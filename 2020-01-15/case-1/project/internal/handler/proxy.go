package handler

import (
	"encoding/json"
	"net/http"

	"tech.avito.ru/service/internal"
)

type Caller interface {
	Call(internal.Action) error
}

func Proxy(caller Caller) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		if err := req.ParseForm(); err != nil {
			http.Error(rw, "please provide a query and data", http.StatusBadRequest)
			return
		}
		uri := req.Form.Get("uri")
		var payload interface{}
		if req.Method != http.MethodGet {
			if err := json.NewDecoder(req.Body).Decode(&payload); err != nil || uri == "" {
				http.Error(rw, "invalid uri and/or request body", http.StatusBadRequest)
				return
			}
		}

		err := caller.Call(internal.Action{
			URI:     uri,
			Payload: payload,
		})

		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}
		rw.WriteHeader(http.StatusOK)
	})
}
