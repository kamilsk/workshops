package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"tech.avito.ru/service/internal"
)

type Caller interface {
	Call(context.Context, internal.Action) error
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

		ctx, cancel := context.WithTimeout(req.Context(), Header(req.Header).Timeout(time.Second, 1.0))
		defer cancel()

		err := caller.Call(ctx, internal.Action{
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
