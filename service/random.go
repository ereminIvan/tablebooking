package service

import (
	"math/rand"
	"time"
)

type IRandom interface {
	Runes(n int) string
}

type randService struct {
}

func NewRand() IRandom {
	rand.Seed(time.Now().UnixNano())
	return &randService{}
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

func (r *randService) Runes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
