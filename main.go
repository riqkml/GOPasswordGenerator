package main

import (
	"fmt"
	"strings"
)

func Reverse(str string) string {
	var backward string

	for i := len(str) - 1; i >= 0; i-- {
		backward += string(str[i])
	}
	return backward
}

func Generate(str string) string {
	str = Reverse(str)

	//contain number or symbol should just reverse
	// for _, number := range str {
	// 	switch number {
	// 	case '1', '2', '3', '4', '5', '6', '7', '8', '9', '0':
	// 		return str
	// 	case '!', '@', '#', '$', '%', '^', '&', '*', '_', '-':
	// 		return str
	// 	}
	// }
	//HOW TO CHANGE VOWEL *basic way
	var replace string
	for _, vowels := range str {
		if string(vowels) == "a" {
			replace += "e"
		} else if string(vowels) == "e" {
			replace += "i"
		} else if string(vowels) == "i" {
			replace += "o"
		} else if string(vowels) == "o" {
			replace += "u"
		} else if string(vowels) == "u" {
			replace += "a"
		} else if string(vowels) == "A" {
			replace += "E"
		} else if string(vowels) == "E" {
			replace += "I"
		} else if string(vowels) == "I" {
			replace += "O"
		} else if string(vowels) == "O" {
			replace += "U"
		} else if string(vowels) == "U" {
			replace += "A"
		} else {
			replace += string(vowels)
		}
	}
	var change string
	for _, Uplow := range replace {
		// change space and capital to lower vice versain
		if string(Uplow) == " " { //belum ke detect ditest but outputnya bener
			continue
		} else if string(Uplow) == strings.ToUpper(string(Uplow)) { // change uppercase to lowercase
			change += strings.ToLower(string(Uplow))
		} else { // change lower to upper
			change += strings.ToUpper(string(Uplow))
		}
	}
	return change
}
func IsLetter(s string) bool {
	for _, r := range s {
		if (r < 'a' || r > 'z') && (r < 'A' || r > 'Z') && (r < '0' || r > '9') {
			return false
		}
	}
	return true
}

func CheckPassword(str string) string {
	str = Generate(str)
	//check password criteria sangat lemah,lemah,sedang,kuat
	var password string
	//numeric := "1234567890"
	//alphabet := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	//Alphasym := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!@#$%^&*()_-"
	//numsym := "1234567890!@#$%^&*()_-"
	// Alphanumeric := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ12345678910"
	//combinasi := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ12345678910!@#$%^&*()_-"

	// gabisa ngedetect ketentuan alphanumeric dll
	if len(str) >= 14 {
		password += "kuat"
	} else if len(str) >= 7 && !IsLetter(str) {
		password += "sedang"
	} else if len(str) >= 7 && IsLetter(str) {
		password += "lemah" //gak kedetect
	} else {
		password += "sangat lemah"
	}
	return password
}

func PasswordGenerator(base string) (string, string) { // error should generate password ??
	var generate = Generate(base)
	var check = CheckPassword(generate)

	return generate, check
}

func main() {
	data := "2133! #123)$#" // bisa digunakan untuk melakukan debug

	res, check := PasswordGenerator(data)

	fmt.Println(res)
	fmt.Println(check)
}
