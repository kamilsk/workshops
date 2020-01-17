package handler

//go:generate mockgen -package $GOPACKAGE -destination mocks_test.go tech.avito.ru/service/internal/handler Caller

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/golang/mock/gomock"
	_ "github.com/golang/mock/mockgen/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"tech.avito.ru/service/internal"
)

func TestProxy(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	action := internal.Action{URI: "https://play.golang.org", Payload: nil}
	caller := NewMockCaller(ctrl)
	record := httptest.NewRecorder()

	caller.EXPECT().Call(action).Return(nil)

	request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/?uri=%s", url.QueryEscape(action.URI)), nil)
	require.NoError(t, err)

	Proxy(caller).ServeHTTP(record, request)
	assert.Equal(t, http.StatusOK, record.Code)
}
