package main

import (
	"errors"
	"log/slog"
	"net/http"
)

type Handler struct {
	Logger     *slog.Logger
	Repository Repository
}

func (h Handler) Register(mux *http.ServeMux) {
	// 普通路由 无访问控制
	mux.HandleFunc("POST /api/v1/login", h.Log(h.Login))
	// 普通路由 demo 工程 未加访问控制中间件
	mux.HandleFunc("GET /api/v1/message/{userName}", h.Log(h.GetMessages))
	mux.HandleFunc("POST /api/v1/broadcast/pet/lost", h.Log(h.BroadcastPetLostMessage))
	mux.HandleFunc("GET /api/v1/user/info/full/{userName}", h.Log(h.GetUserFullInfo))
	mux.HandleFunc("PUT /api/v1/pet/location", h.Log(h.UpdatePetLocation))
	// admin 路由 demo 工程 未加访问控制中间件
	mux.HandleFunc("/api/v1/admin/data/save", h.Log(h.Save))
	mux.HandleFunc("/api/v1/admin/data/load", h.Log(h.Load))
}

func (h Handler) Log(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		h.Logger.Info("request", slog.String("method", r.Method), slog.Any("url", r.URL))
		next(w, r)
	}
}

// 登录
func (h Handler) Login(w http.ResponseWriter, r *http.Request) {
	var a GetUser
	err := ParseRequest(r, &a)
	if err != nil {
		h.Logger.Error("解析请求失败", slog.Any("err", err))
		ResponseErr(w, ErrBadRequest, http.StatusBadRequest)
		return
	}
	// demo 工程 跳过数据校验
	_, err = h.Repository.GetUser(a)
	if errors.Is(err, ErrUserNotFound) {
		h.Logger.Error("用户名或密码不存在", slog.Any("err", err))
		ResponseErr(w, err, http.StatusNotFound)
		return
	}
	if err != nil {
		h.Logger.Error("获取用户失败", slog.Any("err", err))
		ResponseErr(w, err, http.StatusInternalServerError)
		return
	}
	_ = ResponseData(w, Empty{})
}

// 获取信息
func (h Handler) GetMessages(w http.ResponseWriter, r *http.Request) {
	var a GetMessages
	a.UserName = r.PathValue("userName")
	if a.UserName == "" {
		h.Logger.Error("用户名为空")
		ResponseErr(w, ErrBadRequest, http.StatusBadRequest)
		return
	}
	li := h.Repository.GetMessages(a)
	_ = ResponseData(w, li)
}

// 广播宠物走丢信息
func (h Handler) BroadcastPetLostMessage(w http.ResponseWriter, r *http.Request) {
	var a BroadcastPetLostMessage
	err := ParseRequest(r, &a)
	if err != nil {
		h.Logger.Error("解析请求失败", slog.Any("err", err))
		ResponseErr(w, ErrBadRequest, http.StatusBadRequest)
		return
	}
	// demo 工程 跳过数据校验
	h.Repository.BroadcastPetLostMessage(a)
	_ = ResponseData(w, Empty{})
}

// 获取用户全量信息
func (h Handler) GetUserFullInfo(w http.ResponseWriter, r *http.Request) {
	var a GetUserFullInfo
	a.UserName = r.PathValue("userName")
	if a.UserName == "" {
		h.Logger.Error("用户名为空")
		ResponseErr(w, ErrBadRequest, http.StatusBadRequest)
		return
	}
	info, err := h.Repository.GetUserFullInfo(a)
	if err != nil {
		h.Logger.Error("用户数据获取失败", slog.Any("err", err))
		ResponseErr(w, err, http.StatusInternalServerError)
		return
	}
	_ = ResponseData(w, info)
}

// 更新宠物位置
func (h Handler) UpdatePetLocation(w http.ResponseWriter, r *http.Request) {
	var a UpdatePetLocation
	err := ParseRequest(r, &a)
	if err != nil {
		h.Logger.Error("解析请求失败", slog.Any("err", err))
		ResponseErr(w, ErrBadRequest, http.StatusBadRequest)
		return
	}
	// demo 工程 跳过数据校验
	h.Repository.UpdatePetLocation(a)
	_ = ResponseData(w, Empty{})
}

// 保存数据
func (h Handler) Save(w http.ResponseWriter, r *http.Request) {
	err := h.Repository.Save()
	if err != nil {
		h.Logger.Error("数据存储失败", slog.Any("err", err))
		ResponseErr(w, err, http.StatusInternalServerError)
		return
	}
	_ = ResponseData(w, Empty{})
}

// 加载数据
func (h Handler) Load(w http.ResponseWriter, r *http.Request) {
	err := h.Repository.Load()
	if err != nil {
		h.Logger.Warn("数据加载失败", slog.Any("err", err))
		ResponseErr(w, err, http.StatusInternalServerError)
		return
	}
	_ = ResponseData(w, Empty{})
}
