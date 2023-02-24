package utils

import (
	"math/rand"
	"time"
)

func GetRandomNum() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(100)
}
