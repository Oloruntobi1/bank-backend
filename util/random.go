package util

import (
	"math/rand"
	"strings"
	"time"

	"github.com/Oloruntobi1/bankBackend/helper"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max - min + 1)
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n ; i++ {
		c := alphabet[rand.Intn(k)]

		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomOwner() string {
	return RandomString(6)
}

func RandomPassword() string {
	digit := RandomInt(6, 10)
	res := RandomString(int(digit))
	return helper.HashAndSalt([]byte(res))
}

func RandomEmail() string{
	return RandomOwner() + "@gmail.com"
}

func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

func RandomType(types []string) string{

	n := len(types)

	return types[rand.Intn(n)]

}

func RandomBankType() string{

	accType := []string{"Current", "Savings", "Fixed"}

	return RandomType(accType)

}

func RandomCurrencyType() string {
	currencyType := []string{"Yen", "EUR", "Pounds"}

	return RandomType(currencyType)
}