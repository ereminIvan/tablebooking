package service

import (
	"math/rand"
	"time"
)

// IRandom interface of random service
type IRandom interface {
	Runes(n int) string
	Digits(n int) string
}

type randService struct{}

// NewRand create instance of rand service
func NewRand() IRandom {
	rand.Seed(time.Now().UnixNano())
	return &randService{}
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
var digitRunes = []rune("1234567890")

func (r *randService) generate(n int, runes []rune) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = runes[rand.Intn(len(runes))]
	}
	return string(b)
}

// Runes gave string runes with given length
func (r *randService) Runes(n int) string {
	return r.generate(n, letterRunes)
}

// Digits gave string runes of digits with given length
func (r *randService) Digits(n int) string {
	return r.generate(n, digitRunes)
}
