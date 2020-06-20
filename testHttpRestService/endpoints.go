/*
@Time : 2020/6/20 10:26
@Author : zhb
@File : endpoints
@Software: GoLand
*/
package main

import (
	"context"
	"errors"
	"github.com/go-kit/kit/endpoint"
)

var (
	ErrInvalidRequestType = errors.New("RequestType has only four type: Add,Subtract,Multiply,Divide")
)

type ArithmeticRequst struct {
	RequestType string `json:"requst_type"`
	A           int    `json:"a"`
	B           int    `json:"b"`
}
type ArithmeticResponse struct {
	Result int   `json:"result"`
	Error  error `json:"error"`
}

type ArithmeticEndpoint endpoint.Endpoint

func MakeArithmeticEndpoint(s Server) endpoint.Endpoint {

	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(ArithmeticRequst)
		var (
			result, a, b  int
			callbackerErr error
		)
		a = req.A
		b = req.B
		switch req.RequestType {
		case "add":
			result = s.Add(a, b)
		case "sub":
			result = s.Sub(a, b)
		case "mul":
			result = s.Mul(a, b)
		case "div":
			result, callbackerErr = s.Div(a, b)
		default:
			result = s.Add(a, b)
		}
		return result, callbackerErr
	}
}
