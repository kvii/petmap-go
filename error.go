package main

import "errors"

// 错误
var (
	ErrUserNotFound = errors.New("用户名或密码错误")
	ErrBadRequest   = errors.New("请求数据格式错误")
)
