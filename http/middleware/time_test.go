package middleware

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)


func TestTimeMiddleware(t *testing.T) {
	handler := TimeMiddleware(http.HandlerFunc(Echo))
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v1/api/redis/list/len?name=myhyh", nil)
	handler.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "echo", w.Body.String())
}