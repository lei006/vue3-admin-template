package utils

import (
	"math/rand"
	"time"
)

// RandomString 生成指定长度和字符集的随机字符串
func RandomString(length int, numbers, letters, specials bool) string {
	rand.Seed(time.Now().UnixNano())

	var charSet []rune

	if numbers {
		charSet = append(charSet, '0', '1', '2', '3', '4', '5', '6', '7', '8', '9')
	}
	if letters {
		charSet = append(charSet, 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z', 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z')
	}
	if specials {
		charSet = append(charSet, '!', '@', '#', '$', '%', '^', '&', '*', '(', ')', '-', '_', '+', '=', '{', '}', '[', ']', '|', '\\', ';', ':', ',', '.', '<', '>', '/', '?', '`', '~')
	}

	if len(charSet) == 0 {
		panic("At least one character type (numbers, letters, or specials) must be enabled.")
	}

	result := make([]rune, length)
	for i := range result {
		result[i] = charSet[rand.Intn(len(charSet))]
	}
	return string(result)
}
