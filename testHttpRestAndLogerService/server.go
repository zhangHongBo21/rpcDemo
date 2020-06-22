/*
@Time : 2020/6/20 10:20
@Author : zhb
@File : server
@Software: GoLand
*/
package main

import "errors"

type Server interface {
	Add(a, b int) int
	Sub(a, b int) int
	Mul(a, b int) int
	Div(a, b int) (int, error)
}
type ArithmeticService struct {
}

func (t ArithmeticService) Add(a, b int) int {
	return a + b
}
func (t ArithmeticService) Sub(a, b int) int {
	return a - b
}
func (t ArithmeticService) Mul(a, b int) int {
	return a * b
}
func (t ArithmeticService) Div(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("the dividend can not be zero!")
	}

	return a / b, nil
}

type ServerMiddleware func(Server) Server
