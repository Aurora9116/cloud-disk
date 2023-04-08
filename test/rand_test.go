package test

import (
	"cloud-disk/core/define"
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func TestRandCode(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	s := ""
	for i := 0; i < define.CodeLength; i++ {
		s = s + strconv.Itoa(rand.Intn(10))
	}
	fmt.Println(s)
}
