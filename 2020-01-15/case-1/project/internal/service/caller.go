package service

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"tech.avito.ru/service/internal"
)

func New() *implementation {
	return &implementation{}
}

type implementation struct{}

func (i *implementation) Call(action internal.Action) error {
	if action.Payload == nil {
		resp, err := http.Get(action.URI)
		if err == nil {
			discard(resp.Body)
		}
		return err
	}
	body := new(bytes.Buffer)
	if err := json.NewEncoder(body).Encode(action.Payload); err != nil {
		return err
	}
	resp, err := http.Post(action.URI, "application/json", body)
	if err == nil {
		discard(resp.Body)
	}
	return err
}

func discard(body io.ReadCloser) {
	_, _ = io.Copy(ioutil.Discard, body)
	_ = body.Close()
}
