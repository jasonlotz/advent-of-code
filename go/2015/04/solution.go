package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

var SECRET_KEY = "bgvyzdsv"

func main() {
	fmt.Println("Part 1:", findLowestNumber(SECRET_KEY, "00000"))
	fmt.Println("Part 2:", findLowestNumber(SECRET_KEY, "000000"))
}

func findLowestNumber(secretKey string, prefix string) int {
	var i int
	for i = 0; ; i++ {
		if checkHashPrefix(secretKey, prefix, i) {
			break
		}
	}

	return i
}

func checkHashPrefix(secretKey string, prefix string, number int) bool {
	value := secretKey + fmt.Sprint(number)
	hash := md5.Sum([]byte(value))
	hashStr := hex.EncodeToString(hash[:])

	return hashStr[:len(prefix)] == prefix
}
