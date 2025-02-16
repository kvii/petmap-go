package main

import (
	"encoding/json"
	"log/slog"
	"os"
)

type Repository struct {
	Logger *slog.Logger
}

type GetUser struct {
	UserName string `json:"userName"` // 用户名
	Password string `json:"password"` // 密码
}

// 获取用户
func (r Repository) GetUser(p GetUser) (User, error) {
	for _, u := range tables.Users {
		if u.UserName == p.UserName && u.Password == p.Password {
			return u, nil
		}
	}
	r.Logger.Info("用户未找到",
		slog.String("user", p.UserName),
		slog.String("pass", p.Password),
	)
	return User{}, ErrUserNotFound
}

type GetMessages struct {
	UserName string
}

// 获取信息
func (r Repository) GetMessages(p GetMessages) []Message {
	messages := make([]Message, 0, len(tables.Messages))
	for _, message := range tables.Messages {
		if message.Receiver == p.UserName {
			messages = append(messages, message)
		}
	}
	return messages
}

type BroadcastPetLostMessage struct {
	PetName string `json:"petName"`
	Owner   string `json:"owner"`
}

// 广播宠物走丢消息
func (r Repository) BroadcastPetLostMessage(p BroadcastPetLostMessage) {
	for _, user := range tables.Users {
		// 应只查询走丢宠物附近的用户
		// 因是 demo 项目，故不做处理。
		var msg Message
		if user.UserName != p.Owner {
			msg = CreateMessagePetLostRequest(user.UserName, p.Owner, p.PetName)
		} else {
			msg = CreateMessagePetLostHint(p.Owner, p.PetName)
		}
		r.createMessage(msg)
	}
}

type GetUserFullInfo struct {
	UserName string
}

// 获取用户数据
func (r Repository) GetUserFullInfo(p GetUserFullInfo) (UserFullInfo, error) {
	user, err := r.userByName(p.UserName)
	if err != nil {
		return UserFullInfo{}, err
	}
	home, _ := r.getHome(p.UserName)
	pets := r.petsByOwner(p.UserName)
	return UserFullInfo{
		User: user,
		Home: home,
		Pets: pets,
	}, nil
}

type UpdatePetLocation struct {
	PetName   string  `json:"petName"`
	Owner     string  `json:"owner"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// 更新宠物位置
func (r Repository) UpdatePetLocation(p UpdatePetLocation) {
	for i, pet := range tables.Pets {
		if pet.Owner == p.Owner && pet.PetName == p.PetName {
			tables.Pets[i].Latitude = p.Latitude
			tables.Pets[i].Longitude = p.Longitude
			return
		}
	}
}

// 获取用户
func (r Repository) userByName(userName string) (User, error) {
	for _, user := range tables.Users {
		if user.UserName == userName {
			return user, nil
		}
	}
	return User{}, ErrUserNotFound
}

// 用户的家
func (r Repository) getHome(owner string) (Home, bool) {
	for _, home := range tables.Homes {
		if home.Owner == owner {
			return home, true
		}
	}
	return Home{}, false
}

// 用户的宠物
func (r Repository) petsByOwner(owner string) []Pet {
	pets := make([]Pet, 0, len(tables.Pets))
	for _, pet := range tables.Pets {
		if pet.Owner == owner {
			pets = append(pets, pet)
		}
	}
	return pets
}

// 创建消息
func (r Repository) createMessage(p Message) {
	tables.Messages = append(tables.Messages, p)
}

// 保存数据
func (r Repository) Save() error {
	bs, err := json.MarshalIndent(tables, "", "    ")
	if err != nil {
		return err
	}
	return os.WriteFile(tableName, bs, 0666)
}

// 加载数据
func (r Repository) Load() error {
	bs, err := os.ReadFile(tableName)
	if err != nil {
		return err
	}
	return json.Unmarshal(bs, &tables)
}
