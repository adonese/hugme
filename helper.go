package main

import (
	"net/http"
	"time"
	"github.com/google/uuid"
)

func setCookie(w http.ResponseWriter) {
	v := uuid.New().String()
	expiration := time.Now().Add(365 * 24 * time.Hour)

	c := &http.Cookie{
		Name: "hugme", Value: v, Expires: expiration,
	}
	http.SetCookie(w, c)
}
