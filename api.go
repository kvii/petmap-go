package main

import (
	"encoding/json"
	"io"
	"net/http"
)

type Empty = struct{}

func ParseRequest(r *http.Request, a any) error {
	bs, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(bs, a)
}

func ResponseData(w http.ResponseWriter, data any) error {
	bs, err := json.Marshal(data)
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(bs)
	return err
}

func ResponseErr(w http.ResponseWriter, err error, code int) {
	http.Error(w, err.Error(), code)
}
