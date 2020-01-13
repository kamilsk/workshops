package service

import (
	"bytes"
	"context"
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

func (i *implementation) Call(ctx context.Context, action internal.Action) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, action.URI, nil)
	if err != nil {
		return err
	}
	if action.Payload != nil {
		body := new(bytes.Buffer)
		if err := json.NewEncoder(body).Encode(action.Payload); err != nil {
			return err
		}
		if req, err = http.NewRequestWithContext(ctx, http.MethodPost, action.URI, body); err != nil {
			return err
		}
	}
	resp, err := http.DefaultClient.Do(req)
	if err == nil {
		discard(resp.Body)
	}
	return err
}

func discard(body io.ReadCloser) {
	io.Copy(ioutil.Discard, body)
	body.Close()
}
