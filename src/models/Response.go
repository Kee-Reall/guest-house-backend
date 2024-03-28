package models

type Response struct {
	Description string      `json:"description"`
	Status      int         `json:"statusCode"`
	Data        interface{} `json:"data"`
}

type ResponseManager struct {
}

func CreateResponseManager() *ResponseManager {
	return &ResponseManager{}
}

func (r *ResponseManager) NewResponse(data interface{}) *Response {
	return &Response{Data: data, Description: "blaBla", Status: 200}
}
