package main

import (
	"fmt"
)

// 用户
type User struct {
	UserName string `json:"userName"` // 用户名
	Password string `json:"password"` // 密码
}

// 宠物
type Pet struct {
	PetName   string  `json:"petName"`   // 宠物名
	Owner     string  `json:"owner"`     // 拥有者
	Latitude  float64 `json:"latitude"`  // 纬度
	Longitude float64 `json:"longitude"` // 经度
}

// 家
type Home struct {
	Owner     string  `json:"owner"`     // 拥有者
	Latitude  float64 `json:"latitude"`  // 纬度
	Longitude float64 `json:"longitude"` // 经度
}

// 消息
type Message struct {
	Sender   string `json:"sender"`   // 发送者
	Receiver string `json:"receiver"` // 接收者
	Content  string `json:"content"`  // 内容
}

// 数据表
type Tables struct {
	Users    []User    `json:"users"`    // 用户
	Pets     []Pet     `json:"pets"`     // 宠物
	Homes    []Home    `json:"homes"`    // 家
	Messages []Message `json:"messages"` // 消息
}

// 用户全量信息
type UserFullInfo struct {
	User User  `json:"user"`           // 用户
	Home Home  `json:"home,omitempty"` // 家
	Pets []Pet `json:"pets"`           // 宠物
}

// 创建宠物走丢协助消息
func CreateMessagePetLostRequest(receiver, owner, petName string) Message {
	return Message{
		Sender:   owner,
		Receiver: receiver,
		Content:  fmt.Sprintf("请帮我找找走丢的%s吧。", petName),
	}
}

// 创建宠物走丢提示消息
func CreateMessagePetLostHint(owner, petName string) Message {
	return Message{
		Sender:   "系统",
		Receiver: owner,
		Content:  fmt.Sprintf("您的宠物%s已走丢。", petName),
	}
}
