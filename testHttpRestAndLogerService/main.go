/*
@Time : 2020/6/20 11:27
@Author : zhb
@File : main
@Software: GoLand
*/
package main

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	ctx := context.Background()
	errChan := make(chan error)

	var svc Server
	svc = ArithmeticService{}
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.With(logger, "ts", log.DefaultTimestampUTC)
		logger = log.With(logger, "caller", log.DefaultCaller)
	}
	// 增加logger的中间件
	svc = LoggingMiddleware(logger)(svc)
	endpoint := MakeArithmeticEndpoint(svc)

	r := MakeHttpHandler(ctx, endpoint, logger)

	go func() {
		fmt.Println("Http Server start at port:9000")
		errChan <- http.ListenAndServe(":9000", r)
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-c)
	}()

	fmt.Println(<-errChan)
}
