/*
@Time : 2020/6/20 11:06
@Author : zhb
@File : transports
@Software: GoLand
*/
package main

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

var (
	ErrorBadRequest = errors.New("invalid request parameter")
)

func MakeHttpHandler(ctx context.Context, endpoint endpoint.Endpoint, logger log.Logger) http.Handler {
	r := mux.NewRouter()
	options := []kithttp.ServerOption{
		kithttp.ServerErrorLogger(logger),
		kithttp.ServerErrorEncoder(kithttp.DefaultErrorEncoder),
	}

	r.Methods("POST").Path("/calculate/{type}/{a}/{b}").Handler(kithttp.NewServer(
		endpoint,
		decodeArithmeticRequest,
		encodeArithmeticResponse,
		options...,
	))

	//r.Path("/metrics").Handler(promhttp.Handler())

	return r
}

//解码
func decodeArithmeticRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	requestType, ok := vars["type"]
	if !ok {
		return nil, ErrorBadRequest
	}

	pa, ok := vars["a"]
	if !ok {
		return nil, ErrorBadRequest
	}

	pb, ok := vars["b"]
	if !ok {
		return nil, ErrorBadRequest
	}

	a, _ := strconv.Atoi(pa)
	b, _ := strconv.Atoi(pb)

	return ArithmeticRequst{
		RequestType: requestType,
		A:           a,
		B:           b,
	}, nil
}

// 转码
func encodeArithmeticResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}
