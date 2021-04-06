package main

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"goLearning/model"
	"goLearning/server"
	"net/http"
	"net/http/httptest"
	"testing"
)

var r *gin.Engine

func init() {
	r = gin.Default()
	server.SetupRedisRouter(r)
}
func EqualResp(t *testing.T, expected *model.FmtResponse, w *httptest.ResponseRecorder)  {
	assert.Equal(t, 200, w.Code)
	v1, v2, equal, err := expected.EqualResponseRecorder(w)
	if err != nil {
		t.Errorf("err:%v\n", err)
		t.Fail()
	}
	if !equal{
		t.Logf("\nv1:%v\nv2:%v\n", v1, v2)
		t.Fail()
	}
}

func TestLenRoute(t *testing.T) {
	w := httptest.NewRecorder()
	var lenStruct model.ListLenOutQueueStruct
	lenStruct.Name = "abd"
	data, err := json.Marshal(lenStruct)
	if err != nil {
		t.Logf("err:%v\n", err)
		t.Fail()
	}
	req, _ := http.NewRequest("POST", "/v1/api/redis/list/len", bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)
	fmtResp := model.NewFmtResponse()
	fmtResp.Data = 0
	EqualResp(t, fmtResp, w)
}

func TestInQueueRoute(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/v1/api/redis/list/inQueue?element=1&element=2&element=3&name=myhyh", nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, "{\"inData\":[\"1\",\"2\",\"3\"],\"message\":\"success\"}", w)
}

func TestOutQueueRoute(t *testing.T) {
	w := httptest.NewRecorder()
	var inQueueStruct model.ListInQueueStruct
	inQueueStruct.Name = "abd"
	inQueueStruct.Items = []int{1, 2, 3}
	data, err := json.Marshal(inQueueStruct)
	if err != nil {
		t.Logf("err:%v\n", err)
		t.Fail()
	}
	req, _ := http.NewRequest("POST", "/v1/api/redis/list/len", bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)
	assert.Equal(t, "{\"message\":\"success\",\"outData\":[\"1\",\"2\",\"3\"]}", w)
}
