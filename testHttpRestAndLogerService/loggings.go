/*
@Time : 2020/6/22 14:04
@Author : zhb
@File : loggings
@Software: GoLand
*/
package main

import (
	"github.com/go-kit/kit/log"
	"time"
)

type loggingMiddleware struct {
	Server
	logger log.Logger
}

//日志功能的中间件结构体
func LoggingMiddleware(logger log.Logger) ServerMiddleware {
	return func(next Server) Server {
		return loggingMiddleware{next, logger}
	}
}

//重新实现add方法
func (mw loggingMiddleware) Add(a, b int) (ret int) {

	defer func(beign time.Time) {
		mw.logger.Log(
			"function", "Add",
			"a", a,
			"b", b,
			"result", ret,
			"took", time.Since(beign),
		)
	}(time.Now())

	ret = mw.Server.Add(a, b)
	return ret
}

func (mw loggingMiddleware) Sub(a, b int) (ret int) {

	defer func(beign time.Time) {
		mw.logger.Log(
			"function", "Sub",
			"a", a,
			"b", b,
			"result", ret,
			"took", time.Since(beign),
		)
	}(time.Now())

	ret = mw.Server.Sub(a, b)
	return ret
}

func (mw loggingMiddleware) Mul(a, b int) (ret int) {

	defer func(beign time.Time) {
		mw.logger.Log(
			"function", "Mul",
			"a", a,
			"b", b,
			"result", ret,
			"took", time.Since(beign),
		)
	}(time.Now())

	ret = mw.Server.Mul(a, b)
	return ret
}

func (mw loggingMiddleware) Div(a, b int) (ret int, err error) {

	defer func(beign time.Time) {
		mw.logger.Log(
			"function", "Div",
			"a", a,
			"b", b,
			"result", ret,
			"took", time.Since(beign),
		)
	}(time.Now())

	ret, _ = mw.Server.Div(a, b)
	return ret, nil
}
