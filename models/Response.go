//Package models
/*
   @author:xie
   @date:2022/6/1
   @note:
*/
package models

type Response struct {
	Success bool   `json:"success"`
	Data    any    `json:"data"`
	Msg     string `json:"msg"`
}

func NewResponse(success bool, data any, msg string) Response {
	return Response{
		Success: success,
		Data:    data,
		Msg:     msg,
	}
}
func (r *Response) SetSuccess(success bool) *Response {
	r.Success = success
	return r
}
func (r *Response) SetData(data any) *Response {
	r.Data = data
	return r
}

func (r *Response) SetMsg(msg string) *Response {
	r.Msg = msg
	return r
}
