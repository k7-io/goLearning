package main

import (
	"go_learning/server"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLenRoute(t *testing.T) {
	router := server.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/redis/list/len?name=myhyh", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "{\"message\":\"success\",\"size\":0}", w.Body.String())
}

func TestInQueueRoute(t *testing.T) {
	router := server.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/redis/list/inQueue?element=1&element=2&element=3&name=myhyh", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "{\"inData\":[\"1\",\"2\",\"3\"],\"message\":\"success\"}", w.Body.String())
}

func TestOutQueueRoute(t *testing.T) {
	router := server.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/redis/list/outQueue?name=myhyh", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "{\"message\":\"success\",\"outData\":[\"1\",\"2\",\"3\"]}", w.Body.String())
}
