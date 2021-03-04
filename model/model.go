package model

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
