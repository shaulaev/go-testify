package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMainHandlerOK(t *testing.T) {
    req := httptest.NewRequest("GET", "/cafe?count=2&city=moscow", nil)
    rr := httptest.NewRecorder()

    handler := http.HandlerFunc(mainHandle)
    handler.ServeHTTP(rr, req)

    require.Equal(t, http.StatusOK, rr.Code)
    assert.NotEmpty(t, rr.Body.String())
}

func TestMainHandlerWrongCity(t *testing.T) {
    req := httptest.NewRequest("GET", "/cafe?count=2&city=spb", nil)
    rr := httptest.NewRecorder()

    handler := http.HandlerFunc(mainHandle)
    handler.ServeHTTP(rr, req)

    require.Equal(t, http.StatusBadRequest, rr.Code)
    assert.Equal(t, "wrong city value", rr.Body.String())
}

func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
    totalCount := 4

    req := httptest.NewRequest("GET", "/cafe?count=100&city=moscow", nil)
    rr := httptest.NewRecorder()

    handler := http.HandlerFunc(mainHandle)
    handler.ServeHTTP(rr, req)

    require.Equal(t, http.StatusOK, rr.Code)

    result := rr.Body.String()

    cafes := strings.Split(result, ",")

    assert.Len(t, cafes, totalCount)
}
