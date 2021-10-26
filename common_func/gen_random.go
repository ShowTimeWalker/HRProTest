package common_func

import (
	"math/rand"
	"strconv"
	"time"
)

func GenRandPhone(header string) string {
	rand.Seed(time.Now().Unix())
	randNumberString := strconv.Itoa(int(rand.Uint32()%90000000 + 10000000))
	userPhoneNumber := header + randNumberString
	println("phone number:", userPhoneNumber)
	return userPhoneNumber
}
