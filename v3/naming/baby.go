package naming

import "math/rand"

func RandomString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

func CreateBabyName(male bool, minLen int) string {
	name := ""
	if male {
		name = "小王子"
	} else {
		name = "小公主"
	}
	return name + RandomString(minLen)
}
