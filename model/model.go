package model

import (
	"encoding/json"
	"fmt"
	"net/http/httptest"
)

type ResponseMessage struct {
	Message string `json:"message" example:"success" format:"string"`
	Size    int    `json:"size" example:"4"`
}

type ListLenOutQueueStruct struct {
	Name string `json:"name" example:"list name"`
}

type ListInQueueStruct struct {
	Name  string `json:"name" example:"list name"`
	Items []int  `json:"element" example:"1,2,3"`
}

type UserInfo struct {
	UserName string `form:"username" json:"user_name" example:"xiaoming"`
	Password string `form:"password" json:"password" example:"123"`
}

type FmtResponse struct {
	Code int `form:"code" json:"code" example:"200"`
	Msg string `form:"msg" json:"msg" example:"success"`
	Data interface{} `form:"data" json:"data" example:"{}"`
}

func NewFmtResponse() *FmtResponse{
	return &FmtResponse{
		Code: 200,
		Msg: "success",
	}
}

func NewFmtResponseByResponseRecorder(w *httptest.ResponseRecorder) (*FmtResponse, error){
	// 在http响应中使用
	v := NewFmtResponse()
	err := json.Unmarshal(w.Body.Bytes(), v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func (resp *FmtResponse) String() string {
	return fmt.Sprintf("code:%v msg:%v data:%v", resp.Code, resp.Msg, resp.Data)
}

func (resp *FmtResponse) Equal(v *FmtResponse) (v1, v2 string, equal bool) {
	v1 = fmt.Sprintf("code:%v msg:%v data:%v", resp.Code, resp.Msg, resp.Data)
	v2 = fmt.Sprintf("code:%v msg:%v data:%v", v.Code, v.Msg, v.Data)
	return v1, v2, v1 == v2
}

func (resp *FmtResponse) EqualResponseRecorder(w *httptest.ResponseRecorder) (v1, v2 string, equal bool, err error) {
	var v *FmtResponse
	v, err = NewFmtResponseByResponseRecorder(w)
	if err != nil {
		return "", "", false, err
	}
	v1 = resp.String()
	v2 = v.String()
	return v1, v2, v1 == v2, nil
}